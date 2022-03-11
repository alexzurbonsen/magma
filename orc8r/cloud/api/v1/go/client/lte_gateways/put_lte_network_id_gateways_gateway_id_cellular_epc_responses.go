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

// PutLTENetworkIDGatewaysGatewayIDCellularEpcReader is a Reader for the PutLTENetworkIDGatewaysGatewayIDCellularEpc structure.
type PutLTENetworkIDGatewaysGatewayIDCellularEpcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDGatewaysGatewayIDCellularEpcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDGatewaysGatewayIDCellularEpcNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDGatewaysGatewayIDCellularEpcDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularEpcNoContent creates a PutLTENetworkIDGatewaysGatewayIDCellularEpcNoContent with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDCellularEpcNoContent() *PutLTENetworkIDGatewaysGatewayIDCellularEpcNoContent {
	return &PutLTENetworkIDGatewaysGatewayIDCellularEpcNoContent{}
}

/* PutLTENetworkIDGatewaysGatewayIDCellularEpcNoContent describes a response with status code 204, with default header values.

Success
*/
type PutLTENetworkIDGatewaysGatewayIDCellularEpcNoContent struct {
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularEpcNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/cellular/epc][%d] putLteNetworkIdGatewaysGatewayIdCellularEpcNoContent ", 204)
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularEpcNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularEpcDefault creates a PutLTENetworkIDGatewaysGatewayIDCellularEpcDefault with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDCellularEpcDefault(code int) *PutLTENetworkIDGatewaysGatewayIDCellularEpcDefault {
	return &PutLTENetworkIDGatewaysGatewayIDCellularEpcDefault{
		_statusCode: code,
	}
}

/* PutLTENetworkIDGatewaysGatewayIDCellularEpcDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutLTENetworkIDGatewaysGatewayIDCellularEpcDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID gateways gateway ID cellular epc default response
func (o *PutLTENetworkIDGatewaysGatewayIDCellularEpcDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularEpcDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/cellular/epc][%d] PutLTENetworkIDGatewaysGatewayIDCellularEpc default  %+v", o._statusCode, o.Payload)
}
func (o *PutLTENetworkIDGatewaysGatewayIDCellularEpcDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularEpcDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
