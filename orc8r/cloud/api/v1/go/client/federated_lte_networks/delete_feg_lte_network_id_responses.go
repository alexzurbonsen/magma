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

// DeleteFegLTENetworkIDReader is a Reader for the DeleteFegLTENetworkID structure.
type DeleteFegLTENetworkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFegLTENetworkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFegLTENetworkIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteFegLTENetworkIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFegLTENetworkIDNoContent creates a DeleteFegLTENetworkIDNoContent with default headers values
func NewDeleteFegLTENetworkIDNoContent() *DeleteFegLTENetworkIDNoContent {
	return &DeleteFegLTENetworkIDNoContent{}
}

/*DeleteFegLTENetworkIDNoContent handles this case with default header values.

Success
*/
type DeleteFegLTENetworkIDNoContent struct {
}

func (o *DeleteFegLTENetworkIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /feg_lte/{network_id}][%d] deleteFegLteNetworkIdNoContent ", 204)
}

func (o *DeleteFegLTENetworkIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFegLTENetworkIDDefault creates a DeleteFegLTENetworkIDDefault with default headers values
func NewDeleteFegLTENetworkIDDefault(code int) *DeleteFegLTENetworkIDDefault {
	return &DeleteFegLTENetworkIDDefault{
		_statusCode: code,
	}
}

/*DeleteFegLTENetworkIDDefault handles this case with default header values.

Unexpected Error
*/
type DeleteFegLTENetworkIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete feg LTE network ID default response
func (o *DeleteFegLTENetworkIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFegLTENetworkIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /feg_lte/{network_id}][%d] DeleteFegLTENetworkID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteFegLTENetworkIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFegLTENetworkIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
