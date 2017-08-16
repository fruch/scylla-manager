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

// GetRowCapacity1Reader is a Reader for the GetRowCapacity1 structure.
type GetRowCapacity1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRowCapacity1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRowCapacity1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRowCapacity1OK creates a GetRowCapacity1OK with default headers values
func NewGetRowCapacity1OK() *GetRowCapacity1OK {
	return &GetRowCapacity1OK{}
}

/*GetRowCapacity1OK handles this case with default header values.

GetRowCapacity1OK get row capacity1 o k
*/
type GetRowCapacity1OK struct {
	Payload int64
}

func (o *GetRowCapacity1OK) Error() string {
	return fmt.Sprintf("[GET /cache_service/metrics/row/capacity][%d] getRowCapacity1OK  %+v", 200, o.Payload)
}

func (o *GetRowCapacity1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
