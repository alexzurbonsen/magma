// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutLTENetworkIDCellularRanReader is a Reader for the PutLTENetworkIDCellularRan structure.
type PutLTENetworkIDCellularRanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDCellularRanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDCellularRanNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDCellularRanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDCellularRanNoContent creates a PutLTENetworkIDCellularRanNoContent with default headers values
func NewPutLTENetworkIDCellularRanNoContent() *PutLTENetworkIDCellularRanNoContent {
	return &PutLTENetworkIDCellularRanNoContent{}
}

/* PutLTENetworkIDCellularRanNoContent describes a response with status code 204, with default header values.

Success
*/
type PutLTENetworkIDCellularRanNoContent struct {
}

func (o *PutLTENetworkIDCellularRanNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/cellular/ran][%d] putLteNetworkIdCellularRanNoContent ", 204)
}

func (o *PutLTENetworkIDCellularRanNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDCellularRanDefault creates a PutLTENetworkIDCellularRanDefault with default headers values
func NewPutLTENetworkIDCellularRanDefault(code int) *PutLTENetworkIDCellularRanDefault {
	return &PutLTENetworkIDCellularRanDefault{
		_statusCode: code,
	}
}

/* PutLTENetworkIDCellularRanDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutLTENetworkIDCellularRanDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID cellular ran default response
func (o *PutLTENetworkIDCellularRanDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDCellularRanDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/cellular/ran][%d] PutLTENetworkIDCellularRan default  %+v", o._statusCode, o.Payload)
}
func (o *PutLTENetworkIDCellularRanDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDCellularRanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
