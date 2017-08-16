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

// GetDroppableTombstoneRatioReader is a Reader for the GetDroppableTombstoneRatio structure.
type GetDroppableTombstoneRatioReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDroppableTombstoneRatioReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDroppableTombstoneRatioOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDroppableTombstoneRatioOK creates a GetDroppableTombstoneRatioOK with default headers values
func NewGetDroppableTombstoneRatioOK() *GetDroppableTombstoneRatioOK {
	return &GetDroppableTombstoneRatioOK{}
}

/*GetDroppableTombstoneRatioOK handles this case with default header values.

GetDroppableTombstoneRatioOK get droppable tombstone ratio o k
*/
type GetDroppableTombstoneRatioOK struct {
	Payload int32
}

func (o *GetDroppableTombstoneRatioOK) Error() string {
	return fmt.Sprintf("[GET /column_family/droppable_ratio/{name}][%d] getDroppableTombstoneRatioOK  %+v", 200, o.Payload)
}

func (o *GetDroppableTombstoneRatioOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
