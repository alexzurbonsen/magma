// Code generated by go-swagger; DO NOT EDIT.

package upgrades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutChannelsChannelIDReader is a Reader for the PutChannelsChannelID structure.
type PutChannelsChannelIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutChannelsChannelIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutChannelsChannelIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutChannelsChannelIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutChannelsChannelIDNoContent creates a PutChannelsChannelIDNoContent with default headers values
func NewPutChannelsChannelIDNoContent() *PutChannelsChannelIDNoContent {
	return &PutChannelsChannelIDNoContent{}
}

/*PutChannelsChannelIDNoContent handles this case with default header values.

Success
*/
type PutChannelsChannelIDNoContent struct {
}

func (o *PutChannelsChannelIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /channels/{channel_id}][%d] putChannelsChannelIdNoContent ", 204)
}

func (o *PutChannelsChannelIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutChannelsChannelIDDefault creates a PutChannelsChannelIDDefault with default headers values
func NewPutChannelsChannelIDDefault(code int) *PutChannelsChannelIDDefault {
	return &PutChannelsChannelIDDefault{
		_statusCode: code,
	}
}

/*PutChannelsChannelIDDefault handles this case with default header values.

Unexpected Error
*/
type PutChannelsChannelIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put channels channel ID default response
func (o *PutChannelsChannelIDDefault) Code() int {
	return o._statusCode
}

func (o *PutChannelsChannelIDDefault) Error() string {
	return fmt.Sprintf("[PUT /channels/{channel_id}][%d] PutChannelsChannelID default  %+v", o._statusCode, o.Payload)
}

func (o *PutChannelsChannelIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutChannelsChannelIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
