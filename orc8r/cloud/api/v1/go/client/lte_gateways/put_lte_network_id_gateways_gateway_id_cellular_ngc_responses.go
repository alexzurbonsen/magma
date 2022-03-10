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

// PutLTENetworkIDGatewaysGatewayIDCellularNgcReader is a Reader for the PutLTENetworkIDGatewaysGatewayIDCellularNgc structure.
type PutLTENetworkIDGatewaysGatewayIDCellularNgcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDGatewaysGatewayIDCellularNgcNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDGatewaysGatewayIDCellularNgcDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularNgcNoContent creates a PutLTENetworkIDGatewaysGatewayIDCellularNgcNoContent with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDCellularNgcNoContent() *PutLTENetworkIDGatewaysGatewayIDCellularNgcNoContent {
	return &PutLTENetworkIDGatewaysGatewayIDCellularNgcNoContent{}
}

/*PutLTENetworkIDGatewaysGatewayIDCellularNgcNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDGatewaysGatewayIDCellularNgcNoContent struct {
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/cellular/ngc][%d] putLteNetworkIdGatewaysGatewayIdCellularNgcNoContent ", 204)
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularNgcDefault creates a PutLTENetworkIDGatewaysGatewayIDCellularNgcDefault with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDCellularNgcDefault(code int) *PutLTENetworkIDGatewaysGatewayIDCellularNgcDefault {
	return &PutLTENetworkIDGatewaysGatewayIDCellularNgcDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDGatewaysGatewayIDCellularNgcDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDGatewaysGatewayIDCellularNgcDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID gateways gateway ID cellular ngc default response
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/cellular/ngc][%d] PutLTENetworkIDGatewaysGatewayIDCellularNgc default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
