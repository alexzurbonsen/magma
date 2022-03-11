// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetLTENetworkIDSMSSMSPkReader is a Reader for the GetLTENetworkIDSMSSMSPk structure.
type GetLTENetworkIDSMSSMSPkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDSMSSMSPkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDSMSSMSPkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDSMSSMSPkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDSMSSMSPkOK creates a GetLTENetworkIDSMSSMSPkOK with default headers values
func NewGetLTENetworkIDSMSSMSPkOK() *GetLTENetworkIDSMSSMSPkOK {
	return &GetLTENetworkIDSMSSMSPkOK{}
}

/* GetLTENetworkIDSMSSMSPkOK describes a response with status code 200, with default header values.

Requested SMS message
*/
type GetLTENetworkIDSMSSMSPkOK struct {
	Payload *models.SMSMessage
}

func (o *GetLTENetworkIDSMSSMSPkOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/sms/{sms_pk}][%d] getLteNetworkIdSmsSmsPkOK  %+v", 200, o.Payload)
}
func (o *GetLTENetworkIDSMSSMSPkOK) GetPayload() *models.SMSMessage {
	return o.Payload
}

func (o *GetLTENetworkIDSMSSMSPkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SMSMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDSMSSMSPkDefault creates a GetLTENetworkIDSMSSMSPkDefault with default headers values
func NewGetLTENetworkIDSMSSMSPkDefault(code int) *GetLTENetworkIDSMSSMSPkDefault {
	return &GetLTENetworkIDSMSSMSPkDefault{
		_statusCode: code,
	}
}

/* GetLTENetworkIDSMSSMSPkDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetLTENetworkIDSMSSMSPkDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID SMS SMS pk default response
func (o *GetLTENetworkIDSMSSMSPkDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDSMSSMSPkDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/sms/{sms_pk}][%d] GetLTENetworkIDSMSSMSPk default  %+v", o._statusCode, o.Payload)
}
func (o *GetLTENetworkIDSMSSMSPkDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDSMSSMSPkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
