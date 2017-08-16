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

// GetNotStoredHintsCountReader is a Reader for the GetNotStoredHintsCount structure.
type GetNotStoredHintsCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNotStoredHintsCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNotStoredHintsCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNotStoredHintsCountOK creates a GetNotStoredHintsCountOK with default headers values
func NewGetNotStoredHintsCountOK() *GetNotStoredHintsCountOK {
	return &GetNotStoredHintsCountOK{}
}

/*GetNotStoredHintsCountOK handles this case with default header values.

GetNotStoredHintsCountOK get not stored hints count o k
*/
type GetNotStoredHintsCountOK struct {
	Payload int32
}

func (o *GetNotStoredHintsCountOK) Error() string {
	return fmt.Sprintf("[GET /hinted_handoff/metrics/not_stored_hints/{addr}][%d] getNotStoredHintsCountOK  %+v", 200, o.Payload)
}

func (o *GetNotStoredHintsCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
