// Code generated by go-swagger; DO NOT EDIT.

package cbsds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutDpNetworkIDCbsdsCbsdIDReader is a Reader for the PutDpNetworkIDCbsdsCbsdID structure.
type PutDpNetworkIDCbsdsCbsdIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDpNetworkIDCbsdsCbsdIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutDpNetworkIDCbsdsCbsdIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutDpNetworkIDCbsdsCbsdIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutDpNetworkIDCbsdsCbsdIDNoContent creates a PutDpNetworkIDCbsdsCbsdIDNoContent with default headers values
func NewPutDpNetworkIDCbsdsCbsdIDNoContent() *PutDpNetworkIDCbsdsCbsdIDNoContent {
	return &PutDpNetworkIDCbsdsCbsdIDNoContent{}
}

/* PutDpNetworkIDCbsdsCbsdIDNoContent describes a response with status code 204, with default header values.

Success
*/
type PutDpNetworkIDCbsdsCbsdIDNoContent struct {
}

func (o *PutDpNetworkIDCbsdsCbsdIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /dp/{network_id}/cbsds/{cbsd_id}][%d] putDpNetworkIdCbsdsCbsdIdNoContent ", 204)
}

func (o *PutDpNetworkIDCbsdsCbsdIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutDpNetworkIDCbsdsCbsdIDDefault creates a PutDpNetworkIDCbsdsCbsdIDDefault with default headers values
func NewPutDpNetworkIDCbsdsCbsdIDDefault(code int) *PutDpNetworkIDCbsdsCbsdIDDefault {
	return &PutDpNetworkIDCbsdsCbsdIDDefault{
		_statusCode: code,
	}
}

/* PutDpNetworkIDCbsdsCbsdIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutDpNetworkIDCbsdsCbsdIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put dp network ID cbsds cbsd ID default response
func (o *PutDpNetworkIDCbsdsCbsdIDDefault) Code() int {
	return o._statusCode
}

func (o *PutDpNetworkIDCbsdsCbsdIDDefault) Error() string {
	return fmt.Sprintf("[PUT /dp/{network_id}/cbsds/{cbsd_id}][%d] PutDpNetworkIDCbsdsCbsdID default  %+v", o._statusCode, o.Payload)
}
func (o *PutDpNetworkIDCbsdsCbsdIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDpNetworkIDCbsdsCbsdIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
