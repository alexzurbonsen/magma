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

// DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDReader is a Reader for the DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleID structure.
type DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDNoContent creates a DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDNoContent with default headers values
func NewDeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDNoContent() *DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDNoContent {
	return &DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDNoContent{}
}

/*DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDNoContent handles this case with default header values.

Success
*/
type DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDNoContent struct {
}

func (o *DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /feg_lte/{network_id}/subscriber_config/rule_names/{rule_id}][%d] deleteFegLteNetworkIdSubscriberConfigRuleNamesRuleIdNoContent ", 204)
}

func (o *DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault creates a DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault with default headers values
func NewDeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault(code int) *DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault {
	return &DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault{
		_statusCode: code,
	}
}

/*DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault handles this case with default header values.

Unexpected Error
*/
type DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete feg LTE network ID subscriber config rule names rule ID default response
func (o *DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /feg_lte/{network_id}/subscriber_config/rule_names/{rule_id}][%d] DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFegLTENetworkIDSubscriberConfigRuleNamesRuleIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
