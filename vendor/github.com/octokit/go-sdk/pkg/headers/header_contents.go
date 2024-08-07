package headers

const AuthorizationKey = "Authorization"
const AuthType = "bearer"
const UserAgentKey = "User-Agent"

// TODO(kfcampbell): get the version and binary name from build settings rather than hard-coding
const UserAgentValue = "go-sdk@v0.0.0"

const APIVersionKey = "X-GitHub-Api-Version"

// TODO(kfcampbell): get the version from the generated code somehow
const APIVersionValue = "2022-11-28"

// documentation on rate limit headers is available here:
// https://docs.github.com/en/rest/using-the-rest-api/rate-limits-for-the-rest-api?apiVersion=2022-11-28#checking-the-status-of-your-rate-limit
const XRateLimitRemainingKey = "X-Ratelimit-Remaining"
const XRateLimitResetKey = "X-Ratelimit-Reset"
const RetryAfterKey = "Retry-After"

const XRateLimitLimitKey = "X-Ratelimit-Limit"
const XRateLimitUsedKey = "X-Ratelimit-Used"
const XRateLimitResourceKey = "X-Ratelimit-Resource"
