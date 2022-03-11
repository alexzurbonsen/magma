// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetCwfNetworkIDGatewaysReader is a Reader for the GetCwfNetworkIDGateways structure.
type GetCwfNetworkIDGatewaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCwfNetworkIDGatewaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCwfNetworkIDGatewaysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCwfNetworkIDGatewaysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCwfNetworkIDGatewaysOK creates a GetCwfNetworkIDGatewaysOK with default headers values
func NewGetCwfNetworkIDGatewaysOK() *GetCwfNetworkIDGatewaysOK {
	return &GetCwfNetworkIDGatewaysOK{}
}

/* GetCwfNetworkIDGatewaysOK describes a response with status code 200, with default header values.

List of all carrier wifi gateways inside the network
*/
type GetCwfNetworkIDGatewaysOK struct {
	Payload map[string]models.CwfGateway
}

func (o *GetCwfNetworkIDGatewaysOK) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/gateways][%d] getCwfNetworkIdGatewaysOK  %+v", 200, o.Payload)
}
func (o *GetCwfNetworkIDGatewaysOK) GetPayload() map[string]models.CwfGateway {
	return o.Payload
}

func (o *GetCwfNetworkIDGatewaysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCwfNetworkIDGatewaysDefault creates a GetCwfNetworkIDGatewaysDefault with default headers values
func NewGetCwfNetworkIDGatewaysDefault(code int) *GetCwfNetworkIDGatewaysDefault {
	return &GetCwfNetworkIDGatewaysDefault{
		_statusCode: code,
	}
}

/* GetCwfNetworkIDGatewaysDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetCwfNetworkIDGatewaysDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cwf network ID gateways default response
func (o *GetCwfNetworkIDGatewaysDefault) Code() int {
	return o._statusCode
}

func (o *GetCwfNetworkIDGatewaysDefault) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/gateways][%d] GetCwfNetworkIDGateways default  %+v", o._statusCode, o.Payload)
}
func (o *GetCwfNetworkIDGatewaysDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCwfNetworkIDGatewaysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
