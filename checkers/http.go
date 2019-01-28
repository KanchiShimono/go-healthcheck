package checkers

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// HTTPChecker checks the HTTP response code when request url
type HTTPChecker struct {
	URL        *url.URL
	StatusCode int
	Timeout    time.Duration
}

// Check is the implementation of Checker interface.checker.
// Check the status code of GET method.
func (hc *HTTPChecker) Check() error {
	req, err := http.NewRequest("GET", hc.URL.String(), nil)
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: hc.Timeout,
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != hc.StatusCode {
		return errors.New("Returned unexpected code: " + strconv.Itoa(res.StatusCode) + ". Expected: " + strconv.Itoa(hc.StatusCode))
	}

	return nil
}
