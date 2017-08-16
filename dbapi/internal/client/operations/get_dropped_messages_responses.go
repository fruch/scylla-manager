// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/scylladb/mermaid/dbapi/internal/models"
)

// GetDroppedMessagesReader is a Reader for the GetDroppedMessages structure.
type GetDroppedMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDroppedMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDroppedMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDroppedMessagesOK creates a GetDroppedMessagesOK with default headers values
func NewGetDroppedMessagesOK() *GetDroppedMessagesOK {
	return &GetDroppedMessagesOK{}
}

/*GetDroppedMessagesOK handles this case with default header values.

GetDroppedMessagesOK get dropped messages o k
*/
type GetDroppedMessagesOK struct {
	Payload []*models.MessageCounter
}

func (o *GetDroppedMessagesOK) Error() string {
	return fmt.Sprintf("[GET /messaging_service/messages/dropped][%d] getDroppedMessagesOK  %+v", 200, o.Payload)
}

func (o *GetDroppedMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
