// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetEventsNetworkIDAboutCountReader is a Reader for the GetEventsNetworkIDAboutCount structure.
type GetEventsNetworkIDAboutCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventsNetworkIDAboutCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventsNetworkIDAboutCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEventsNetworkIDAboutCountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEventsNetworkIDAboutCountOK creates a GetEventsNetworkIDAboutCountOK with default headers values
func NewGetEventsNetworkIDAboutCountOK() *GetEventsNetworkIDAboutCountOK {
	return &GetEventsNetworkIDAboutCountOK{}
}

/*GetEventsNetworkIDAboutCountOK handles this case with default header values.

Success
*/
type GetEventsNetworkIDAboutCountOK struct {
	Payload int64
}

func (o *GetEventsNetworkIDAboutCountOK) Error() string {
	return fmt.Sprintf("[GET /events/{network_id}/about/count][%d] getEventsNetworkIdAboutCountOK  %+v", 200, o.Payload)
}

func (o *GetEventsNetworkIDAboutCountOK) GetPayload() int64 {
	return o.Payload
}

func (o *GetEventsNetworkIDAboutCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsNetworkIDAboutCountDefault creates a GetEventsNetworkIDAboutCountDefault with default headers values
func NewGetEventsNetworkIDAboutCountDefault(code int) *GetEventsNetworkIDAboutCountDefault {
	return &GetEventsNetworkIDAboutCountDefault{
		_statusCode: code,
	}
}

/*GetEventsNetworkIDAboutCountDefault handles this case with default header values.

Unexpected Error
*/
type GetEventsNetworkIDAboutCountDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get events network ID about count default response
func (o *GetEventsNetworkIDAboutCountDefault) Code() int {
	return o._statusCode
}

func (o *GetEventsNetworkIDAboutCountDefault) Error() string {
	return fmt.Sprintf("[GET /events/{network_id}/about/count][%d] GetEventsNetworkIDAboutCount default  %+v", o._statusCode, o.Payload)
}

func (o *GetEventsNetworkIDAboutCountDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEventsNetworkIDAboutCountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
