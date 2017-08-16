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

// GetCompressionMetadataOffHeapMemoryUsedReader is a Reader for the GetCompressionMetadataOffHeapMemoryUsed structure.
type GetCompressionMetadataOffHeapMemoryUsedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCompressionMetadataOffHeapMemoryUsedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCompressionMetadataOffHeapMemoryUsedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCompressionMetadataOffHeapMemoryUsedOK creates a GetCompressionMetadataOffHeapMemoryUsedOK with default headers values
func NewGetCompressionMetadataOffHeapMemoryUsedOK() *GetCompressionMetadataOffHeapMemoryUsedOK {
	return &GetCompressionMetadataOffHeapMemoryUsedOK{}
}

/*GetCompressionMetadataOffHeapMemoryUsedOK handles this case with default header values.

GetCompressionMetadataOffHeapMemoryUsedOK get compression metadata off heap memory used o k
*/
type GetCompressionMetadataOffHeapMemoryUsedOK struct {
	Payload int64
}

func (o *GetCompressionMetadataOffHeapMemoryUsedOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/compression_metadata_off_heap_memory_used/{name}][%d] getCompressionMetadataOffHeapMemoryUsedOK  %+v", 200, o.Payload)
}

func (o *GetCompressionMetadataOffHeapMemoryUsedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
