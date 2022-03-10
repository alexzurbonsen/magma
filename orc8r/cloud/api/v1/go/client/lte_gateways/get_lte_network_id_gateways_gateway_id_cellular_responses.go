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

// GetLTENetworkIDGatewaysGatewayIDCellularReader is a Reader for the GetLTENetworkIDGatewaysGatewayIDCellular structure.
type GetLTENetworkIDGatewaysGatewayIDCellularReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewaysGatewayIDCellularReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewGetLTENetworkIDGatewaysGatewayIDCellularNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewaysGatewayIDCellularDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularNoContent creates a GetLTENetworkIDGatewaysGatewayIDCellularNoContent with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDCellularNoContent() *GetLTENetworkIDGatewaysGatewayIDCellularNoContent {
	return &GetLTENetworkIDGatewaysGatewayIDCellularNoContent{}
}

/*GetLTENetworkIDGatewaysGatewayIDCellularNoContent handles this case with default header values.

Cellular configuration
*/
type GetLTENetworkIDGatewaysGatewayIDCellularNoContent struct {
	Payload *models.GatewayCellularConfigs
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularNoContent) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/cellular][%d] getLteNetworkIdGatewaysGatewayIdCellularNoContent  %+v", 204, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularNoContent) GetPayload() *models.GatewayCellularConfigs {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayCellularConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularDefault creates a GetLTENetworkIDGatewaysGatewayIDCellularDefault with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDCellularDefault(code int) *GetLTENetworkIDGatewaysGatewayIDCellularDefault {
	return &GetLTENetworkIDGatewaysGatewayIDCellularDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDGatewaysGatewayIDCellularDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewaysGatewayIDCellularDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateways gateway ID cellular default response
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/cellular][%d] GetLTENetworkIDGatewaysGatewayIDCellular default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
