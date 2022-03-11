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

// GetLTENetworkIDCellularReader is a Reader for the GetLTENetworkIDCellular structure.
type GetLTENetworkIDCellularReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDCellularReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDCellularOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDCellularDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDCellularOK creates a GetLTENetworkIDCellularOK with default headers values
func NewGetLTENetworkIDCellularOK() *GetLTENetworkIDCellularOK {
	return &GetLTENetworkIDCellularOK{}
}

/* GetLTENetworkIDCellularOK describes a response with status code 200, with default header values.

Cellular configuration of the network
*/
type GetLTENetworkIDCellularOK struct {
	Payload *models.NetworkCellularConfigs
}

func (o *GetLTENetworkIDCellularOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/cellular][%d] getLteNetworkIdCellularOK  %+v", 200, o.Payload)
}
func (o *GetLTENetworkIDCellularOK) GetPayload() *models.NetworkCellularConfigs {
	return o.Payload
}

func (o *GetLTENetworkIDCellularOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkCellularConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDCellularDefault creates a GetLTENetworkIDCellularDefault with default headers values
func NewGetLTENetworkIDCellularDefault(code int) *GetLTENetworkIDCellularDefault {
	return &GetLTENetworkIDCellularDefault{
		_statusCode: code,
	}
}

/* GetLTENetworkIDCellularDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDCellularDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID cellular default response
func (o *GetLTENetworkIDCellularDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDCellularDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/cellular][%d] GetLTENetworkIDCellular default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDCellularDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDCellularDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
