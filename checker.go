package checker

import "net/http"

type ckeckFunc func() error

// Creates a simple HTTP server listening on port 8080 with only GET:/live and
// GET:/ready endpoints. It takes liveness and readiness functions as
// parameters. Each parameter is a ckeckFunc function. It simply returns an
// error if the probe fails.
//
// Example:
//
//  // Use a separate coroutine
//  go checker.HTTP(
//    func() error {
//      if !app.Ready() {
//        return errors.New("App is not ready")
//      }
//
//      return nil
//    },
//    func() error {
//      if !app.Live() {
//        return errors.New("App is not live")
//      }
//
//      return nil
//    },
//  )
func HTTP(liveness, readiness ckeckFunc) error {
	http.HandleFunc("/live", handleHTTPCheck(liveness))
	http.HandleFunc("/ready", handleHTTPCheck(readiness))

	return http.ListenAndServe(":8080", nil)
}

func handleHTTPCheck(c ckeckFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if err := c(); err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
	}
}
