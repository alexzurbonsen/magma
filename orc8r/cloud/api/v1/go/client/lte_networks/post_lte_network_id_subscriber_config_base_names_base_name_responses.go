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

// PostLTENetworkIDSubscriberConfigBaseNamesBaseNameReader is a Reader for the PostLTENetworkIDSubscriberConfigBaseNamesBaseName structure.
type PostLTENetworkIDSubscriberConfigBaseNamesBaseNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLTENetworkIDSubscriberConfigBaseNamesBaseNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated creates a PostLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated with default headers values
func NewPostLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated() *PostLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated {
	return &PostLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated{}
}

/* PostLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated describes a response with status code 201, with default header values.

Success
*/
type PostLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated struct {
}

func (o *PostLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/subscriber_config/base_names/{base_name}][%d] postLteNetworkIdSubscriberConfigBaseNamesBaseNameCreated ", 201)
}

func (o *PostLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault creates a PostLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault with default headers values
func NewPostLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault(code int) *PostLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault {
	return &PostLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault{
		_statusCode: code,
	}
}

/* PostLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PostLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post LTE network ID subscriber config base names base name default response
func (o *PostLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault) Code() int {
	return o._statusCode
}

func (o *PostLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/subscriber_config/base_names/{base_name}][%d] PostLTENetworkIDSubscriberConfigBaseNamesBaseName default  %+v", o._statusCode, o.Payload)
}
func (o *PostLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
