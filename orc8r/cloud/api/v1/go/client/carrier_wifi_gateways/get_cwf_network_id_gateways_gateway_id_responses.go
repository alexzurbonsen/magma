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

// GetCwfNetworkIDGatewaysGatewayIDReader is a Reader for the GetCwfNetworkIDGatewaysGatewayID structure.
type GetCwfNetworkIDGatewaysGatewayIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCwfNetworkIDGatewaysGatewayIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCwfNetworkIDGatewaysGatewayIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCwfNetworkIDGatewaysGatewayIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDOK creates a GetCwfNetworkIDGatewaysGatewayIDOK with default headers values
func NewGetCwfNetworkIDGatewaysGatewayIDOK() *GetCwfNetworkIDGatewaysGatewayIDOK {
	return &GetCwfNetworkIDGatewaysGatewayIDOK{}
}

/* GetCwfNetworkIDGatewaysGatewayIDOK describes a response with status code 200, with default header values.

The requested carrier wifi gateway
*/
type GetCwfNetworkIDGatewaysGatewayIDOK struct {
	Payload *models.CwfGateway
}

func (o *GetCwfNetworkIDGatewaysGatewayIDOK) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/gateways/{gateway_id}][%d] getCwfNetworkIdGatewaysGatewayIdOK  %+v", 200, o.Payload)
}
func (o *GetCwfNetworkIDGatewaysGatewayIDOK) GetPayload() *models.CwfGateway {
	return o.Payload
}

func (o *GetCwfNetworkIDGatewaysGatewayIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CwfGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCwfNetworkIDGatewaysGatewayIDDefault creates a GetCwfNetworkIDGatewaysGatewayIDDefault with default headers values
func NewGetCwfNetworkIDGatewaysGatewayIDDefault(code int) *GetCwfNetworkIDGatewaysGatewayIDDefault {
	return &GetCwfNetworkIDGatewaysGatewayIDDefault{
		_statusCode: code,
	}
}

/* GetCwfNetworkIDGatewaysGatewayIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetCwfNetworkIDGatewaysGatewayIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cwf network ID gateways gateway ID default response
func (o *GetCwfNetworkIDGatewaysGatewayIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCwfNetworkIDGatewaysGatewayIDDefault) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/gateways/{gateway_id}][%d] GetCwfNetworkIDGatewaysGatewayID default  %+v", o._statusCode, o.Payload)
}
func (o *GetCwfNetworkIDGatewaysGatewayIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCwfNetworkIDGatewaysGatewayIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
