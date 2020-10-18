package http_request

import (
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
)

type HttpClient struct {
	Client  *resty.Client
	Request *resty.Request
}

func New() HttpClient {

	// Create a Resty Client
	client := resty.New()

	// Retries are configured per client
	client.
		// Set retry count to non zero to enable retries
		SetRetryCount(10).
		// TimeOut
		SetTimeout(5 * time.Second).
		// You can override initial retry wait time.
		// Default is 100 milliseconds.
		SetRetryWaitTime(2 * time.Second).
		// MaxWaitTime can be overridden as well.
		// Default is 2 seconds.
		SetRetryMaxWaitTime(20 * time.Second).
		// SetRetryAfter sets callback to calculate wait time between retries.
		// Default (nil) implies exponential backoff with jitter
		SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
			return 0, errors.New("quota exceeded")
		})

	return HttpClient{
		Client:  client,
		Request: client.R(),
	}
}

func (h HttpClient) SetTrace(header interface{}) HttpClient {
	trace := SetHeader(header)

	h.Request.SetHeader(X_REQUEST_ID, trace.Http_Header.Get(X_REQUEST_ID))
	h.Request.SetHeader(X_B3_TRACEID, trace.Http_Header.Get(X_B3_TRACEID))
	h.Request.SetHeader(X_B3_SPANID, trace.Http_Header.Get(X_B3_SPANID))
	h.Request.SetHeader(X_B3_PARENTSPANID, trace.Http_Header.Get(X_B3_PARENTSPANID))
	h.Request.SetHeader(X_B3_SAMPLED, trace.Http_Header.Get(X_B3_SAMPLED))
	h.Request.SetHeader(X_B3_FLAGS, trace.Http_Header.Get(X_B3_FLAGS))
	h.Request.SetHeader(B3, trace.Http_Header.Get(B3))
	h.Request.SetHeader(X_OT_SPAN_CONTEXT, trace.Http_Header.Get(X_OT_SPAN_CONTEXT))

	h.Request.Header = trace.Http_Header
	return h
}
