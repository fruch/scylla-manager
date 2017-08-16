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

// NewForceRemoveCompletionParams creates a new ForceRemoveCompletionParams object
// with the default values initialized.
func NewForceRemoveCompletionParams() *ForceRemoveCompletionParams {

	return &ForceRemoveCompletionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewForceRemoveCompletionParamsWithTimeout creates a new ForceRemoveCompletionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewForceRemoveCompletionParamsWithTimeout(timeout time.Duration) *ForceRemoveCompletionParams {

	return &ForceRemoveCompletionParams{

		timeout: timeout,
	}
}

// NewForceRemoveCompletionParamsWithContext creates a new ForceRemoveCompletionParams object
// with the default values initialized, and the ability to set a context for a request
func NewForceRemoveCompletionParamsWithContext(ctx context.Context) *ForceRemoveCompletionParams {

	return &ForceRemoveCompletionParams{

		Context: ctx,
	}
}

// NewForceRemoveCompletionParamsWithHTTPClient creates a new ForceRemoveCompletionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewForceRemoveCompletionParamsWithHTTPClient(client *http.Client) *ForceRemoveCompletionParams {

	return &ForceRemoveCompletionParams{
		HTTPClient: client,
	}
}

/*ForceRemoveCompletionParams contains all the parameters to send to the API endpoint
for the force remove completion operation typically these are written to a http.Request
*/
type ForceRemoveCompletionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the force remove completion params
func (o *ForceRemoveCompletionParams) WithTimeout(timeout time.Duration) *ForceRemoveCompletionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the force remove completion params
func (o *ForceRemoveCompletionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the force remove completion params
func (o *ForceRemoveCompletionParams) WithContext(ctx context.Context) *ForceRemoveCompletionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the force remove completion params
func (o *ForceRemoveCompletionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the force remove completion params
func (o *ForceRemoveCompletionParams) WithHTTPClient(client *http.Client) *ForceRemoveCompletionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the force remove completion params
func (o *ForceRemoveCompletionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ForceRemoveCompletionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
