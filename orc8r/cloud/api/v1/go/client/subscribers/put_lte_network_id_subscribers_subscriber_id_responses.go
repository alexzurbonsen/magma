// Code generated by go-swagger; DO NOT EDIT.

package subscribers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutLTENetworkIDSubscribersSubscriberIDReader is a Reader for the PutLTENetworkIDSubscribersSubscriberID structure.
type PutLTENetworkIDSubscribersSubscriberIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDSubscribersSubscriberIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDSubscribersSubscriberIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDSubscribersSubscriberIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDSubscribersSubscriberIDNoContent creates a PutLTENetworkIDSubscribersSubscriberIDNoContent with default headers values
func NewPutLTENetworkIDSubscribersSubscriberIDNoContent() *PutLTENetworkIDSubscribersSubscriberIDNoContent {
	return &PutLTENetworkIDSubscribersSubscriberIDNoContent{}
}

/* PutLTENetworkIDSubscribersSubscriberIDNoContent describes a response with status code 204, with default header values.

Success
*/
type PutLTENetworkIDSubscribersSubscriberIDNoContent struct {
}

func (o *PutLTENetworkIDSubscribersSubscriberIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/subscribers/{subscriber_id}][%d] putLteNetworkIdSubscribersSubscriberIdNoContent ", 204)
}

func (o *PutLTENetworkIDSubscribersSubscriberIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDSubscribersSubscriberIDDefault creates a PutLTENetworkIDSubscribersSubscriberIDDefault with default headers values
func NewPutLTENetworkIDSubscribersSubscriberIDDefault(code int) *PutLTENetworkIDSubscribersSubscriberIDDefault {
	return &PutLTENetworkIDSubscribersSubscriberIDDefault{
		_statusCode: code,
	}
}

/* PutLTENetworkIDSubscribersSubscriberIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutLTENetworkIDSubscribersSubscriberIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID subscribers subscriber ID default response
func (o *PutLTENetworkIDSubscribersSubscriberIDDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDSubscribersSubscriberIDDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/subscribers/{subscriber_id}][%d] PutLTENetworkIDSubscribersSubscriberID default  %+v", o._statusCode, o.Payload)
}
func (o *PutLTENetworkIDSubscribersSubscriberIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDSubscribersSubscriberIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
