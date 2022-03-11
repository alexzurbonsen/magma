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

// GetNetworksNetworkIDPrometheusSeriesReader is a Reader for the GetNetworksNetworkIDPrometheusSeries structure.
type GetNetworksNetworkIDPrometheusSeriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDPrometheusSeriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDPrometheusSeriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDPrometheusSeriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDPrometheusSeriesOK creates a GetNetworksNetworkIDPrometheusSeriesOK with default headers values
func NewGetNetworksNetworkIDPrometheusSeriesOK() *GetNetworksNetworkIDPrometheusSeriesOK {
	return &GetNetworksNetworkIDPrometheusSeriesOK{}
}

/* GetNetworksNetworkIDPrometheusSeriesOK describes a response with status code 200, with default header values.

List of metric names
*/
type GetNetworksNetworkIDPrometheusSeriesOK struct {
	Payload []models.PrometheusLabelset
}

func (o *GetNetworksNetworkIDPrometheusSeriesOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/prometheus/series][%d] getNetworksNetworkIdPrometheusSeriesOK  %+v", 200, o.Payload)
}
func (o *GetNetworksNetworkIDPrometheusSeriesOK) GetPayload() []models.PrometheusLabelset {
	return o.Payload
}

func (o *GetNetworksNetworkIDPrometheusSeriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDPrometheusSeriesDefault creates a GetNetworksNetworkIDPrometheusSeriesDefault with default headers values
func NewGetNetworksNetworkIDPrometheusSeriesDefault(code int) *GetNetworksNetworkIDPrometheusSeriesDefault {
	return &GetNetworksNetworkIDPrometheusSeriesDefault{
		_statusCode: code,
	}
}

/* GetNetworksNetworkIDPrometheusSeriesDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDPrometheusSeriesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID prometheus series default response
func (o *GetNetworksNetworkIDPrometheusSeriesDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDPrometheusSeriesDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/prometheus/series][%d] GetNetworksNetworkIDPrometheusSeries default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworksNetworkIDPrometheusSeriesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDPrometheusSeriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
