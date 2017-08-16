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

// NewGetCasWriteTimeoutsParams creates a new GetCasWriteTimeoutsParams object
// with the default values initialized.
func NewGetCasWriteTimeoutsParams() *GetCasWriteTimeoutsParams {

	return &GetCasWriteTimeoutsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCasWriteTimeoutsParamsWithTimeout creates a new GetCasWriteTimeoutsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCasWriteTimeoutsParamsWithTimeout(timeout time.Duration) *GetCasWriteTimeoutsParams {

	return &GetCasWriteTimeoutsParams{

		timeout: timeout,
	}
}

// NewGetCasWriteTimeoutsParamsWithContext creates a new GetCasWriteTimeoutsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCasWriteTimeoutsParamsWithContext(ctx context.Context) *GetCasWriteTimeoutsParams {

	return &GetCasWriteTimeoutsParams{

		Context: ctx,
	}
}

// NewGetCasWriteTimeoutsParamsWithHTTPClient creates a new GetCasWriteTimeoutsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCasWriteTimeoutsParamsWithHTTPClient(client *http.Client) *GetCasWriteTimeoutsParams {

	return &GetCasWriteTimeoutsParams{
		HTTPClient: client,
	}
}

/*GetCasWriteTimeoutsParams contains all the parameters to send to the API endpoint
for the get cas write timeouts operation typically these are written to a http.Request
*/
type GetCasWriteTimeoutsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cas write timeouts params
func (o *GetCasWriteTimeoutsParams) WithTimeout(timeout time.Duration) *GetCasWriteTimeoutsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cas write timeouts params
func (o *GetCasWriteTimeoutsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cas write timeouts params
func (o *GetCasWriteTimeoutsParams) WithContext(ctx context.Context) *GetCasWriteTimeoutsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cas write timeouts params
func (o *GetCasWriteTimeoutsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cas write timeouts params
func (o *GetCasWriteTimeoutsParams) WithHTTPClient(client *http.Client) *GetCasWriteTimeoutsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cas write timeouts params
func (o *GetCasWriteTimeoutsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCasWriteTimeoutsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
