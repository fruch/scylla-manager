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

// NewIsGossipRunningParams creates a new IsGossipRunningParams object
// with the default values initialized.
func NewIsGossipRunningParams() *IsGossipRunningParams {

	return &IsGossipRunningParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIsGossipRunningParamsWithTimeout creates a new IsGossipRunningParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIsGossipRunningParamsWithTimeout(timeout time.Duration) *IsGossipRunningParams {

	return &IsGossipRunningParams{

		timeout: timeout,
	}
}

// NewIsGossipRunningParamsWithContext creates a new IsGossipRunningParams object
// with the default values initialized, and the ability to set a context for a request
func NewIsGossipRunningParamsWithContext(ctx context.Context) *IsGossipRunningParams {

	return &IsGossipRunningParams{

		Context: ctx,
	}
}

// NewIsGossipRunningParamsWithHTTPClient creates a new IsGossipRunningParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIsGossipRunningParamsWithHTTPClient(client *http.Client) *IsGossipRunningParams {

	return &IsGossipRunningParams{
		HTTPClient: client,
	}
}

/*IsGossipRunningParams contains all the parameters to send to the API endpoint
for the is gossip running operation typically these are written to a http.Request
*/
type IsGossipRunningParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the is gossip running params
func (o *IsGossipRunningParams) WithTimeout(timeout time.Duration) *IsGossipRunningParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is gossip running params
func (o *IsGossipRunningParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is gossip running params
func (o *IsGossipRunningParams) WithContext(ctx context.Context) *IsGossipRunningParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is gossip running params
func (o *IsGossipRunningParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is gossip running params
func (o *IsGossipRunningParams) WithHTTPClient(client *http.Client) *IsGossipRunningParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is gossip running params
func (o *IsGossipRunningParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IsGossipRunningParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
