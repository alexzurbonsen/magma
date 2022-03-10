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

// GetNetworksNetworkIDPrometheusQueryRangeReader is a Reader for the GetNetworksNetworkIDPrometheusQueryRange structure.
type GetNetworksNetworkIDPrometheusQueryRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDPrometheusQueryRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDPrometheusQueryRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDPrometheusQueryRangeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDPrometheusQueryRangeOK creates a GetNetworksNetworkIDPrometheusQueryRangeOK with default headers values
func NewGetNetworksNetworkIDPrometheusQueryRangeOK() *GetNetworksNetworkIDPrometheusQueryRangeOK {
	return &GetNetworksNetworkIDPrometheusQueryRangeOK{}
}

/*GetNetworksNetworkIDPrometheusQueryRangeOK handles this case with default header values.

List of PromQL metrics results
*/
type GetNetworksNetworkIDPrometheusQueryRangeOK struct {
	Payload *models.PromqlReturnObject
}

func (o *GetNetworksNetworkIDPrometheusQueryRangeOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/prometheus/query_range][%d] getNetworksNetworkIdPrometheusQueryRangeOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDPrometheusQueryRangeOK) GetPayload() *models.PromqlReturnObject {
	return o.Payload
}

func (o *GetNetworksNetworkIDPrometheusQueryRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PromqlReturnObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDPrometheusQueryRangeDefault creates a GetNetworksNetworkIDPrometheusQueryRangeDefault with default headers values
func NewGetNetworksNetworkIDPrometheusQueryRangeDefault(code int) *GetNetworksNetworkIDPrometheusQueryRangeDefault {
	return &GetNetworksNetworkIDPrometheusQueryRangeDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDPrometheusQueryRangeDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDPrometheusQueryRangeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID prometheus query range default response
func (o *GetNetworksNetworkIDPrometheusQueryRangeDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDPrometheusQueryRangeDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/prometheus/query_range][%d] GetNetworksNetworkIDPrometheusQueryRange default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDPrometheusQueryRangeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDPrometheusQueryRangeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
