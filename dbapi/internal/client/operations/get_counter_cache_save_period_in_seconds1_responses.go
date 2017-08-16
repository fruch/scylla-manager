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

// GetCounterCacheSavePeriodInSeconds1Reader is a Reader for the GetCounterCacheSavePeriodInSeconds1 structure.
type GetCounterCacheSavePeriodInSeconds1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCounterCacheSavePeriodInSeconds1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCounterCacheSavePeriodInSeconds1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCounterCacheSavePeriodInSeconds1OK creates a GetCounterCacheSavePeriodInSeconds1OK with default headers values
func NewGetCounterCacheSavePeriodInSeconds1OK() *GetCounterCacheSavePeriodInSeconds1OK {
	return &GetCounterCacheSavePeriodInSeconds1OK{}
}

/*GetCounterCacheSavePeriodInSeconds1OK handles this case with default header values.

GetCounterCacheSavePeriodInSeconds1OK get counter cache save period in seconds1 o k
*/
type GetCounterCacheSavePeriodInSeconds1OK struct {
	Payload int32
}

func (o *GetCounterCacheSavePeriodInSeconds1OK) Error() string {
	return fmt.Sprintf("[GET /cache_service/counter_cache_save_period][%d] getCounterCacheSavePeriodInSeconds1OK  %+v", 200, o.Payload)
}

func (o *GetCounterCacheSavePeriodInSeconds1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
