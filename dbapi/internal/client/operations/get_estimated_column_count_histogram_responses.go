// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetEstimatedColumnCountHistogramReader is a Reader for the GetEstimatedColumnCountHistogram structure.
type GetEstimatedColumnCountHistogramReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEstimatedColumnCountHistogramReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEstimatedColumnCountHistogramOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEstimatedColumnCountHistogramOK creates a GetEstimatedColumnCountHistogramOK with default headers values
func NewGetEstimatedColumnCountHistogramOK() *GetEstimatedColumnCountHistogramOK {
	return &GetEstimatedColumnCountHistogramOK{}
}

/*GetEstimatedColumnCountHistogramOK handles this case with default header values.

GetEstimatedColumnCountHistogramOK get estimated column count histogram o k
*/
type GetEstimatedColumnCountHistogramOK struct {
	Payload []int64
}

func (o *GetEstimatedColumnCountHistogramOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/estimated_column_count_histogram/{name}][%d] getEstimatedColumnCountHistogramOK  %+v", 200, o.Payload)
}

func (o *GetEstimatedColumnCountHistogramOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
