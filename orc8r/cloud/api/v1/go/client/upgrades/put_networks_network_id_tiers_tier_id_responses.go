// Code generated by go-swagger; DO NOT EDIT.

package upgrades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutNetworksNetworkIDTiersTierIDReader is a Reader for the PutNetworksNetworkIDTiersTierID structure.
type PutNetworksNetworkIDTiersTierIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDTiersTierIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDTiersTierIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDTiersTierIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDTiersTierIDNoContent creates a PutNetworksNetworkIDTiersTierIDNoContent with default headers values
func NewPutNetworksNetworkIDTiersTierIDNoContent() *PutNetworksNetworkIDTiersTierIDNoContent {
	return &PutNetworksNetworkIDTiersTierIDNoContent{}
}

/*PutNetworksNetworkIDTiersTierIDNoContent handles this case with default header values.

Success
*/
type PutNetworksNetworkIDTiersTierIDNoContent struct {
}

func (o *PutNetworksNetworkIDTiersTierIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/tiers/{tier_id}][%d] putNetworksNetworkIdTiersTierIdNoContent ", 204)
}

func (o *PutNetworksNetworkIDTiersTierIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDTiersTierIDDefault creates a PutNetworksNetworkIDTiersTierIDDefault with default headers values
func NewPutNetworksNetworkIDTiersTierIDDefault(code int) *PutNetworksNetworkIDTiersTierIDDefault {
	return &PutNetworksNetworkIDTiersTierIDDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDTiersTierIDDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDTiersTierIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID tiers tier ID default response
func (o *PutNetworksNetworkIDTiersTierIDDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDTiersTierIDDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/tiers/{tier_id}][%d] PutNetworksNetworkIDTiersTierID default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDTiersTierIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDTiersTierIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
