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

// DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameReader is a Reader for the DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseName structure.
type DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameNoContent creates a DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameNoContent with default headers values
func NewDeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameNoContent() *DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameNoContent {
	return &DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameNoContent{}
}

/*DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameNoContent handles this case with default header values.

Success
*/
type DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameNoContent struct {
}

func (o *DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /cwf/{network_id}/subscriber_config/base_names/{base_name}][%d] deleteCwfNetworkIdSubscriberConfigBaseNamesBaseNameNoContent ", 204)
}

func (o *DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault creates a DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault with default headers values
func NewDeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault(code int) *DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault {
	return &DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault{
		_statusCode: code,
	}
}

/*DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault handles this case with default header values.

Unexpected Error
*/
type DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete cwf network ID subscriber config base names base name default response
func (o *DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /cwf/{network_id}/subscriber_config/base_names/{base_name}][%d] DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCwfNetworkIDSubscriberConfigBaseNamesBaseNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
