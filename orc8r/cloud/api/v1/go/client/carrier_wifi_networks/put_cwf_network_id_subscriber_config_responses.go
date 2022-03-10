// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutCwfNetworkIDSubscriberConfigReader is a Reader for the PutCwfNetworkIDSubscriberConfig structure.
type PutCwfNetworkIDSubscriberConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCwfNetworkIDSubscriberConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutCwfNetworkIDSubscriberConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCwfNetworkIDSubscriberConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCwfNetworkIDSubscriberConfigNoContent creates a PutCwfNetworkIDSubscriberConfigNoContent with default headers values
func NewPutCwfNetworkIDSubscriberConfigNoContent() *PutCwfNetworkIDSubscriberConfigNoContent {
	return &PutCwfNetworkIDSubscriberConfigNoContent{}
}

/*PutCwfNetworkIDSubscriberConfigNoContent handles this case with default header values.

Success
*/
type PutCwfNetworkIDSubscriberConfigNoContent struct {
}

func (o *PutCwfNetworkIDSubscriberConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/subscriber_config][%d] putCwfNetworkIdSubscriberConfigNoContent ", 204)
}

func (o *PutCwfNetworkIDSubscriberConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCwfNetworkIDSubscriberConfigDefault creates a PutCwfNetworkIDSubscriberConfigDefault with default headers values
func NewPutCwfNetworkIDSubscriberConfigDefault(code int) *PutCwfNetworkIDSubscriberConfigDefault {
	return &PutCwfNetworkIDSubscriberConfigDefault{
		_statusCode: code,
	}
}

/*PutCwfNetworkIDSubscriberConfigDefault handles this case with default header values.

Unexpected Error
*/
type PutCwfNetworkIDSubscriberConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put cwf network ID subscriber config default response
func (o *PutCwfNetworkIDSubscriberConfigDefault) Code() int {
	return o._statusCode
}

func (o *PutCwfNetworkIDSubscriberConfigDefault) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/subscriber_config][%d] PutCwfNetworkIDSubscriberConfig default  %+v", o._statusCode, o.Payload)
}

func (o *PutCwfNetworkIDSubscriberConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCwfNetworkIDSubscriberConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
