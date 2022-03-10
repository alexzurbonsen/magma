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

// GetNetworksNetworkIDPrometheusAlertReceiverRouteReader is a Reader for the GetNetworksNetworkIDPrometheusAlertReceiverRoute structure.
type GetNetworksNetworkIDPrometheusAlertReceiverRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDPrometheusAlertReceiverRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDPrometheusAlertReceiverRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworksNetworkIDPrometheusAlertReceiverRouteOK creates a GetNetworksNetworkIDPrometheusAlertReceiverRouteOK with default headers values
func NewGetNetworksNetworkIDPrometheusAlertReceiverRouteOK() *GetNetworksNetworkIDPrometheusAlertReceiverRouteOK {
	return &GetNetworksNetworkIDPrometheusAlertReceiverRouteOK{}
}

/*GetNetworksNetworkIDPrometheusAlertReceiverRouteOK handles this case with default header values.

Alerting tree
*/
type GetNetworksNetworkIDPrometheusAlertReceiverRouteOK struct {
	Payload *models.AlertRoutingTree
}

func (o *GetNetworksNetworkIDPrometheusAlertReceiverRouteOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/prometheus/alert_receiver/route][%d] getNetworksNetworkIdPrometheusAlertReceiverRouteOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDPrometheusAlertReceiverRouteOK) GetPayload() *models.AlertRoutingTree {
	return o.Payload
}

func (o *GetNetworksNetworkIDPrometheusAlertReceiverRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertRoutingTree)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
