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

// GetLTENetworkIDCellularNgcReader is a Reader for the GetLTENetworkIDCellularNgc structure.
type GetLTENetworkIDCellularNgcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDCellularNgcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDCellularNgcOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDCellularNgcDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDCellularNgcOK creates a GetLTENetworkIDCellularNgcOK with default headers values
func NewGetLTENetworkIDCellularNgcOK() *GetLTENetworkIDCellularNgcOK {
	return &GetLTENetworkIDCellularNgcOK{}
}

/* GetLTENetworkIDCellularNgcOK describes a response with status code 200, with default header values.

NGC configuration of the network
*/
type GetLTENetworkIDCellularNgcOK struct {
	Payload *models.NetworkNgcConfigs
}

func (o *GetLTENetworkIDCellularNgcOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/cellular/ngc][%d] getLteNetworkIdCellularNgcOK  %+v", 200, o.Payload)
}
func (o *GetLTENetworkIDCellularNgcOK) GetPayload() *models.NetworkNgcConfigs {
	return o.Payload
}

func (o *GetLTENetworkIDCellularNgcOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkNgcConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDCellularNgcDefault creates a GetLTENetworkIDCellularNgcDefault with default headers values
func NewGetLTENetworkIDCellularNgcDefault(code int) *GetLTENetworkIDCellularNgcDefault {
	return &GetLTENetworkIDCellularNgcDefault{
		_statusCode: code,
	}
}

/* GetLTENetworkIDCellularNgcDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDCellularNgcDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID cellular ngc default response
func (o *GetLTENetworkIDCellularNgcDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDCellularNgcDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/cellular/ngc][%d] GetLTENetworkIDCellularNgc default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDCellularNgcDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDCellularNgcDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
