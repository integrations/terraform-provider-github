# ghinstallation

[![GoDoc](https://godoc.org/github.com/kfcampbell/ghinstallation?status.svg)](https://godoc.org/github.com/kfcampbell/ghinstallation)

This library is forked from [bradleyfalzon/ghinstallation](https://github.com/bradleyfalzon/ghinstallation), which was created to provide GitHub Apps authentication helpers for [google/go-github](https://github.com/google/go-github). This fork is designed to work with [octokit/go-sdk](https://github.com/octokit/go-sdk).

`ghinstallation` provides `Transport`, which implements `http.RoundTripper` to
provide authentication as an installation for GitHub Apps.

See
https://developer.github.com/apps/building-integrations/setting-up-and-registering-github-apps/about-authentication-options-for-github-apps/

## Installation

Get the package:

```bash
go get -u github.com/kfcampbell/ghinstallation
```

## go-sdk example

See [go-sdk](https://github.com/octokit/go-sdk/blob/apps-poc/cmd/app-example/main.go) examples for instructions on usage with the go-sdk library.

## Manual usage example

Create or obtain a transport, then wrap it with an Apps transport using your private key, client ID, and installation ID.

```go

existingTransport := http.DefaultTransport

appTransport, err := ghinstallation.NewKeyFromFile(existingTransport, "your-client-ID", yourInstallationIDInt, "path/to/your/pem/file.pem")
if err != nil {
  return nil, fmt.Errorf("failed to create transport from GitHub App using clientID: %v", err)
}
// use the created appTransport in your HttpClient
```

### What are client ID, app ID, and installation ID?

These are fields unique to your application that GitHub uses to identify your App and issue its credentials. Client ID and App ID are interchangeable, and client ID is preferred. Both can be obtained by visiting the App's page > App settings. The URL for your App is "https://github.com/apps/{yourAppName}".

The App page will look something like this, with the App settings link on the right sidebar:

![App page](https://github.com/kfcampbell/ghinstallation/assets/9327688/db529326-e994-443e-bef5-98fcf0f5ab20)

The App settings page will look like this, and allow you to copy the App ID and the client ID:

![App settings page](https://github.com/kfcampbell/ghinstallation/assets/9327688/dc409378-0fba-45c9-bbbe-6490e967140f)

The installation ID is specific to an installation of that App to a specific organization or user. You can find it in the URL when viewing an App's installation, when the URL will look something like this: "https://github.com/organizations/{yourAppName}/settings/installations/{installationID}".

## Customizing signing behavior

Users can customize signing behavior by passing in a
[Signer](https://pkg.go.dev/github.com/kfcampbell/ghinstallation#Signer)
implementation when creating an
[AppsTransport](https://pkg.go.dev/github.com/kfcampbell/ghinstallation#AppsTransport).
For example, this can be used to create tokens backed by keys in a KMS system.

```go
signer := &myCustomSigner{
  key: "https://url/to/key/vault",
}
appsTransport := NewAppsTransportWithOptions(http.DefaultTransport, "your-client-ID", WithSigner(signer))
transport := NewFromAppsTransport(appsTransport, yourInstallationIDInt)
```

## License

[Apache 2.0](LICENSE)

## Dependencies

- [github.com/golang-jwt/jwt-go](https://github.com/golang-jwt/jwt-go)
