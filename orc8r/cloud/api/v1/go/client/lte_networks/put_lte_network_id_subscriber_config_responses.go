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

// PutLTENetworkIDSubscriberConfigReader is a Reader for the PutLTENetworkIDSubscriberConfig structure.
type PutLTENetworkIDSubscriberConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDSubscriberConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDSubscriberConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDSubscriberConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDSubscriberConfigNoContent creates a PutLTENetworkIDSubscriberConfigNoContent with default headers values
func NewPutLTENetworkIDSubscriberConfigNoContent() *PutLTENetworkIDSubscriberConfigNoContent {
	return &PutLTENetworkIDSubscriberConfigNoContent{}
}

/* PutLTENetworkIDSubscriberConfigNoContent describes a response with status code 204, with default header values.

Success
*/
type PutLTENetworkIDSubscriberConfigNoContent struct {
}

func (o *PutLTENetworkIDSubscriberConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/subscriber_config][%d] putLteNetworkIdSubscriberConfigNoContent ", 204)
}

func (o *PutLTENetworkIDSubscriberConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDSubscriberConfigDefault creates a PutLTENetworkIDSubscriberConfigDefault with default headers values
func NewPutLTENetworkIDSubscriberConfigDefault(code int) *PutLTENetworkIDSubscriberConfigDefault {
	return &PutLTENetworkIDSubscriberConfigDefault{
		_statusCode: code,
	}
}

/* PutLTENetworkIDSubscriberConfigDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutLTENetworkIDSubscriberConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID subscriber config default response
func (o *PutLTENetworkIDSubscriberConfigDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDSubscriberConfigDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/subscriber_config][%d] PutLTENetworkIDSubscriberConfig default  %+v", o._statusCode, o.Payload)
}
func (o *PutLTENetworkIDSubscriberConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDSubscriberConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
