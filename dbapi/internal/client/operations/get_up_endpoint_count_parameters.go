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

// NewGetUpEndpointCountParams creates a new GetUpEndpointCountParams object
// with the default values initialized.
func NewGetUpEndpointCountParams() *GetUpEndpointCountParams {

	return &GetUpEndpointCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUpEndpointCountParamsWithTimeout creates a new GetUpEndpointCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUpEndpointCountParamsWithTimeout(timeout time.Duration) *GetUpEndpointCountParams {

	return &GetUpEndpointCountParams{

		timeout: timeout,
	}
}

// NewGetUpEndpointCountParamsWithContext creates a new GetUpEndpointCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUpEndpointCountParamsWithContext(ctx context.Context) *GetUpEndpointCountParams {

	return &GetUpEndpointCountParams{

		Context: ctx,
	}
}

// NewGetUpEndpointCountParamsWithHTTPClient creates a new GetUpEndpointCountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUpEndpointCountParamsWithHTTPClient(client *http.Client) *GetUpEndpointCountParams {

	return &GetUpEndpointCountParams{
		HTTPClient: client,
	}
}

/*GetUpEndpointCountParams contains all the parameters to send to the API endpoint
for the get up endpoint count operation typically these are written to a http.Request
*/
type GetUpEndpointCountParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get up endpoint count params
func (o *GetUpEndpointCountParams) WithTimeout(timeout time.Duration) *GetUpEndpointCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get up endpoint count params
func (o *GetUpEndpointCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get up endpoint count params
func (o *GetUpEndpointCountParams) WithContext(ctx context.Context) *GetUpEndpointCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get up endpoint count params
func (o *GetUpEndpointCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get up endpoint count params
func (o *GetUpEndpointCountParams) WithHTTPClient(client *http.Client) *GetUpEndpointCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get up endpoint count params
func (o *GetUpEndpointCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUpEndpointCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
