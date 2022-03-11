// Code generated by go-swagger; DO NOT EDIT.

package network_probes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDReader is a Reader for the DeleteLTENetworkIDNetworkProbeDestinationsDestinationID structure.
type DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteLTENetworkIDNetworkProbeDestinationsDestinationIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent creates a DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent with default headers values
func NewDeleteLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent() *DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent {
	return &DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent{}
}

/* DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent struct {
}

func (o *DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/network_probe/destinations/{destination_id}][%d] deleteLteNetworkIdNetworkProbeDestinationsDestinationIdNoContent ", 204)
}

func (o *DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLTENetworkIDNetworkProbeDestinationsDestinationIDDefault creates a DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDDefault with default headers values
func NewDeleteLTENetworkIDNetworkProbeDestinationsDestinationIDDefault(code int) *DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDDefault {
	return &DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDDefault{
		_statusCode: code,
	}
}

/* DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete LTE network ID network probe destinations destination ID default response
func (o *DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/network_probe/destinations/{destination_id}][%d] DeleteLTENetworkIDNetworkProbeDestinationsDestinationID default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
