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

// NewGetReadLatencyEstimatedHistogramParams creates a new GetReadLatencyEstimatedHistogramParams object
// with the default values initialized.
func NewGetReadLatencyEstimatedHistogramParams() *GetReadLatencyEstimatedHistogramParams {
	var ()
	return &GetReadLatencyEstimatedHistogramParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReadLatencyEstimatedHistogramParamsWithTimeout creates a new GetReadLatencyEstimatedHistogramParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReadLatencyEstimatedHistogramParamsWithTimeout(timeout time.Duration) *GetReadLatencyEstimatedHistogramParams {
	var ()
	return &GetReadLatencyEstimatedHistogramParams{

		timeout: timeout,
	}
}

// NewGetReadLatencyEstimatedHistogramParamsWithContext creates a new GetReadLatencyEstimatedHistogramParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReadLatencyEstimatedHistogramParamsWithContext(ctx context.Context) *GetReadLatencyEstimatedHistogramParams {
	var ()
	return &GetReadLatencyEstimatedHistogramParams{

		Context: ctx,
	}
}

// NewGetReadLatencyEstimatedHistogramParamsWithHTTPClient creates a new GetReadLatencyEstimatedHistogramParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReadLatencyEstimatedHistogramParamsWithHTTPClient(client *http.Client) *GetReadLatencyEstimatedHistogramParams {
	var ()
	return &GetReadLatencyEstimatedHistogramParams{
		HTTPClient: client,
	}
}

/*GetReadLatencyEstimatedHistogramParams contains all the parameters to send to the API endpoint
for the get read latency estimated histogram operation typically these are written to a http.Request
*/
type GetReadLatencyEstimatedHistogramParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get read latency estimated histogram params
func (o *GetReadLatencyEstimatedHistogramParams) WithTimeout(timeout time.Duration) *GetReadLatencyEstimatedHistogramParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get read latency estimated histogram params
func (o *GetReadLatencyEstimatedHistogramParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get read latency estimated histogram params
func (o *GetReadLatencyEstimatedHistogramParams) WithContext(ctx context.Context) *GetReadLatencyEstimatedHistogramParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get read latency estimated histogram params
func (o *GetReadLatencyEstimatedHistogramParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get read latency estimated histogram params
func (o *GetReadLatencyEstimatedHistogramParams) WithHTTPClient(client *http.Client) *GetReadLatencyEstimatedHistogramParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get read latency estimated histogram params
func (o *GetReadLatencyEstimatedHistogramParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get read latency estimated histogram params
func (o *GetReadLatencyEstimatedHistogramParams) WithName(name string) *GetReadLatencyEstimatedHistogramParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get read latency estimated histogram params
func (o *GetReadLatencyEstimatedHistogramParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetReadLatencyEstimatedHistogramParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
