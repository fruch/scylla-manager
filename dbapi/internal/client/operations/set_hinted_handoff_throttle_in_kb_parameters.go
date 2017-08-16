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

// NewSetHintedHandoffThrottleInKbParams creates a new SetHintedHandoffThrottleInKbParams object
// with the default values initialized.
func NewSetHintedHandoffThrottleInKbParams() *SetHintedHandoffThrottleInKbParams {
	var ()
	return &SetHintedHandoffThrottleInKbParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetHintedHandoffThrottleInKbParamsWithTimeout creates a new SetHintedHandoffThrottleInKbParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetHintedHandoffThrottleInKbParamsWithTimeout(timeout time.Duration) *SetHintedHandoffThrottleInKbParams {
	var ()
	return &SetHintedHandoffThrottleInKbParams{

		timeout: timeout,
	}
}

// NewSetHintedHandoffThrottleInKbParamsWithContext creates a new SetHintedHandoffThrottleInKbParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetHintedHandoffThrottleInKbParamsWithContext(ctx context.Context) *SetHintedHandoffThrottleInKbParams {
	var ()
	return &SetHintedHandoffThrottleInKbParams{

		Context: ctx,
	}
}

// NewSetHintedHandoffThrottleInKbParamsWithHTTPClient creates a new SetHintedHandoffThrottleInKbParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetHintedHandoffThrottleInKbParamsWithHTTPClient(client *http.Client) *SetHintedHandoffThrottleInKbParams {
	var ()
	return &SetHintedHandoffThrottleInKbParams{
		HTTPClient: client,
	}
}

/*SetHintedHandoffThrottleInKbParams contains all the parameters to send to the API endpoint
for the set hinted handoff throttle in kb operation typically these are written to a http.Request
*/
type SetHintedHandoffThrottleInKbParams struct {

	/*Throttle
	  throttle in kb

	*/
	Throttle int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set hinted handoff throttle in kb params
func (o *SetHintedHandoffThrottleInKbParams) WithTimeout(timeout time.Duration) *SetHintedHandoffThrottleInKbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set hinted handoff throttle in kb params
func (o *SetHintedHandoffThrottleInKbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set hinted handoff throttle in kb params
func (o *SetHintedHandoffThrottleInKbParams) WithContext(ctx context.Context) *SetHintedHandoffThrottleInKbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set hinted handoff throttle in kb params
func (o *SetHintedHandoffThrottleInKbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set hinted handoff throttle in kb params
func (o *SetHintedHandoffThrottleInKbParams) WithHTTPClient(client *http.Client) *SetHintedHandoffThrottleInKbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set hinted handoff throttle in kb params
func (o *SetHintedHandoffThrottleInKbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithThrottle adds the throttle to the set hinted handoff throttle in kb params
func (o *SetHintedHandoffThrottleInKbParams) WithThrottle(throttle int32) *SetHintedHandoffThrottleInKbParams {
	o.SetThrottle(throttle)
	return o
}

// SetThrottle adds the throttle to the set hinted handoff throttle in kb params
func (o *SetHintedHandoffThrottleInKbParams) SetThrottle(throttle int32) {
	o.Throttle = throttle
}

// WriteToRequest writes these params to a swagger request
func (o *SetHintedHandoffThrottleInKbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param throttle
	qrThrottle := o.Throttle
	qThrottle := swag.FormatInt32(qrThrottle)
	if qThrottle != "" {
		if err := r.SetQueryParam("throttle", qThrottle); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
