// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PostNetworksNetworkIDPrometheusAlertConfigReader is a Reader for the PostNetworksNetworkIDPrometheusAlertConfig structure.
type PostNetworksNetworkIDPrometheusAlertConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworksNetworkIDPrometheusAlertConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostNetworksNetworkIDPrometheusAlertConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostNetworksNetworkIDPrometheusAlertConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostNetworksNetworkIDPrometheusAlertConfigCreated creates a PostNetworksNetworkIDPrometheusAlertConfigCreated with default headers values
func NewPostNetworksNetworkIDPrometheusAlertConfigCreated() *PostNetworksNetworkIDPrometheusAlertConfigCreated {
	return &PostNetworksNetworkIDPrometheusAlertConfigCreated{}
}

/* PostNetworksNetworkIDPrometheusAlertConfigCreated describes a response with status code 201, with default header values.

Created
*/
type PostNetworksNetworkIDPrometheusAlertConfigCreated struct {
}

func (o *PostNetworksNetworkIDPrometheusAlertConfigCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/prometheus/alert_config][%d] postNetworksNetworkIdPrometheusAlertConfigCreated ", 201)
}

func (o *PostNetworksNetworkIDPrometheusAlertConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostNetworksNetworkIDPrometheusAlertConfigDefault creates a PostNetworksNetworkIDPrometheusAlertConfigDefault with default headers values
func NewPostNetworksNetworkIDPrometheusAlertConfigDefault(code int) *PostNetworksNetworkIDPrometheusAlertConfigDefault {
	return &PostNetworksNetworkIDPrometheusAlertConfigDefault{
		_statusCode: code,
	}
}

/* PostNetworksNetworkIDPrometheusAlertConfigDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PostNetworksNetworkIDPrometheusAlertConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post networks network ID prometheus alert config default response
func (o *PostNetworksNetworkIDPrometheusAlertConfigDefault) Code() int {
	return o._statusCode
}

func (o *PostNetworksNetworkIDPrometheusAlertConfigDefault) Error() string {
	return fmt.Sprintf("[POST /networks/{network_id}/prometheus/alert_config][%d] PostNetworksNetworkIDPrometheusAlertConfig default  %+v", o._statusCode, o.Payload)
}
func (o *PostNetworksNetworkIDPrometheusAlertConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNetworksNetworkIDPrometheusAlertConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
