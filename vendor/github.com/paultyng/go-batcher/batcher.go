package batcher // import "github.com/paultyng/go-batcher"

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type result struct {
	value interface{}
	err   error
}

type paramCh struct {
	Param interface{}
	Ch    chan result
}

// New returns a new Batcher that will run the getBatch func after enough time
// has elapsed since the last Get call. getBatch MUST return the same len as the
// input slice it is passed. If a value is not found, it can return a nil for
// that index.
func New(after time.Duration, getBatch func([]interface{}) ([]interface{}, error)) *Batcher {
	return &Batcher{
		after:    after,
		getBatch: getBatch,
	}
}

// Batcher represents the Batcher instance. Consumers should use the Get method to
// queue new requests for the batch.
type Batcher struct {
	after time.Duration

	// startOnce is used to ensure the start method is only invoked once
	startOnce sync.Once

	getBatch func([]interface{}) ([]interface{}, error)

	// the mutex covers only ch, timer, and params
	mu     sync.Mutex
	ch     chan paramCh
	timer  *time.Timer
	params []paramCh
}

func (b *Batcher) start() {
	b.startOnce.Do(func() {
		func() {
			b.mu.Lock()
			defer b.mu.Unlock()

			b.ch = make(chan paramCh)
		}()

		go func() {
			for {
				n := <-b.ch

				func() {
					b.mu.Lock()
					defer b.mu.Unlock()

					b.params = append(b.params, n)

					if b.timer != nil {
						b.timer.Stop()
					}
					b.timer = time.AfterFunc(b.after, func() {
						var params []paramCh
						func() {
							b.mu.Lock()
							defer b.mu.Unlock()

							params = b.params
							b.params = nil
						}()

						b.handleBatchInternal(params)
					})
				}()
			}
		}()
	})
}

func (b *Batcher) handleBatchInternal(paramChs []paramCh) {
	params := make([]interface{}, len(paramChs))
	for i, pCh := range paramChs {
		params[i] = pCh.Param
	}

	values, err := b.getBatch(params)
	// if you get an error on any request, just return the error for all items
	if err != nil {
		for _, pCh := range paramChs {
			pCh.Ch <- result{
				err: err,
			}
		}
		return
	}
	if len(values) != len(params) {
		// this is a bad implementation of getBatch, just panic
		panic("getBatch must return the same len of slice as is provided and the indicies must match the input")
	}
	for i, pCh := range paramChs {
		v := values[i]
		pCh.Ch <- result{
			value: v,
		}
	}
}

// Get queues a new request for the batch. If enough time has elapsed since the last Get
// call, the batch will be processed.
func (b *Batcher) Get(ctx context.Context, p interface{}) (interface{}, error) {
	b.start()

	result := make(chan result)

	b.ch <- paramCh{
		Param: p,
		Ch:    result,
	}

	select {
	case r := <-result:
		return r.value, r.err
	case <-ctx.Done():
		err := ctx.Err()
		if err == nil {
			// this is possibly unnecessary as Err should always be populated
			err = fmt.Errorf("context terminated unexpectedly fetching item")
		}
		return nil, err
	}
}
