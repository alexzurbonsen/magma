// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDStateReader is a Reader for the GetNetworksNetworkIDState structure.
type GetNetworksNetworkIDStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDStateOK creates a GetNetworksNetworkIDStateOK with default headers values
func NewGetNetworksNetworkIDStateOK() *GetNetworksNetworkIDStateOK {
	return &GetNetworksNetworkIDStateOK{}
}

/* GetNetworksNetworkIDStateOK describes a response with status code 200, with default header values.

State configuration
*/
type GetNetworksNetworkIDStateOK struct {
	Payload *models.StateConfig
}

func (o *GetNetworksNetworkIDStateOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/state][%d] getNetworksNetworkIdStateOK  %+v", 200, o.Payload)
}
func (o *GetNetworksNetworkIDStateOK) GetPayload() *models.StateConfig {
	return o.Payload
}

func (o *GetNetworksNetworkIDStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StateConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDStateDefault creates a GetNetworksNetworkIDStateDefault with default headers values
func NewGetNetworksNetworkIDStateDefault(code int) *GetNetworksNetworkIDStateDefault {
	return &GetNetworksNetworkIDStateDefault{
		_statusCode: code,
	}
}

/* GetNetworksNetworkIDStateDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDStateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID state default response
func (o *GetNetworksNetworkIDStateDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDStateDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/state][%d] GetNetworksNetworkIDState default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworksNetworkIDStateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
