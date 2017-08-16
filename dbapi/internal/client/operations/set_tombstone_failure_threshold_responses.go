// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SetTombstoneFailureThresholdReader is a Reader for the SetTombstoneFailureThreshold structure.
type SetTombstoneFailureThresholdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetTombstoneFailureThresholdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetTombstoneFailureThresholdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetTombstoneFailureThresholdOK creates a SetTombstoneFailureThresholdOK with default headers values
func NewSetTombstoneFailureThresholdOK() *SetTombstoneFailureThresholdOK {
	return &SetTombstoneFailureThresholdOK{}
}

/*SetTombstoneFailureThresholdOK handles this case with default header values.

SetTombstoneFailureThresholdOK set tombstone failure threshold o k
*/
type SetTombstoneFailureThresholdOK struct {
}

func (o *SetTombstoneFailureThresholdOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/tombstone_failure_threshold][%d] setTombstoneFailureThresholdOK ", 200)
}

func (o *SetTombstoneFailureThresholdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
