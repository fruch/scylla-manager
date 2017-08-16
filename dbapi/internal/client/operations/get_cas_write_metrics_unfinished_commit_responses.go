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

// GetCasWriteMetricsUnfinishedCommitReader is a Reader for the GetCasWriteMetricsUnfinishedCommit structure.
type GetCasWriteMetricsUnfinishedCommitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCasWriteMetricsUnfinishedCommitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCasWriteMetricsUnfinishedCommitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCasWriteMetricsUnfinishedCommitOK creates a GetCasWriteMetricsUnfinishedCommitOK with default headers values
func NewGetCasWriteMetricsUnfinishedCommitOK() *GetCasWriteMetricsUnfinishedCommitOK {
	return &GetCasWriteMetricsUnfinishedCommitOK{}
}

/*GetCasWriteMetricsUnfinishedCommitOK handles this case with default header values.

GetCasWriteMetricsUnfinishedCommitOK get cas write metrics unfinished commit o k
*/
type GetCasWriteMetricsUnfinishedCommitOK struct {
	Payload int32
}

func (o *GetCasWriteMetricsUnfinishedCommitOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/metrics/cas_write/unfinished_commit][%d] getCasWriteMetricsUnfinishedCommitOK  %+v", 200, o.Payload)
}

func (o *GetCasWriteMetricsUnfinishedCommitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
