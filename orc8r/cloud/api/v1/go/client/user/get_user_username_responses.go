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

// GetUserUsernameReader is a Reader for the GetUserUsername structure.
type GetUserUsernameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserUsernameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGetUserUsernameCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUserUsernameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUserUsernameCreated creates a GetUserUsernameCreated with default headers values
func NewGetUserUsernameCreated() *GetUserUsernameCreated {
	return &GetUserUsernameCreated{}
}

/*GetUserUsernameCreated handles this case with default header values.

Success
*/
type GetUserUsernameCreated struct {
	Payload string
}

func (o *GetUserUsernameCreated) Error() string {
	return fmt.Sprintf("[GET /user/{username}][%d] getUserUsernameCreated  %+v", 201, o.Payload)
}

func (o *GetUserUsernameCreated) GetPayload() string {
	return o.Payload
}

func (o *GetUserUsernameCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserUsernameDefault creates a GetUserUsernameDefault with default headers values
func NewGetUserUsernameDefault(code int) *GetUserUsernameDefault {
	return &GetUserUsernameDefault{
		_statusCode: code,
	}
}

/*GetUserUsernameDefault handles this case with default header values.

Unexpected Error
*/
type GetUserUsernameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get user username default response
func (o *GetUserUsernameDefault) Code() int {
	return o._statusCode
}

func (o *GetUserUsernameDefault) Error() string {
	return fmt.Sprintf("[GET /user/{username}][%d] GetUserUsername default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserUsernameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserUsernameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
