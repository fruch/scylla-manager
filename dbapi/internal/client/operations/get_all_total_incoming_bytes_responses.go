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

// GetAllTotalIncomingBytesReader is a Reader for the GetAllTotalIncomingBytes structure.
type GetAllTotalIncomingBytesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTotalIncomingBytesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllTotalIncomingBytesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllTotalIncomingBytesOK creates a GetAllTotalIncomingBytesOK with default headers values
func NewGetAllTotalIncomingBytesOK() *GetAllTotalIncomingBytesOK {
	return &GetAllTotalIncomingBytesOK{}
}

/*GetAllTotalIncomingBytesOK handles this case with default header values.

GetAllTotalIncomingBytesOK get all total incoming bytes o k
*/
type GetAllTotalIncomingBytesOK struct {
	Payload int32
}

func (o *GetAllTotalIncomingBytesOK) Error() string {
	return fmt.Sprintf("[GET /stream_manager/metrics/incoming][%d] getAllTotalIncomingBytesOK  %+v", 200, o.Payload)
}

func (o *GetAllTotalIncomingBytesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
