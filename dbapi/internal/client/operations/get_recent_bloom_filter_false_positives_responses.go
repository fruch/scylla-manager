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

// GetRecentBloomFilterFalsePositivesReader is a Reader for the GetRecentBloomFilterFalsePositives structure.
type GetRecentBloomFilterFalsePositivesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecentBloomFilterFalsePositivesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRecentBloomFilterFalsePositivesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRecentBloomFilterFalsePositivesOK creates a GetRecentBloomFilterFalsePositivesOK with default headers values
func NewGetRecentBloomFilterFalsePositivesOK() *GetRecentBloomFilterFalsePositivesOK {
	return &GetRecentBloomFilterFalsePositivesOK{}
}

/*GetRecentBloomFilterFalsePositivesOK handles this case with default header values.

GetRecentBloomFilterFalsePositivesOK get recent bloom filter false positives o k
*/
type GetRecentBloomFilterFalsePositivesOK struct {
	Payload int64
}

func (o *GetRecentBloomFilterFalsePositivesOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/recent_bloom_filter_false_positives/{name}][%d] getRecentBloomFilterFalsePositivesOK  %+v", 200, o.Payload)
}

func (o *GetRecentBloomFilterFalsePositivesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
