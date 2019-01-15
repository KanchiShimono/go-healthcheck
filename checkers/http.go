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
	url        *url.URL
	statusCode int
	timeout    time.Duration
}

// Check is the implementation of Checker interface.checker.
// Check the status code of GET method.
func (hc *HTTPChecker) Check() error {
	req, err := http.NewRequest("GET", hc.url.String(), nil)
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: hc.timeout,
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != hc.statusCode {
		return errors.New("Returned unexpected code: " + strconv.Itoa(res.StatusCode) + ". Expected: " + strconv.Itoa(hc.statusCode))
	}

	return nil
}
