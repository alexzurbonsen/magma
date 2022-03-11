// Code generated by go-swagger; DO NOT EDIT.

package federation_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutFegNetworkIDGatewaysGatewayIDFederationReader is a Reader for the PutFegNetworkIDGatewaysGatewayIDFederation structure.
type PutFegNetworkIDGatewaysGatewayIDFederationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFegNetworkIDGatewaysGatewayIDFederationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutFegNetworkIDGatewaysGatewayIDFederationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutFegNetworkIDGatewaysGatewayIDFederationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutFegNetworkIDGatewaysGatewayIDFederationOK creates a PutFegNetworkIDGatewaysGatewayIDFederationOK with default headers values
func NewPutFegNetworkIDGatewaysGatewayIDFederationOK() *PutFegNetworkIDGatewaysGatewayIDFederationOK {
	return &PutFegNetworkIDGatewaysGatewayIDFederationOK{}
}

/* PutFegNetworkIDGatewaysGatewayIDFederationOK describes a response with status code 200, with default header values.

Success
*/
type PutFegNetworkIDGatewaysGatewayIDFederationOK struct {
}

func (o *PutFegNetworkIDGatewaysGatewayIDFederationOK) Error() string {
	return fmt.Sprintf("[PUT /feg/{network_id}/gateways/{gateway_id}/federation][%d] putFegNetworkIdGatewaysGatewayIdFederationOK ", 200)
}

func (o *PutFegNetworkIDGatewaysGatewayIDFederationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFegNetworkIDGatewaysGatewayIDFederationDefault creates a PutFegNetworkIDGatewaysGatewayIDFederationDefault with default headers values
func NewPutFegNetworkIDGatewaysGatewayIDFederationDefault(code int) *PutFegNetworkIDGatewaysGatewayIDFederationDefault {
	return &PutFegNetworkIDGatewaysGatewayIDFederationDefault{
		_statusCode: code,
	}
}

/* PutFegNetworkIDGatewaysGatewayIDFederationDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutFegNetworkIDGatewaysGatewayIDFederationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put feg network ID gateways gateway ID federation default response
func (o *PutFegNetworkIDGatewaysGatewayIDFederationDefault) Code() int {
	return o._statusCode
}

func (o *PutFegNetworkIDGatewaysGatewayIDFederationDefault) Error() string {
	return fmt.Sprintf("[PUT /feg/{network_id}/gateways/{gateway_id}/federation][%d] PutFegNetworkIDGatewaysGatewayIDFederation default  %+v", o._statusCode, o.Payload)
}
func (o *PutFegNetworkIDGatewaysGatewayIDFederationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutFegNetworkIDGatewaysGatewayIDFederationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
