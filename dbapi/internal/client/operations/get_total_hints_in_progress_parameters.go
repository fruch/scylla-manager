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

// NewGetTotalHintsInProgressParams creates a new GetTotalHintsInProgressParams object
// with the default values initialized.
func NewGetTotalHintsInProgressParams() *GetTotalHintsInProgressParams {

	return &GetTotalHintsInProgressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTotalHintsInProgressParamsWithTimeout creates a new GetTotalHintsInProgressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTotalHintsInProgressParamsWithTimeout(timeout time.Duration) *GetTotalHintsInProgressParams {

	return &GetTotalHintsInProgressParams{

		timeout: timeout,
	}
}

// NewGetTotalHintsInProgressParamsWithContext creates a new GetTotalHintsInProgressParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTotalHintsInProgressParamsWithContext(ctx context.Context) *GetTotalHintsInProgressParams {

	return &GetTotalHintsInProgressParams{

		Context: ctx,
	}
}

// NewGetTotalHintsInProgressParamsWithHTTPClient creates a new GetTotalHintsInProgressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTotalHintsInProgressParamsWithHTTPClient(client *http.Client) *GetTotalHintsInProgressParams {

	return &GetTotalHintsInProgressParams{
		HTTPClient: client,
	}
}

/*GetTotalHintsInProgressParams contains all the parameters to send to the API endpoint
for the get total hints in progress operation typically these are written to a http.Request
*/
type GetTotalHintsInProgressParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get total hints in progress params
func (o *GetTotalHintsInProgressParams) WithTimeout(timeout time.Duration) *GetTotalHintsInProgressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get total hints in progress params
func (o *GetTotalHintsInProgressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get total hints in progress params
func (o *GetTotalHintsInProgressParams) WithContext(ctx context.Context) *GetTotalHintsInProgressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get total hints in progress params
func (o *GetTotalHintsInProgressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get total hints in progress params
func (o *GetTotalHintsInProgressParams) WithHTTPClient(client *http.Client) *GetTotalHintsInProgressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get total hints in progress params
func (o *GetTotalHintsInProgressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTotalHintsInProgressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
