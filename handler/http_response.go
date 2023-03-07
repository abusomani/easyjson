package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// HttpResponseHandler implements the Handler interface
// It takes in pointer to a http.Response on which the operations have to be performed.
type HttpResponseHandler struct {
	r *http.Response
}

// NewHttpResponseHandler takes in a pointer to http.Response.
// It returns a HttpResponseHandler instance.
func NewHttpResponseHandler(res *http.Response) *HttpResponseHandler {
	return &HttpResponseHandler{
		r: res,
	}
}

// Read function returns the bytes read from the given http.Response or an error in case something went wrong.
func (rh *HttpResponseHandler) Read() ([]byte, error) {
	if rh.r == nil {
		return nil, fmt.Errorf("response is nil")
	}
	if rh.r.Body == nil {
		return nil, fmt.Errorf("response body is nil")
	}
	defer func() {
		if err := rh.r.Body.Close(); err != nil {
			log.Fatalf("error closing : %s", err.Error())
		}
	}()
	body, err := io.ReadAll(rh.r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading body: %s", err.Error())
	}
	return body, nil
}

// Write operation is not allowed on a http.Response
func (rh *HttpResponseHandler) Write(input []byte) error {
	return fmt.Errorf("not allowed")
}
