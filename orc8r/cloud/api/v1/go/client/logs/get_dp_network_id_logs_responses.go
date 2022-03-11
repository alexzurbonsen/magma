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

// GetDpNetworkIDLogsReader is a Reader for the GetDpNetworkIDLogs structure.
type GetDpNetworkIDLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDpNetworkIDLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDpNetworkIDLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDpNetworkIDLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDpNetworkIDLogsOK creates a GetDpNetworkIDLogsOK with default headers values
func NewGetDpNetworkIDLogsOK() *GetDpNetworkIDLogsOK {
	return &GetDpNetworkIDLogsOK{}
}

/* GetDpNetworkIDLogsOK describes a response with status code 200, with default header values.

Messages between DP and SAS
*/
type GetDpNetworkIDLogsOK struct {
	Payload []*models.Message
}

func (o *GetDpNetworkIDLogsOK) Error() string {
	return fmt.Sprintf("[GET /dp/{network_id}/logs][%d] getDpNetworkIdLogsOK  %+v", 200, o.Payload)
}
func (o *GetDpNetworkIDLogsOK) GetPayload() []*models.Message {
	return o.Payload
}

func (o *GetDpNetworkIDLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDpNetworkIDLogsDefault creates a GetDpNetworkIDLogsDefault with default headers values
func NewGetDpNetworkIDLogsDefault(code int) *GetDpNetworkIDLogsDefault {
	return &GetDpNetworkIDLogsDefault{
		_statusCode: code,
	}
}

/* GetDpNetworkIDLogsDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetDpNetworkIDLogsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get dp network ID logs default response
func (o *GetDpNetworkIDLogsDefault) Code() int {
	return o._statusCode
}

func (o *GetDpNetworkIDLogsDefault) Error() string {
	return fmt.Sprintf("[GET /dp/{network_id}/logs][%d] GetDpNetworkIDLogs default  %+v", o._statusCode, o.Payload)
}
func (o *GetDpNetworkIDLogsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDpNetworkIDLogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
