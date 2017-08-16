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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetEndpointDowntimeParams creates a new GetEndpointDowntimeParams object
// with the default values initialized.
func NewGetEndpointDowntimeParams() *GetEndpointDowntimeParams {
	var ()
	return &GetEndpointDowntimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEndpointDowntimeParamsWithTimeout creates a new GetEndpointDowntimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEndpointDowntimeParamsWithTimeout(timeout time.Duration) *GetEndpointDowntimeParams {
	var ()
	return &GetEndpointDowntimeParams{

		timeout: timeout,
	}
}

// NewGetEndpointDowntimeParamsWithContext creates a new GetEndpointDowntimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEndpointDowntimeParamsWithContext(ctx context.Context) *GetEndpointDowntimeParams {
	var ()
	return &GetEndpointDowntimeParams{

		Context: ctx,
	}
}

// NewGetEndpointDowntimeParamsWithHTTPClient creates a new GetEndpointDowntimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEndpointDowntimeParamsWithHTTPClient(client *http.Client) *GetEndpointDowntimeParams {
	var ()
	return &GetEndpointDowntimeParams{
		HTTPClient: client,
	}
}

/*GetEndpointDowntimeParams contains all the parameters to send to the API endpoint
for the get endpoint downtime operation typically these are written to a http.Request
*/
type GetEndpointDowntimeParams struct {

	/*Addr
	  The endpoint address

	*/
	Addr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get endpoint downtime params
func (o *GetEndpointDowntimeParams) WithTimeout(timeout time.Duration) *GetEndpointDowntimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get endpoint downtime params
func (o *GetEndpointDowntimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get endpoint downtime params
func (o *GetEndpointDowntimeParams) WithContext(ctx context.Context) *GetEndpointDowntimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get endpoint downtime params
func (o *GetEndpointDowntimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get endpoint downtime params
func (o *GetEndpointDowntimeParams) WithHTTPClient(client *http.Client) *GetEndpointDowntimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get endpoint downtime params
func (o *GetEndpointDowntimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddr adds the addr to the get endpoint downtime params
func (o *GetEndpointDowntimeParams) WithAddr(addr string) *GetEndpointDowntimeParams {
	o.SetAddr(addr)
	return o
}

// SetAddr adds the addr to the get endpoint downtime params
func (o *GetEndpointDowntimeParams) SetAddr(addr string) {
	o.Addr = addr
}

// WriteToRequest writes these params to a swagger request
func (o *GetEndpointDowntimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addr
	if err := r.SetPathParam("addr", o.Addr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
