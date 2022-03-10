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

// PutNetworksNetworkIDPrometheusAlertReceiverReceiverReader is a Reader for the PutNetworksNetworkIDPrometheusAlertReceiverReceiver structure.
type PutNetworksNetworkIDPrometheusAlertReceiverReceiverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverOK creates a PutNetworksNetworkIDPrometheusAlertReceiverReceiverOK with default headers values
func NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverOK() *PutNetworksNetworkIDPrometheusAlertReceiverReceiverOK {
	return &PutNetworksNetworkIDPrometheusAlertReceiverReceiverOK{}
}

/*PutNetworksNetworkIDPrometheusAlertReceiverReceiverOK handles this case with default header values.

Updated
*/
type PutNetworksNetworkIDPrometheusAlertReceiverReceiverOK struct {
}

func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/prometheus/alert_receiver/{receiver}][%d] putNetworksNetworkIdPrometheusAlertReceiverReceiverOK ", 200)
}

func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverDefault creates a PutNetworksNetworkIDPrometheusAlertReceiverReceiverDefault with default headers values
func NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverDefault(code int) *PutNetworksNetworkIDPrometheusAlertReceiverReceiverDefault {
	return &PutNetworksNetworkIDPrometheusAlertReceiverReceiverDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDPrometheusAlertReceiverReceiverDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDPrometheusAlertReceiverReceiverDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID prometheus alert receiver receiver default response
func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/prometheus/alert_receiver/{receiver}][%d] PutNetworksNetworkIDPrometheusAlertReceiverReceiver default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDPrometheusAlertReceiverReceiverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
