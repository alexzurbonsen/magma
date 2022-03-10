// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDLogsCountReader is a Reader for the GetNetworksNetworkIDLogsCount structure.
type GetNetworksNetworkIDLogsCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDLogsCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDLogsCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDLogsCountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDLogsCountOK creates a GetNetworksNetworkIDLogsCountOK with default headers values
func NewGetNetworksNetworkIDLogsCountOK() *GetNetworksNetworkIDLogsCountOK {
	return &GetNetworksNetworkIDLogsCountOK{}
}

/*GetNetworksNetworkIDLogsCountOK handles this case with default header values.

Success
*/
type GetNetworksNetworkIDLogsCountOK struct {
	Payload models.ElasticHitCount
}

func (o *GetNetworksNetworkIDLogsCountOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/logs/count][%d] getNetworksNetworkIdLogsCountOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDLogsCountOK) GetPayload() models.ElasticHitCount {
	return o.Payload
}

func (o *GetNetworksNetworkIDLogsCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDLogsCountDefault creates a GetNetworksNetworkIDLogsCountDefault with default headers values
func NewGetNetworksNetworkIDLogsCountDefault(code int) *GetNetworksNetworkIDLogsCountDefault {
	return &GetNetworksNetworkIDLogsCountDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDLogsCountDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDLogsCountDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID logs count default response
func (o *GetNetworksNetworkIDLogsCountDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDLogsCountDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/logs/count][%d] GetNetworksNetworkIDLogsCount default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDLogsCountDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDLogsCountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
