// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetFegNetworkIDFederationReader is a Reader for the GetFegNetworkIDFederation structure.
type GetFegNetworkIDFederationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFegNetworkIDFederationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFegNetworkIDFederationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFegNetworkIDFederationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFegNetworkIDFederationOK creates a GetFegNetworkIDFederationOK with default headers values
func NewGetFegNetworkIDFederationOK() *GetFegNetworkIDFederationOK {
	return &GetFegNetworkIDFederationOK{}
}

/*GetFegNetworkIDFederationOK handles this case with default header values.

Retrieved Network Federation Configs
*/
type GetFegNetworkIDFederationOK struct {
	Payload *models.NetworkFederationConfigs
}

func (o *GetFegNetworkIDFederationOK) Error() string {
	return fmt.Sprintf("[GET /feg/{network_id}/federation][%d] getFegNetworkIdFederationOK  %+v", 200, o.Payload)
}

func (o *GetFegNetworkIDFederationOK) GetPayload() *models.NetworkFederationConfigs {
	return o.Payload
}

func (o *GetFegNetworkIDFederationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkFederationConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFegNetworkIDFederationDefault creates a GetFegNetworkIDFederationDefault with default headers values
func NewGetFegNetworkIDFederationDefault(code int) *GetFegNetworkIDFederationDefault {
	return &GetFegNetworkIDFederationDefault{
		_statusCode: code,
	}
}

/*GetFegNetworkIDFederationDefault handles this case with default header values.

Unexpected Error
*/
type GetFegNetworkIDFederationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get feg network ID federation default response
func (o *GetFegNetworkIDFederationDefault) Code() int {
	return o._statusCode
}

func (o *GetFegNetworkIDFederationDefault) Error() string {
	return fmt.Sprintf("[GET /feg/{network_id}/federation][%d] GetFegNetworkIDFederation default  %+v", o._statusCode, o.Payload)
}

func (o *GetFegNetworkIDFederationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFegNetworkIDFederationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
