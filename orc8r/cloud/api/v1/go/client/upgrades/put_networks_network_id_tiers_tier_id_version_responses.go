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

// PutNetworksNetworkIDTiersTierIDVersionReader is a Reader for the PutNetworksNetworkIDTiersTierIDVersion structure.
type PutNetworksNetworkIDTiersTierIDVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDTiersTierIDVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDTiersTierIDVersionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDTiersTierIDVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDTiersTierIDVersionNoContent creates a PutNetworksNetworkIDTiersTierIDVersionNoContent with default headers values
func NewPutNetworksNetworkIDTiersTierIDVersionNoContent() *PutNetworksNetworkIDTiersTierIDVersionNoContent {
	return &PutNetworksNetworkIDTiersTierIDVersionNoContent{}
}

/*PutNetworksNetworkIDTiersTierIDVersionNoContent handles this case with default header values.

Success
*/
type PutNetworksNetworkIDTiersTierIDVersionNoContent struct {
}

func (o *PutNetworksNetworkIDTiersTierIDVersionNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/tiers/{tier_id}/version][%d] putNetworksNetworkIdTiersTierIdVersionNoContent ", 204)
}

func (o *PutNetworksNetworkIDTiersTierIDVersionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDTiersTierIDVersionDefault creates a PutNetworksNetworkIDTiersTierIDVersionDefault with default headers values
func NewPutNetworksNetworkIDTiersTierIDVersionDefault(code int) *PutNetworksNetworkIDTiersTierIDVersionDefault {
	return &PutNetworksNetworkIDTiersTierIDVersionDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDTiersTierIDVersionDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDTiersTierIDVersionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID tiers tier ID version default response
func (o *PutNetworksNetworkIDTiersTierIDVersionDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDTiersTierIDVersionDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/tiers/{tier_id}/version][%d] PutNetworksNetworkIDTiersTierIDVersion default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDTiersTierIDVersionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDTiersTierIDVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
