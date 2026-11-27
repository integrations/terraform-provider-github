package httpx

import "time"

// headerRetryAfter is the header a server uses to say how long to wait before retrying.
const headerRetryAfter = "Retry-After"

// maxRetryAfterSeconds is the largest Retry-After value in seconds that fits in a [time.Duration].
const maxRetryAfterSeconds = uint64((1<<63 - 1) / time.Second)

const (
	// DefaultRetryMaxRetries is the default maximum number of retries for failed requests.
	DefaultRetryMaxRetries = 3

	// DefaultRetryWaitMin is the default minimum wait duration between retries.
	DefaultRetryWaitMin = time.Second

	// DefaultRetryWaitMax is the default maximum wait duration between retries.
	DefaultRetryWaitMax = time.Minute

	// DefaultRetryJitter is the default jitter duration added to the wait between retries.
	DefaultRetryJitter = 100 * time.Millisecond
)
