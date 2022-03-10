// Code generated by go-swagger; DO NOT EDIT.

package enode_bs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutLTENetworkIDEnodebsENODEBSerialReader is a Reader for the PutLTENetworkIDEnodebsENODEBSerial structure.
type PutLTENetworkIDEnodebsENODEBSerialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDEnodebsENODEBSerialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDEnodebsENODEBSerialNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDEnodebsENODEBSerialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDEnodebsENODEBSerialNoContent creates a PutLTENetworkIDEnodebsENODEBSerialNoContent with default headers values
func NewPutLTENetworkIDEnodebsENODEBSerialNoContent() *PutLTENetworkIDEnodebsENODEBSerialNoContent {
	return &PutLTENetworkIDEnodebsENODEBSerialNoContent{}
}

/*PutLTENetworkIDEnodebsENODEBSerialNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDEnodebsENODEBSerialNoContent struct {
}

func (o *PutLTENetworkIDEnodebsENODEBSerialNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/enodebs/{enodeb_serial}][%d] putLteNetworkIdEnodebsEnodebSerialNoContent ", 204)
}

func (o *PutLTENetworkIDEnodebsENODEBSerialNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDEnodebsENODEBSerialDefault creates a PutLTENetworkIDEnodebsENODEBSerialDefault with default headers values
func NewPutLTENetworkIDEnodebsENODEBSerialDefault(code int) *PutLTENetworkIDEnodebsENODEBSerialDefault {
	return &PutLTENetworkIDEnodebsENODEBSerialDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDEnodebsENODEBSerialDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDEnodebsENODEBSerialDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID enodebs ENODEB serial default response
func (o *PutLTENetworkIDEnodebsENODEBSerialDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDEnodebsENODEBSerialDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/enodebs/{enodeb_serial}][%d] PutLTENetworkIDEnodebsENODEBSerial default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDEnodebsENODEBSerialDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDEnodebsENODEBSerialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
