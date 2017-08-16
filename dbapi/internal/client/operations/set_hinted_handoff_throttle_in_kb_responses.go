// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SetHintedHandoffThrottleInKbReader is a Reader for the SetHintedHandoffThrottleInKb structure.
type SetHintedHandoffThrottleInKbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetHintedHandoffThrottleInKbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetHintedHandoffThrottleInKbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetHintedHandoffThrottleInKbOK creates a SetHintedHandoffThrottleInKbOK with default headers values
func NewSetHintedHandoffThrottleInKbOK() *SetHintedHandoffThrottleInKbOK {
	return &SetHintedHandoffThrottleInKbOK{}
}

/*SetHintedHandoffThrottleInKbOK handles this case with default header values.

SetHintedHandoffThrottleInKbOK set hinted handoff throttle in kb o k
*/
type SetHintedHandoffThrottleInKbOK struct {
}

func (o *SetHintedHandoffThrottleInKbOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/hinted_handoff][%d] setHintedHandoffThrottleInKbOK ", 200)
}

func (o *SetHintedHandoffThrottleInKbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
