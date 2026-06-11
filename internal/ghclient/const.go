package ghclient

import "time"

// clientTimeout defines the timeout duration for GitHub API requests. This is set to 5 minutes to accommodate potentially long-running operations while still providing a reasonable upper limit on request duration.
const clientTimeout = 5 * time.Minute
