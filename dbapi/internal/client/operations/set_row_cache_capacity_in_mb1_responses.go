// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SetRowCacheCapacityInMb1Reader is a Reader for the SetRowCacheCapacityInMb1 structure.
type SetRowCacheCapacityInMb1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetRowCacheCapacityInMb1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetRowCacheCapacityInMb1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetRowCacheCapacityInMb1OK creates a SetRowCacheCapacityInMb1OK with default headers values
func NewSetRowCacheCapacityInMb1OK() *SetRowCacheCapacityInMb1OK {
	return &SetRowCacheCapacityInMb1OK{}
}

/*SetRowCacheCapacityInMb1OK handles this case with default header values.

SetRowCacheCapacityInMb1OK set row cache capacity in mb1 o k
*/
type SetRowCacheCapacityInMb1OK struct {
}

func (o *SetRowCacheCapacityInMb1OK) Error() string {
	return fmt.Sprintf("[POST /cache_service/row_cache_capacity][%d] setRowCacheCapacityInMb1OK ", 200)
}

func (o *SetRowCacheCapacityInMb1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
