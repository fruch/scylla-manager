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

// GetCreateHintCountReader is a Reader for the GetCreateHintCount structure.
type GetCreateHintCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCreateHintCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCreateHintCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCreateHintCountOK creates a GetCreateHintCountOK with default headers values
func NewGetCreateHintCountOK() *GetCreateHintCountOK {
	return &GetCreateHintCountOK{}
}

/*GetCreateHintCountOK handles this case with default header values.

GetCreateHintCountOK get create hint count o k
*/
type GetCreateHintCountOK struct {
	Payload int32
}

func (o *GetCreateHintCountOK) Error() string {
	return fmt.Sprintf("[GET /hinted_handoff/metrics/create_hint/{addr}][%d] getCreateHintCountOK  %+v", 200, o.Payload)
}

func (o *GetCreateHintCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
