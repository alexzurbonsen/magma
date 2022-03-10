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

// GetNetworksNetworkIDTiersTierIDNameReader is a Reader for the GetNetworksNetworkIDTiersTierIDName structure.
type GetNetworksNetworkIDTiersTierIDNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDTiersTierIDNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDTiersTierIDNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDTiersTierIDNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDTiersTierIDNameOK creates a GetNetworksNetworkIDTiersTierIDNameOK with default headers values
func NewGetNetworksNetworkIDTiersTierIDNameOK() *GetNetworksNetworkIDTiersTierIDNameOK {
	return &GetNetworksNetworkIDTiersTierIDNameOK{}
}

/*GetNetworksNetworkIDTiersTierIDNameOK handles this case with default header values.

Success
*/
type GetNetworksNetworkIDTiersTierIDNameOK struct {
	Payload models.TierName
}

func (o *GetNetworksNetworkIDTiersTierIDNameOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/tiers/{tier_id}/name][%d] getNetworksNetworkIdTiersTierIdNameOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDTiersTierIDNameOK) GetPayload() models.TierName {
	return o.Payload
}

func (o *GetNetworksNetworkIDTiersTierIDNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDTiersTierIDNameDefault creates a GetNetworksNetworkIDTiersTierIDNameDefault with default headers values
func NewGetNetworksNetworkIDTiersTierIDNameDefault(code int) *GetNetworksNetworkIDTiersTierIDNameDefault {
	return &GetNetworksNetworkIDTiersTierIDNameDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDTiersTierIDNameDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDTiersTierIDNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID tiers tier ID name default response
func (o *GetNetworksNetworkIDTiersTierIDNameDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDTiersTierIDNameDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/tiers/{tier_id}/name][%d] GetNetworksNetworkIDTiersTierIDName default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDTiersTierIDNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDTiersTierIDNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
