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

// SetMaximumCompactionThresholdReader is a Reader for the SetMaximumCompactionThreshold structure.
type SetMaximumCompactionThresholdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetMaximumCompactionThresholdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetMaximumCompactionThresholdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetMaximumCompactionThresholdOK creates a SetMaximumCompactionThresholdOK with default headers values
func NewSetMaximumCompactionThresholdOK() *SetMaximumCompactionThresholdOK {
	return &SetMaximumCompactionThresholdOK{}
}

/*SetMaximumCompactionThresholdOK handles this case with default header values.

SetMaximumCompactionThresholdOK set maximum compaction threshold o k
*/
type SetMaximumCompactionThresholdOK struct {
	Payload string
}

func (o *SetMaximumCompactionThresholdOK) Error() string {
	return fmt.Sprintf("[POST /column_family/maximum_compaction/{name}][%d] setMaximumCompactionThresholdOK  %+v", 200, o.Payload)
}

func (o *SetMaximumCompactionThresholdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
