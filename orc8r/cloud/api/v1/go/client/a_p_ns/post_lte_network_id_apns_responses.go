// Code generated by go-swagger; DO NOT EDIT.

package a_p_ns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PostLTENetworkIDAPNSReader is a Reader for the PostLTENetworkIDAPNS structure.
type PostLTENetworkIDAPNSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLTENetworkIDAPNSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLTENetworkIDAPNSCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostLTENetworkIDAPNSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostLTENetworkIDAPNSCreated creates a PostLTENetworkIDAPNSCreated with default headers values
func NewPostLTENetworkIDAPNSCreated() *PostLTENetworkIDAPNSCreated {
	return &PostLTENetworkIDAPNSCreated{}
}

/*PostLTENetworkIDAPNSCreated handles this case with default header values.

Success
*/
type PostLTENetworkIDAPNSCreated struct {
}

func (o *PostLTENetworkIDAPNSCreated) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/apns][%d] postLteNetworkIdApnsCreated ", 201)
}

func (o *PostLTENetworkIDAPNSCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLTENetworkIDAPNSDefault creates a PostLTENetworkIDAPNSDefault with default headers values
func NewPostLTENetworkIDAPNSDefault(code int) *PostLTENetworkIDAPNSDefault {
	return &PostLTENetworkIDAPNSDefault{
		_statusCode: code,
	}
}

/*PostLTENetworkIDAPNSDefault handles this case with default header values.

Unexpected Error
*/
type PostLTENetworkIDAPNSDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post LTE network ID APNS default response
func (o *PostLTENetworkIDAPNSDefault) Code() int {
	return o._statusCode
}

func (o *PostLTENetworkIDAPNSDefault) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/apns][%d] PostLTENetworkIDAPNS default  %+v", o._statusCode, o.Payload)
}

func (o *PostLTENetworkIDAPNSDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostLTENetworkIDAPNSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
