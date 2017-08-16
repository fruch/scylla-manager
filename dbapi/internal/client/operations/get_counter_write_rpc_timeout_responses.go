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

// GetCounterWriteRPCTimeoutReader is a Reader for the GetCounterWriteRPCTimeout structure.
type GetCounterWriteRPCTimeoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCounterWriteRPCTimeoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCounterWriteRPCTimeoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCounterWriteRPCTimeoutOK creates a GetCounterWriteRPCTimeoutOK with default headers values
func NewGetCounterWriteRPCTimeoutOK() *GetCounterWriteRPCTimeoutOK {
	return &GetCounterWriteRPCTimeoutOK{}
}

/*GetCounterWriteRPCTimeoutOK handles this case with default header values.

GetCounterWriteRPCTimeoutOK get counter write Rpc timeout o k
*/
type GetCounterWriteRPCTimeoutOK struct {
	Payload int32
}

func (o *GetCounterWriteRPCTimeoutOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/counter_write_rpc_timeout][%d] getCounterWriteRpcTimeoutOK  %+v", 200, o.Payload)
}

func (o *GetCounterWriteRPCTimeoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
