// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutCwfNetworkIDCarrierWifiReader is a Reader for the PutCwfNetworkIDCarrierWifi structure.
type PutCwfNetworkIDCarrierWifiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCwfNetworkIDCarrierWifiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCwfNetworkIDCarrierWifiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCwfNetworkIDCarrierWifiDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCwfNetworkIDCarrierWifiOK creates a PutCwfNetworkIDCarrierWifiOK with default headers values
func NewPutCwfNetworkIDCarrierWifiOK() *PutCwfNetworkIDCarrierWifiOK {
	return &PutCwfNetworkIDCarrierWifiOK{}
}

/*PutCwfNetworkIDCarrierWifiOK handles this case with default header values.

Success
*/
type PutCwfNetworkIDCarrierWifiOK struct {
}

func (o *PutCwfNetworkIDCarrierWifiOK) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/carrier_wifi][%d] putCwfNetworkIdCarrierWifiOK ", 200)
}

func (o *PutCwfNetworkIDCarrierWifiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCwfNetworkIDCarrierWifiDefault creates a PutCwfNetworkIDCarrierWifiDefault with default headers values
func NewPutCwfNetworkIDCarrierWifiDefault(code int) *PutCwfNetworkIDCarrierWifiDefault {
	return &PutCwfNetworkIDCarrierWifiDefault{
		_statusCode: code,
	}
}

/*PutCwfNetworkIDCarrierWifiDefault handles this case with default header values.

Unexpected Error
*/
type PutCwfNetworkIDCarrierWifiDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put cwf network ID carrier wifi default response
func (o *PutCwfNetworkIDCarrierWifiDefault) Code() int {
	return o._statusCode
}

func (o *PutCwfNetworkIDCarrierWifiDefault) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/carrier_wifi][%d] PutCwfNetworkIDCarrierWifi default  %+v", o._statusCode, o.Payload)
}

func (o *PutCwfNetworkIDCarrierWifiDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCwfNetworkIDCarrierWifiDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
