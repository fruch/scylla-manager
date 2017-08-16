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

// NewGetRowSize1Params creates a new GetRowSize1Params object
// with the default values initialized.
func NewGetRowSize1Params() *GetRowSize1Params {

	return &GetRowSize1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRowSize1ParamsWithTimeout creates a new GetRowSize1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRowSize1ParamsWithTimeout(timeout time.Duration) *GetRowSize1Params {

	return &GetRowSize1Params{

		timeout: timeout,
	}
}

// NewGetRowSize1ParamsWithContext creates a new GetRowSize1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetRowSize1ParamsWithContext(ctx context.Context) *GetRowSize1Params {

	return &GetRowSize1Params{

		Context: ctx,
	}
}

// NewGetRowSize1ParamsWithHTTPClient creates a new GetRowSize1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRowSize1ParamsWithHTTPClient(client *http.Client) *GetRowSize1Params {

	return &GetRowSize1Params{
		HTTPClient: client,
	}
}

/*GetRowSize1Params contains all the parameters to send to the API endpoint
for the get row size1 operation typically these are written to a http.Request
*/
type GetRowSize1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get row size1 params
func (o *GetRowSize1Params) WithTimeout(timeout time.Duration) *GetRowSize1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get row size1 params
func (o *GetRowSize1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get row size1 params
func (o *GetRowSize1Params) WithContext(ctx context.Context) *GetRowSize1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get row size1 params
func (o *GetRowSize1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get row size1 params
func (o *GetRowSize1Params) WithHTTPClient(client *http.Client) *GetRowSize1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get row size1 params
func (o *GetRowSize1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRowSize1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
