// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetTenantsTargetsMetadataReader is a Reader for the GetTenantsTargetsMetadata structure.
type GetTenantsTargetsMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantsTargetsMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantsTargetsMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTenantsTargetsMetadataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTenantsTargetsMetadataOK creates a GetTenantsTargetsMetadataOK with default headers values
func NewGetTenantsTargetsMetadataOK() *GetTenantsTargetsMetadataOK {
	return &GetTenantsTargetsMetadataOK{}
}

/* GetTenantsTargetsMetadataOK describes a response with status code 200, with default header values.

Info of metrics
*/
type GetTenantsTargetsMetadataOK struct {
	Payload []*models.PrometheusTargetsMetadata
}

func (o *GetTenantsTargetsMetadataOK) Error() string {
	return fmt.Sprintf("[GET /tenants/targets_metadata][%d] getTenantsTargetsMetadataOK  %+v", 200, o.Payload)
}
func (o *GetTenantsTargetsMetadataOK) GetPayload() []*models.PrometheusTargetsMetadata {
	return o.Payload
}

func (o *GetTenantsTargetsMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantsTargetsMetadataDefault creates a GetTenantsTargetsMetadataDefault with default headers values
func NewGetTenantsTargetsMetadataDefault(code int) *GetTenantsTargetsMetadataDefault {
	return &GetTenantsTargetsMetadataDefault{
		_statusCode: code,
	}
}

/* GetTenantsTargetsMetadataDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetTenantsTargetsMetadataDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get tenants targets metadata default response
func (o *GetTenantsTargetsMetadataDefault) Code() int {
	return o._statusCode
}

func (o *GetTenantsTargetsMetadataDefault) Error() string {
	return fmt.Sprintf("[GET /tenants/targets_metadata][%d] GetTenantsTargetsMetadata default  %+v", o._statusCode, o.Payload)
}
func (o *GetTenantsTargetsMetadataDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTenantsTargetsMetadataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
