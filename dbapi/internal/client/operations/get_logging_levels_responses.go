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

// GetLoggingLevelsReader is a Reader for the GetLoggingLevels structure.
type GetLoggingLevelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoggingLevelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLoggingLevelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLoggingLevelsOK creates a GetLoggingLevelsOK with default headers values
func NewGetLoggingLevelsOK() *GetLoggingLevelsOK {
	return &GetLoggingLevelsOK{}
}

/*GetLoggingLevelsOK handles this case with default header values.

GetLoggingLevelsOK get logging levels o k
*/
type GetLoggingLevelsOK struct {
	Payload []*models.Mapper
}

func (o *GetLoggingLevelsOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/logging_level][%d] getLoggingLevelsOK  %+v", 200, o.Payload)
}

func (o *GetLoggingLevelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
