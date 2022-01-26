# Checker

[![GoReference](https://pkg.go.dev/badge/github.com/remyduthu/checker.svg)](https://pkg.go.dev/github.com/remyduthu/checker)

Checker provides simple liveness and readiness checkers. You can use this
package to configure [Kubernetes
probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/).

## Checkers

### HTTP

Creates a simple HTTP server listening on port `8080` with only `GET:/live` and
`GET:/ready` endpoints. It takes `liveness` and `readiness` functions as
parameters. Each parameter is a `ckeckFunc` function. It simply returns an error
if the probe fails.

Example:

```go
// Use a separate coroutine
go checker.HTTP(
  func() error {
    if !app.Ready() {
      return errors.New("App is not ready")
    }

    return nil
  },
	func() error {
    if !app.Live() {
      return errors.New("App is not live")
    }

    return nil
	},
)
```
