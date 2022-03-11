// Code generated by go-swagger; DO NOT EDIT.

package enode_bs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetLTENetworkIDEnodebsENODEBSerialStateReader is a Reader for the GetLTENetworkIDEnodebsENODEBSerialState structure.
type GetLTENetworkIDEnodebsENODEBSerialStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDEnodebsENODEBSerialStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDEnodebsENODEBSerialStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDEnodebsENODEBSerialStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDEnodebsENODEBSerialStateOK creates a GetLTENetworkIDEnodebsENODEBSerialStateOK with default headers values
func NewGetLTENetworkIDEnodebsENODEBSerialStateOK() *GetLTENetworkIDEnodebsENODEBSerialStateOK {
	return &GetLTENetworkIDEnodebsENODEBSerialStateOK{}
}

/* GetLTENetworkIDEnodebsENODEBSerialStateOK describes a response with status code 200, with default header values.

The requested enodeB's configuration
*/
type GetLTENetworkIDEnodebsENODEBSerialStateOK struct {
	Payload *models.ENODEBState
}

func (o *GetLTENetworkIDEnodebsENODEBSerialStateOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/enodebs/{enodeb_serial}/state][%d] getLteNetworkIdEnodebsEnodebSerialStateOK  %+v", 200, o.Payload)
}
func (o *GetLTENetworkIDEnodebsENODEBSerialStateOK) GetPayload() *models.ENODEBState {
	return o.Payload
}

func (o *GetLTENetworkIDEnodebsENODEBSerialStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ENODEBState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDEnodebsENODEBSerialStateDefault creates a GetLTENetworkIDEnodebsENODEBSerialStateDefault with default headers values
func NewGetLTENetworkIDEnodebsENODEBSerialStateDefault(code int) *GetLTENetworkIDEnodebsENODEBSerialStateDefault {
	return &GetLTENetworkIDEnodebsENODEBSerialStateDefault{
		_statusCode: code,
	}
}

/* GetLTENetworkIDEnodebsENODEBSerialStateDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDEnodebsENODEBSerialStateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID enodebs ENODEB serial state default response
func (o *GetLTENetworkIDEnodebsENODEBSerialStateDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDEnodebsENODEBSerialStateDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/enodebs/{enodeb_serial}/state][%d] GetLTENetworkIDEnodebsENODEBSerialState default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDEnodebsENODEBSerialStateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDEnodebsENODEBSerialStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
