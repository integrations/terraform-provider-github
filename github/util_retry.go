package github

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

const (
	defaultRetryDelay      = 1 * time.Second
	defaultRetryMaxRetries = 5
	defaultRetryTimeout    = 1 * time.Minute
)

// retryOptions defines the options for retrying a function call.
type retryOptions struct {
	delay      time.Duration
	maxRetries int
	timeout    time.Duration
}

// retryUntilOK retries the given function until it returns a value and true within the default timeout. If the function returns an error, it will be returned immediately. If the function returns false, it will be retried until the timeout is reached.
func retryUntilOK[T any](ctx context.Context, f func() (T, bool, error), opts *retryOptions) (T, error) {
	if opts == nil {
		opts = &retryOptions{
			delay:      defaultRetryDelay,
			maxRetries: defaultRetryMaxRetries,
			timeout:    defaultRetryTimeout,
		}
	}

	conf := &retry.StateChangeConf{
		Pending: []string{"missing"},
		Target:  []string{"found"},
		Refresh: func() (any, string, error) {
			val, ok, err := f()
			if err != nil {
				return nil, "", err
			}
			if !ok {
				return nil, "missing", nil
			}
			return val, "found", nil
		},
		Timeout:    opts.timeout,
		Delay:      opts.delay,
		MinTimeout: 1 * time.Second,
	}

	a, err := conf.WaitForStateContext(ctx)
	if err != nil {
		var zero T
		return zero, err
	}

	val, ok := a.(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("unexpected type %T", a)
	}

	return val, nil
}

// retryUntilResourceFound retries the given function and if there is no error it returns the value. If the error is a [github.ErrorResponse] with a 404 status code, it will be retried until the timeout is reached. If the error is any other error, it will be returned immediately.
func retryUntilResourceFound[T any](ctx context.Context, f func() (T, error), opts *retryOptions) (T, error) {
	return retryUntilOK(ctx, func() (T, bool, error) {
		val, err := f()
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response != nil && ghErr.Response.StatusCode == 404 {
				return val, false, nil
			}
			return val, false, err
		}
		return val, true, nil
	}, opts)
}
