// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// DeleteUserUsernameReader is a Reader for the DeleteUserUsername structure.
type DeleteUserUsernameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserUsernameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserUsernameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteUserUsernameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteUserUsernameNoContent creates a DeleteUserUsernameNoContent with default headers values
func NewDeleteUserUsernameNoContent() *DeleteUserUsernameNoContent {
	return &DeleteUserUsernameNoContent{}
}

/* DeleteUserUsernameNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteUserUsernameNoContent struct {
}

func (o *DeleteUserUsernameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /user/{username}][%d] deleteUserUsernameNoContent ", 204)
}

func (o *DeleteUserUsernameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserUsernameDefault creates a DeleteUserUsernameDefault with default headers values
func NewDeleteUserUsernameDefault(code int) *DeleteUserUsernameDefault {
	return &DeleteUserUsernameDefault{
		_statusCode: code,
	}
}

/* DeleteUserUsernameDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteUserUsernameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete user username default response
func (o *DeleteUserUsernameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteUserUsernameDefault) Error() string {
	return fmt.Sprintf("[DELETE /user/{username}][%d] DeleteUserUsername default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteUserUsernameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserUsernameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
