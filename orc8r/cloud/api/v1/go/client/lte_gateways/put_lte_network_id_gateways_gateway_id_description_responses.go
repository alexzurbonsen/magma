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

// PutLTENetworkIDGatewaysGatewayIDDescriptionReader is a Reader for the PutLTENetworkIDGatewaysGatewayIDDescription structure.
type PutLTENetworkIDGatewaysGatewayIDDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDGatewaysGatewayIDDescriptionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDGatewaysGatewayIDDescriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDDescriptionNoContent creates a PutLTENetworkIDGatewaysGatewayIDDescriptionNoContent with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDDescriptionNoContent() *PutLTENetworkIDGatewaysGatewayIDDescriptionNoContent {
	return &PutLTENetworkIDGatewaysGatewayIDDescriptionNoContent{}
}

/*PutLTENetworkIDGatewaysGatewayIDDescriptionNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDGatewaysGatewayIDDescriptionNoContent struct {
}

func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/description][%d] putLteNetworkIdGatewaysGatewayIdDescriptionNoContent ", 204)
}

func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDGatewaysGatewayIDDescriptionDefault creates a PutLTENetworkIDGatewaysGatewayIDDescriptionDefault with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDDescriptionDefault(code int) *PutLTENetworkIDGatewaysGatewayIDDescriptionDefault {
	return &PutLTENetworkIDGatewaysGatewayIDDescriptionDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDGatewaysGatewayIDDescriptionDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDGatewaysGatewayIDDescriptionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID gateways gateway ID description default response
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/description][%d] PutLTENetworkIDGatewaysGatewayIDDescription default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
