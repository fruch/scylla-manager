// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSetWriteRPCTimeoutParams creates a new SetWriteRPCTimeoutParams object
// with the default values initialized.
func NewSetWriteRPCTimeoutParams() *SetWriteRPCTimeoutParams {
	var ()
	return &SetWriteRPCTimeoutParams{

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewSetWriteRPCTimeoutParamsWithTimeout creates a new SetWriteRPCTimeoutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetWriteRPCTimeoutParamsWithTimeout(timeout time.Duration) *SetWriteRPCTimeoutParams {
	var ()
	return &SetWriteRPCTimeoutParams{

		requestTimeout: timeout,
	}
}

// NewSetWriteRPCTimeoutParamsWithContext creates a new SetWriteRPCTimeoutParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetWriteRPCTimeoutParamsWithContext(ctx context.Context) *SetWriteRPCTimeoutParams {
	var ()
	return &SetWriteRPCTimeoutParams{

		Context: ctx,
	}
}

// NewSetWriteRPCTimeoutParamsWithHTTPClient creates a new SetWriteRPCTimeoutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetWriteRPCTimeoutParamsWithHTTPClient(client *http.Client) *SetWriteRPCTimeoutParams {
	var ()
	return &SetWriteRPCTimeoutParams{
		HTTPClient: client,
	}
}

/*SetWriteRPCTimeoutParams contains all the parameters to send to the API endpoint
for the set write rpc timeout operation typically these are written to a http.Request
*/
type SetWriteRPCTimeoutParams struct {

	/*Timeout
	  timeout in seconds

	*/
	Timeout int32

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the set write rpc timeout params
func (o *SetWriteRPCTimeoutParams) WithRequestTimeout(timeout time.Duration) *SetWriteRPCTimeoutParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the set write rpc timeout params
func (o *SetWriteRPCTimeoutParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the set write rpc timeout params
func (o *SetWriteRPCTimeoutParams) WithContext(ctx context.Context) *SetWriteRPCTimeoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set write rpc timeout params
func (o *SetWriteRPCTimeoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set write rpc timeout params
func (o *SetWriteRPCTimeoutParams) WithHTTPClient(client *http.Client) *SetWriteRPCTimeoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set write rpc timeout params
func (o *SetWriteRPCTimeoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimeout adds the timeout to the set write rpc timeout params
func (o *SetWriteRPCTimeoutParams) WithTimeout(timeout int32) *SetWriteRPCTimeoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set write rpc timeout params
func (o *SetWriteRPCTimeoutParams) SetTimeout(timeout int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *SetWriteRPCTimeoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	// query param timeout
	qrTimeout := o.Timeout
	qTimeout := swag.FormatInt32(qrTimeout)
	if qTimeout != "" {
		if err := r.SetQueryParam("timeout", qTimeout); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
