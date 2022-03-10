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

// GetNetworksNetworkIDTiersTierIDReader is a Reader for the GetNetworksNetworkIDTiersTierID structure.
type GetNetworksNetworkIDTiersTierIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDTiersTierIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDTiersTierIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDTiersTierIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDTiersTierIDOK creates a GetNetworksNetworkIDTiersTierIDOK with default headers values
func NewGetNetworksNetworkIDTiersTierIDOK() *GetNetworksNetworkIDTiersTierIDOK {
	return &GetNetworksNetworkIDTiersTierIDOK{}
}

/*GetNetworksNetworkIDTiersTierIDOK handles this case with default header values.

Success
*/
type GetNetworksNetworkIDTiersTierIDOK struct {
	Payload *models.Tier
}

func (o *GetNetworksNetworkIDTiersTierIDOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/tiers/{tier_id}][%d] getNetworksNetworkIdTiersTierIdOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDTiersTierIDOK) GetPayload() *models.Tier {
	return o.Payload
}

func (o *GetNetworksNetworkIDTiersTierIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDTiersTierIDDefault creates a GetNetworksNetworkIDTiersTierIDDefault with default headers values
func NewGetNetworksNetworkIDTiersTierIDDefault(code int) *GetNetworksNetworkIDTiersTierIDDefault {
	return &GetNetworksNetworkIDTiersTierIDDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDTiersTierIDDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDTiersTierIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID tiers tier ID default response
func (o *GetNetworksNetworkIDTiersTierIDDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDTiersTierIDDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/tiers/{tier_id}][%d] GetNetworksNetworkIDTiersTierID default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDTiersTierIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDTiersTierIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
