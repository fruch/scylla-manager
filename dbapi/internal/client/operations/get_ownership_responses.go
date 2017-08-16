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

// GetOwnershipReader is a Reader for the GetOwnership structure.
type GetOwnershipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOwnershipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOwnershipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOwnershipOK creates a GetOwnershipOK with default headers values
func NewGetOwnershipOK() *GetOwnershipOK {
	return &GetOwnershipOK{}
}

/*GetOwnershipOK handles this case with default header values.

GetOwnershipOK get ownership o k
*/
type GetOwnershipOK struct {
	Payload []*models.Mapper
}

func (o *GetOwnershipOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/ownership/][%d] getOwnershipOK  %+v", 200, o.Payload)
}

func (o *GetOwnershipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
