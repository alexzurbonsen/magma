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

// DeleteNetworksNetworkIDPrometheusAlertReceiverReader is a Reader for the DeleteNetworksNetworkIDPrometheusAlertReceiver structure.
type DeleteNetworksNetworkIDPrometheusAlertReceiverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworksNetworkIDPrometheusAlertReceiverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNetworksNetworkIDPrometheusAlertReceiverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteNetworksNetworkIDPrometheusAlertReceiverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworksNetworkIDPrometheusAlertReceiverOK creates a DeleteNetworksNetworkIDPrometheusAlertReceiverOK with default headers values
func NewDeleteNetworksNetworkIDPrometheusAlertReceiverOK() *DeleteNetworksNetworkIDPrometheusAlertReceiverOK {
	return &DeleteNetworksNetworkIDPrometheusAlertReceiverOK{}
}

/* DeleteNetworksNetworkIDPrometheusAlertReceiverOK describes a response with status code 200, with default header values.

Deleted
*/
type DeleteNetworksNetworkIDPrometheusAlertReceiverOK struct {
}

func (o *DeleteNetworksNetworkIDPrometheusAlertReceiverOK) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/prometheus/alert_receiver][%d] deleteNetworksNetworkIdPrometheusAlertReceiverOK ", 200)
}

func (o *DeleteNetworksNetworkIDPrometheusAlertReceiverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworksNetworkIDPrometheusAlertReceiverDefault creates a DeleteNetworksNetworkIDPrometheusAlertReceiverDefault with default headers values
func NewDeleteNetworksNetworkIDPrometheusAlertReceiverDefault(code int) *DeleteNetworksNetworkIDPrometheusAlertReceiverDefault {
	return &DeleteNetworksNetworkIDPrometheusAlertReceiverDefault{
		_statusCode: code,
	}
}

/* DeleteNetworksNetworkIDPrometheusAlertReceiverDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteNetworksNetworkIDPrometheusAlertReceiverDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete networks network ID prometheus alert receiver default response
func (o *DeleteNetworksNetworkIDPrometheusAlertReceiverDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworksNetworkIDPrometheusAlertReceiverDefault) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/prometheus/alert_receiver][%d] DeleteNetworksNetworkIDPrometheusAlertReceiver default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNetworksNetworkIDPrometheusAlertReceiverDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNetworksNetworkIDPrometheusAlertReceiverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
