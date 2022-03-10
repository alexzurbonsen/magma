// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetLTENetworkIDGatewaysGatewayIDReader is a Reader for the GetLTENetworkIDGatewaysGatewayID structure.
type GetLTENetworkIDGatewaysGatewayIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewaysGatewayIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDGatewaysGatewayIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewaysGatewayIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDOK creates a GetLTENetworkIDGatewaysGatewayIDOK with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDOK() *GetLTENetworkIDGatewaysGatewayIDOK {
	return &GetLTENetworkIDGatewaysGatewayIDOK{}
}

/*GetLTENetworkIDGatewaysGatewayIDOK handles this case with default header values.

The requested LTE gateway
*/
type GetLTENetworkIDGatewaysGatewayIDOK struct {
	Payload *models.LTEGateway
}

func (o *GetLTENetworkIDGatewaysGatewayIDOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}][%d] getLteNetworkIdGatewaysGatewayIdOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDOK) GetPayload() *models.LTEGateway {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LTEGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewaysGatewayIDDefault creates a GetLTENetworkIDGatewaysGatewayIDDefault with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDDefault(code int) *GetLTENetworkIDGatewaysGatewayIDDefault {
	return &GetLTENetworkIDGatewaysGatewayIDDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDGatewaysGatewayIDDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewaysGatewayIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateways gateway ID default response
func (o *GetLTENetworkIDGatewaysGatewayIDDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewaysGatewayIDDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}][%d] GetLTENetworkIDGatewaysGatewayID default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
