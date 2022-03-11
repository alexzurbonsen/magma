// Code generated by go-swagger; DO NOT EDIT.

package call_tracing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDTracingTraceIDReader is a Reader for the GetNetworksNetworkIDTracingTraceID structure.
type GetNetworksNetworkIDTracingTraceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDTracingTraceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDTracingTraceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDTracingTraceIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDTracingTraceIDOK creates a GetNetworksNetworkIDTracingTraceIDOK with default headers values
func NewGetNetworksNetworkIDTracingTraceIDOK() *GetNetworksNetworkIDTracingTraceIDOK {
	return &GetNetworksNetworkIDTracingTraceIDOK{}
}

/* GetNetworksNetworkIDTracingTraceIDOK describes a response with status code 200, with default header values.

Show tracing status
*/
type GetNetworksNetworkIDTracingTraceIDOK struct {
	Payload *models.CallTrace
}

func (o *GetNetworksNetworkIDTracingTraceIDOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/tracing/{trace_id}][%d] getNetworksNetworkIdTracingTraceIdOK  %+v", 200, o.Payload)
}
func (o *GetNetworksNetworkIDTracingTraceIDOK) GetPayload() *models.CallTrace {
	return o.Payload
}

func (o *GetNetworksNetworkIDTracingTraceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallTrace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDTracingTraceIDDefault creates a GetNetworksNetworkIDTracingTraceIDDefault with default headers values
func NewGetNetworksNetworkIDTracingTraceIDDefault(code int) *GetNetworksNetworkIDTracingTraceIDDefault {
	return &GetNetworksNetworkIDTracingTraceIDDefault{
		_statusCode: code,
	}
}

/* GetNetworksNetworkIDTracingTraceIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDTracingTraceIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID tracing trace ID default response
func (o *GetNetworksNetworkIDTracingTraceIDDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDTracingTraceIDDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/tracing/{trace_id}][%d] GetNetworksNetworkIDTracingTraceID default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworksNetworkIDTracingTraceIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDTracingTraceIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
