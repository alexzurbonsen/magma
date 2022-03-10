// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutUserUsernameReader is a Reader for the PutUserUsername structure.
type PutUserUsernameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUserUsernameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutUserUsernameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutUserUsernameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutUserUsernameNoContent creates a PutUserUsernameNoContent with default headers values
func NewPutUserUsernameNoContent() *PutUserUsernameNoContent {
	return &PutUserUsernameNoContent{}
}

/*PutUserUsernameNoContent handles this case with default header values.

Success
*/
type PutUserUsernameNoContent struct {
}

func (o *PutUserUsernameNoContent) Error() string {
	return fmt.Sprintf("[PUT /user/{username}][%d] putUserUsernameNoContent ", 204)
}

func (o *PutUserUsernameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUserUsernameDefault creates a PutUserUsernameDefault with default headers values
func NewPutUserUsernameDefault(code int) *PutUserUsernameDefault {
	return &PutUserUsernameDefault{
		_statusCode: code,
	}
}

/*PutUserUsernameDefault handles this case with default header values.

Unexpected Error
*/
type PutUserUsernameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put user username default response
func (o *PutUserUsernameDefault) Code() int {
	return o._statusCode
}

func (o *PutUserUsernameDefault) Error() string {
	return fmt.Sprintf("[PUT /user/{username}][%d] PutUserUsername default  %+v", o._statusCode, o.Payload)
}

func (o *PutUserUsernameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutUserUsernameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutUserUsernameBody put user username body
swagger:model PutUserUsernameBody
*/
type PutUserUsernameBody struct {

	// password
	// Format: password
	Password strfmt.Password `json:"password,omitempty"`
}

// Validate validates this put user username body
func (o *PutUserUsernameBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutUserUsernameBody) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(o.Password) { // not required
		return nil
	}

	if err := validate.FormatOf("password"+"."+"password", "body", "password", o.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutUserUsernameBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutUserUsernameBody) UnmarshalBinary(b []byte) error {
	var res PutUserUsernameBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
