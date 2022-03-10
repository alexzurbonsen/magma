// Code generated by go-swagger; DO NOT EDIT.

package rating_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutNetworksNetworkIDRatingGroupsRatingGroupIDReader is a Reader for the PutNetworksNetworkIDRatingGroupsRatingGroupID structure.
type PutNetworksNetworkIDRatingGroupsRatingGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDRatingGroupsRatingGroupIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDRatingGroupsRatingGroupIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDRatingGroupsRatingGroupIDNoContent creates a PutNetworksNetworkIDRatingGroupsRatingGroupIDNoContent with default headers values
func NewPutNetworksNetworkIDRatingGroupsRatingGroupIDNoContent() *PutNetworksNetworkIDRatingGroupsRatingGroupIDNoContent {
	return &PutNetworksNetworkIDRatingGroupsRatingGroupIDNoContent{}
}

/*PutNetworksNetworkIDRatingGroupsRatingGroupIDNoContent handles this case with default header values.

Success
*/
type PutNetworksNetworkIDRatingGroupsRatingGroupIDNoContent struct {
}

func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/rating_groups/{rating_group_id}][%d] putNetworksNetworkIdRatingGroupsRatingGroupIdNoContent ", 204)
}

func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDRatingGroupsRatingGroupIDDefault creates a PutNetworksNetworkIDRatingGroupsRatingGroupIDDefault with default headers values
func NewPutNetworksNetworkIDRatingGroupsRatingGroupIDDefault(code int) *PutNetworksNetworkIDRatingGroupsRatingGroupIDDefault {
	return &PutNetworksNetworkIDRatingGroupsRatingGroupIDDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDRatingGroupsRatingGroupIDDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDRatingGroupsRatingGroupIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID rating groups rating group ID default response
func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/rating_groups/{rating_group_id}][%d] PutNetworksNetworkIDRatingGroupsRatingGroupID default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDRatingGroupsRatingGroupIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
