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

// GetStreamThroughputMbPerSecReader is a Reader for the GetStreamThroughputMbPerSec structure.
type GetStreamThroughputMbPerSecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStreamThroughputMbPerSecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStreamThroughputMbPerSecOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStreamThroughputMbPerSecOK creates a GetStreamThroughputMbPerSecOK with default headers values
func NewGetStreamThroughputMbPerSecOK() *GetStreamThroughputMbPerSecOK {
	return &GetStreamThroughputMbPerSecOK{}
}

/*GetStreamThroughputMbPerSecOK handles this case with default header values.

GetStreamThroughputMbPerSecOK get stream throughput mb per sec o k
*/
type GetStreamThroughputMbPerSecOK struct {
	Payload int32
}

func (o *GetStreamThroughputMbPerSecOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/stream_throughput][%d] getStreamThroughputMbPerSecOK  %+v", 200, o.Payload)
}

func (o *GetStreamThroughputMbPerSecOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
