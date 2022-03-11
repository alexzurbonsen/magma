// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PostLTENetworkIDGatewayPoolsReader is a Reader for the PostLTENetworkIDGatewayPools structure.
type PostLTENetworkIDGatewayPoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLTENetworkIDGatewayPoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLTENetworkIDGatewayPoolsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostLTENetworkIDGatewayPoolsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostLTENetworkIDGatewayPoolsCreated creates a PostLTENetworkIDGatewayPoolsCreated with default headers values
func NewPostLTENetworkIDGatewayPoolsCreated() *PostLTENetworkIDGatewayPoolsCreated {
	return &PostLTENetworkIDGatewayPoolsCreated{}
}

/* PostLTENetworkIDGatewayPoolsCreated describes a response with status code 201, with default header values.

Success
*/
type PostLTENetworkIDGatewayPoolsCreated struct {
}

func (o *PostLTENetworkIDGatewayPoolsCreated) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/gateway_pools][%d] postLteNetworkIdGatewayPoolsCreated ", 201)
}

func (o *PostLTENetworkIDGatewayPoolsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLTENetworkIDGatewayPoolsDefault creates a PostLTENetworkIDGatewayPoolsDefault with default headers values
func NewPostLTENetworkIDGatewayPoolsDefault(code int) *PostLTENetworkIDGatewayPoolsDefault {
	return &PostLTENetworkIDGatewayPoolsDefault{
		_statusCode: code,
	}
}

/* PostLTENetworkIDGatewayPoolsDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PostLTENetworkIDGatewayPoolsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post LTE network ID gateway pools default response
func (o *PostLTENetworkIDGatewayPoolsDefault) Code() int {
	return o._statusCode
}

func (o *PostLTENetworkIDGatewayPoolsDefault) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/gateway_pools][%d] PostLTENetworkIDGatewayPools default  %+v", o._statusCode, o.Payload)
}
func (o *PostLTENetworkIDGatewayPoolsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostLTENetworkIDGatewayPoolsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
