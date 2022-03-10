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

// PostLTENetworkIDMsisdnsReader is a Reader for the PostLTENetworkIDMsisdns structure.
type PostLTENetworkIDMsisdnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLTENetworkIDMsisdnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLTENetworkIDMsisdnsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostLTENetworkIDMsisdnsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostLTENetworkIDMsisdnsCreated creates a PostLTENetworkIDMsisdnsCreated with default headers values
func NewPostLTENetworkIDMsisdnsCreated() *PostLTENetworkIDMsisdnsCreated {
	return &PostLTENetworkIDMsisdnsCreated{}
}

/*PostLTENetworkIDMsisdnsCreated handles this case with default header values.

Success
*/
type PostLTENetworkIDMsisdnsCreated struct {
}

func (o *PostLTENetworkIDMsisdnsCreated) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/msisdns][%d] postLteNetworkIdMsisdnsCreated ", 201)
}

func (o *PostLTENetworkIDMsisdnsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLTENetworkIDMsisdnsDefault creates a PostLTENetworkIDMsisdnsDefault with default headers values
func NewPostLTENetworkIDMsisdnsDefault(code int) *PostLTENetworkIDMsisdnsDefault {
	return &PostLTENetworkIDMsisdnsDefault{
		_statusCode: code,
	}
}

/*PostLTENetworkIDMsisdnsDefault handles this case with default header values.

Unexpected Error
*/
type PostLTENetworkIDMsisdnsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post LTE network ID msisdns default response
func (o *PostLTENetworkIDMsisdnsDefault) Code() int {
	return o._statusCode
}

func (o *PostLTENetworkIDMsisdnsDefault) Error() string {
	return fmt.Sprintf("[POST /lte/{network_id}/msisdns][%d] PostLTENetworkIDMsisdns default  %+v", o._statusCode, o.Payload)
}

func (o *PostLTENetworkIDMsisdnsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostLTENetworkIDMsisdnsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
