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

// NewSetReadRPCTimeoutParams creates a new SetReadRPCTimeoutParams object
// with the default values initialized.
func NewSetReadRPCTimeoutParams() *SetReadRPCTimeoutParams {
	var ()
	return &SetReadRPCTimeoutParams{

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewSetReadRPCTimeoutParamsWithTimeout creates a new SetReadRPCTimeoutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetReadRPCTimeoutParamsWithTimeout(timeout time.Duration) *SetReadRPCTimeoutParams {
	var ()
	return &SetReadRPCTimeoutParams{

		requestTimeout: timeout,
	}
}

// NewSetReadRPCTimeoutParamsWithContext creates a new SetReadRPCTimeoutParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetReadRPCTimeoutParamsWithContext(ctx context.Context) *SetReadRPCTimeoutParams {
	var ()
	return &SetReadRPCTimeoutParams{

		Context: ctx,
	}
}

// NewSetReadRPCTimeoutParamsWithHTTPClient creates a new SetReadRPCTimeoutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetReadRPCTimeoutParamsWithHTTPClient(client *http.Client) *SetReadRPCTimeoutParams {
	var ()
	return &SetReadRPCTimeoutParams{
		HTTPClient: client,
	}
}

/*SetReadRPCTimeoutParams contains all the parameters to send to the API endpoint
for the set read rpc timeout operation typically these are written to a http.Request
*/
type SetReadRPCTimeoutParams struct {

	/*Timeout
	  The timeout in second

	*/
	Timeout int32

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the set read rpc timeout params
func (o *SetReadRPCTimeoutParams) WithRequestTimeout(timeout time.Duration) *SetReadRPCTimeoutParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the set read rpc timeout params
func (o *SetReadRPCTimeoutParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the set read rpc timeout params
func (o *SetReadRPCTimeoutParams) WithContext(ctx context.Context) *SetReadRPCTimeoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set read rpc timeout params
func (o *SetReadRPCTimeoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set read rpc timeout params
func (o *SetReadRPCTimeoutParams) WithHTTPClient(client *http.Client) *SetReadRPCTimeoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set read rpc timeout params
func (o *SetReadRPCTimeoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimeout adds the timeout to the set read rpc timeout params
func (o *SetReadRPCTimeoutParams) WithTimeout(timeout int32) *SetReadRPCTimeoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set read rpc timeout params
func (o *SetReadRPCTimeoutParams) SetTimeout(timeout int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *SetReadRPCTimeoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
