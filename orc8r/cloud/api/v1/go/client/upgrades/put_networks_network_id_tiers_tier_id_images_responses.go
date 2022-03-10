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

// PutNetworksNetworkIDTiersTierIDImagesReader is a Reader for the PutNetworksNetworkIDTiersTierIDImages structure.
type PutNetworksNetworkIDTiersTierIDImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDTiersTierIDImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDTiersTierIDImagesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDTiersTierIDImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDTiersTierIDImagesNoContent creates a PutNetworksNetworkIDTiersTierIDImagesNoContent with default headers values
func NewPutNetworksNetworkIDTiersTierIDImagesNoContent() *PutNetworksNetworkIDTiersTierIDImagesNoContent {
	return &PutNetworksNetworkIDTiersTierIDImagesNoContent{}
}

/*PutNetworksNetworkIDTiersTierIDImagesNoContent handles this case with default header values.

Success
*/
type PutNetworksNetworkIDTiersTierIDImagesNoContent struct {
}

func (o *PutNetworksNetworkIDTiersTierIDImagesNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/tiers/{tier_id}/images][%d] putNetworksNetworkIdTiersTierIdImagesNoContent ", 204)
}

func (o *PutNetworksNetworkIDTiersTierIDImagesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDTiersTierIDImagesDefault creates a PutNetworksNetworkIDTiersTierIDImagesDefault with default headers values
func NewPutNetworksNetworkIDTiersTierIDImagesDefault(code int) *PutNetworksNetworkIDTiersTierIDImagesDefault {
	return &PutNetworksNetworkIDTiersTierIDImagesDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDTiersTierIDImagesDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDTiersTierIDImagesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID tiers tier ID images default response
func (o *PutNetworksNetworkIDTiersTierIDImagesDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDTiersTierIDImagesDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/tiers/{tier_id}/images][%d] PutNetworksNetworkIDTiersTierIDImages default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDTiersTierIDImagesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDTiersTierIDImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
