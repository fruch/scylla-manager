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

// NewGetKeySize1Params creates a new GetKeySize1Params object
// with the default values initialized.
func NewGetKeySize1Params() *GetKeySize1Params {

	return &GetKeySize1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKeySize1ParamsWithTimeout creates a new GetKeySize1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKeySize1ParamsWithTimeout(timeout time.Duration) *GetKeySize1Params {

	return &GetKeySize1Params{

		timeout: timeout,
	}
}

// NewGetKeySize1ParamsWithContext creates a new GetKeySize1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetKeySize1ParamsWithContext(ctx context.Context) *GetKeySize1Params {

	return &GetKeySize1Params{

		Context: ctx,
	}
}

// NewGetKeySize1ParamsWithHTTPClient creates a new GetKeySize1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKeySize1ParamsWithHTTPClient(client *http.Client) *GetKeySize1Params {

	return &GetKeySize1Params{
		HTTPClient: client,
	}
}

/*GetKeySize1Params contains all the parameters to send to the API endpoint
for the get key size1 operation typically these are written to a http.Request
*/
type GetKeySize1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get key size1 params
func (o *GetKeySize1Params) WithTimeout(timeout time.Duration) *GetKeySize1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get key size1 params
func (o *GetKeySize1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get key size1 params
func (o *GetKeySize1Params) WithContext(ctx context.Context) *GetKeySize1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get key size1 params
func (o *GetKeySize1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get key size1 params
func (o *GetKeySize1Params) WithHTTPClient(client *http.Client) *GetKeySize1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get key size1 params
func (o *GetKeySize1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetKeySize1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
