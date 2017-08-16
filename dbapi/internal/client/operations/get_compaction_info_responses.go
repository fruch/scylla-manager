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

// GetCompactionInfoReader is a Reader for the GetCompactionInfo structure.
type GetCompactionInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCompactionInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCompactionInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCompactionInfoOK creates a GetCompactionInfoOK with default headers values
func NewGetCompactionInfoOK() *GetCompactionInfoOK {
	return &GetCompactionInfoOK{}
}

/*GetCompactionInfoOK handles this case with default header values.

GetCompactionInfoOK get compaction info o k
*/
type GetCompactionInfoOK struct {
	Payload []*models.CompactionInfo
}

func (o *GetCompactionInfoOK) Error() string {
	return fmt.Sprintf("[GET /compaction_manager/compaction_info][%d] getCompactionInfoOK  %+v", 200, o.Payload)
}

func (o *GetCompactionInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
