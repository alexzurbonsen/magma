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

// GetLTENetworkIDCellularRanReader is a Reader for the GetLTENetworkIDCellularRan structure.
type GetLTENetworkIDCellularRanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDCellularRanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDCellularRanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDCellularRanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDCellularRanOK creates a GetLTENetworkIDCellularRanOK with default headers values
func NewGetLTENetworkIDCellularRanOK() *GetLTENetworkIDCellularRanOK {
	return &GetLTENetworkIDCellularRanOK{}
}

/*GetLTENetworkIDCellularRanOK handles this case with default header values.

RAN configuration of the network
*/
type GetLTENetworkIDCellularRanOK struct {
	Payload *models.NetworkRanConfigs
}

func (o *GetLTENetworkIDCellularRanOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/cellular/ran][%d] getLteNetworkIdCellularRanOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDCellularRanOK) GetPayload() *models.NetworkRanConfigs {
	return o.Payload
}

func (o *GetLTENetworkIDCellularRanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkRanConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDCellularRanDefault creates a GetLTENetworkIDCellularRanDefault with default headers values
func NewGetLTENetworkIDCellularRanDefault(code int) *GetLTENetworkIDCellularRanDefault {
	return &GetLTENetworkIDCellularRanDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDCellularRanDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDCellularRanDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID cellular ran default response
func (o *GetLTENetworkIDCellularRanDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDCellularRanDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/cellular/ran][%d] GetLTENetworkIDCellularRan default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDCellularRanDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDCellularRanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
