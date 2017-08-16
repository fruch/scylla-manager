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

// GetAllEndpointStatesReader is a Reader for the GetAllEndpointStates structure.
type GetAllEndpointStatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllEndpointStatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllEndpointStatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllEndpointStatesOK creates a GetAllEndpointStatesOK with default headers values
func NewGetAllEndpointStatesOK() *GetAllEndpointStatesOK {
	return &GetAllEndpointStatesOK{}
}

/*GetAllEndpointStatesOK handles this case with default header values.

GetAllEndpointStatesOK get all endpoint states o k
*/
type GetAllEndpointStatesOK struct {
	Payload []*models.EndpointState
}

func (o *GetAllEndpointStatesOK) Error() string {
	return fmt.Sprintf("[GET /failure_detector/endpoints/][%d] getAllEndpointStatesOK  %+v", 200, o.Payload)
}

func (o *GetAllEndpointStatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
