// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutCwfNetworkIDHaPairsHaPairIDReader is a Reader for the PutCwfNetworkIDHaPairsHaPairID structure.
type PutCwfNetworkIDHaPairsHaPairIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCwfNetworkIDHaPairsHaPairIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCwfNetworkIDHaPairsHaPairIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCwfNetworkIDHaPairsHaPairIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCwfNetworkIDHaPairsHaPairIDOK creates a PutCwfNetworkIDHaPairsHaPairIDOK with default headers values
func NewPutCwfNetworkIDHaPairsHaPairIDOK() *PutCwfNetworkIDHaPairsHaPairIDOK {
	return &PutCwfNetworkIDHaPairsHaPairIDOK{}
}

/* PutCwfNetworkIDHaPairsHaPairIDOK describes a response with status code 200, with default header values.

Success
*/
type PutCwfNetworkIDHaPairsHaPairIDOK struct {
}

func (o *PutCwfNetworkIDHaPairsHaPairIDOK) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/ha_pairs/{ha_pair_id}][%d] putCwfNetworkIdHaPairsHaPairIdOK ", 200)
}

func (o *PutCwfNetworkIDHaPairsHaPairIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCwfNetworkIDHaPairsHaPairIDDefault creates a PutCwfNetworkIDHaPairsHaPairIDDefault with default headers values
func NewPutCwfNetworkIDHaPairsHaPairIDDefault(code int) *PutCwfNetworkIDHaPairsHaPairIDDefault {
	return &PutCwfNetworkIDHaPairsHaPairIDDefault{
		_statusCode: code,
	}
}

/* PutCwfNetworkIDHaPairsHaPairIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutCwfNetworkIDHaPairsHaPairIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put cwf network ID ha pairs ha pair ID default response
func (o *PutCwfNetworkIDHaPairsHaPairIDDefault) Code() int {
	return o._statusCode
}

func (o *PutCwfNetworkIDHaPairsHaPairIDDefault) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/ha_pairs/{ha_pair_id}][%d] PutCwfNetworkIDHaPairsHaPairID default  %+v", o._statusCode, o.Payload)
}
func (o *PutCwfNetworkIDHaPairsHaPairIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCwfNetworkIDHaPairsHaPairIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
