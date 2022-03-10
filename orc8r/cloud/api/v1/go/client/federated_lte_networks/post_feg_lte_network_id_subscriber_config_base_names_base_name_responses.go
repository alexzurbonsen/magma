// Code generated by go-swagger; DO NOT EDIT.

package federated_lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameReader is a Reader for the PostFegLTENetworkIDSubscriberConfigBaseNamesBaseName structure.
type PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated creates a PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated with default headers values
func NewPostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated() *PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated {
	return &PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated{}
}

/*PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated handles this case with default header values.

Success
*/
type PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated struct {
}

func (o *PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated) Error() string {
	return fmt.Sprintf("[POST /feg_lte/{network_id}/subscriber_config/base_names/{base_name}][%d] postFegLteNetworkIdSubscriberConfigBaseNamesBaseNameCreated ", 201)
}

func (o *PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault creates a PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault with default headers values
func NewPostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault(code int) *PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault {
	return &PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault{
		_statusCode: code,
	}
}

/*PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault handles this case with default header values.

Unexpected Error
*/
type PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post feg LTE network ID subscriber config base names base name default response
func (o *PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault) Code() int {
	return o._statusCode
}

func (o *PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault) Error() string {
	return fmt.Sprintf("[POST /feg_lte/{network_id}/subscriber_config/base_names/{base_name}][%d] PostFegLTENetworkIDSubscriberConfigBaseNamesBaseName default  %+v", o._statusCode, o.Payload)
}

func (o *PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFegLTENetworkIDSubscriberConfigBaseNamesBaseNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
