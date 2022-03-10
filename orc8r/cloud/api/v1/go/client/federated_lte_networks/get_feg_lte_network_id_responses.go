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

// GetFegLTENetworkIDReader is a Reader for the GetFegLTENetworkID structure.
type GetFegLTENetworkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFegLTENetworkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFegLTENetworkIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFegLTENetworkIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFegLTENetworkIDOK creates a GetFegLTENetworkIDOK with default headers values
func NewGetFegLTENetworkIDOK() *GetFegLTENetworkIDOK {
	return &GetFegLTENetworkIDOK{}
}

/*GetFegLTENetworkIDOK handles this case with default header values.

Full description of a federated LTE network
*/
type GetFegLTENetworkIDOK struct {
	Payload *models.FegLTENetwork
}

func (o *GetFegLTENetworkIDOK) Error() string {
	return fmt.Sprintf("[GET /feg_lte/{network_id}][%d] getFegLteNetworkIdOK  %+v", 200, o.Payload)
}

func (o *GetFegLTENetworkIDOK) GetPayload() *models.FegLTENetwork {
	return o.Payload
}

func (o *GetFegLTENetworkIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FegLTENetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFegLTENetworkIDDefault creates a GetFegLTENetworkIDDefault with default headers values
func NewGetFegLTENetworkIDDefault(code int) *GetFegLTENetworkIDDefault {
	return &GetFegLTENetworkIDDefault{
		_statusCode: code,
	}
}

/*GetFegLTENetworkIDDefault handles this case with default header values.

Unexpected Error
*/
type GetFegLTENetworkIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get feg LTE network ID default response
func (o *GetFegLTENetworkIDDefault) Code() int {
	return o._statusCode
}

func (o *GetFegLTENetworkIDDefault) Error() string {
	return fmt.Sprintf("[GET /feg_lte/{network_id}][%d] GetFegLTENetworkID default  %+v", o._statusCode, o.Payload)
}

func (o *GetFegLTENetworkIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFegLTENetworkIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
