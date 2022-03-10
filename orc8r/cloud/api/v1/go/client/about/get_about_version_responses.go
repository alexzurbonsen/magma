// Code generated by go-swagger; DO NOT EDIT.

package about

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetAboutVersionReader is a Reader for the GetAboutVersion structure.
type GetAboutVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAboutVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAboutVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAboutVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAboutVersionOK creates a GetAboutVersionOK with default headers values
func NewGetAboutVersionOK() *GetAboutVersionOK {
	return &GetAboutVersionOK{}
}

/*GetAboutVersionOK handles this case with default header values.

Version
*/
type GetAboutVersionOK struct {
	Payload *models.VersionInfo
}

func (o *GetAboutVersionOK) Error() string {
	return fmt.Sprintf("[GET /about/version][%d] getAboutVersionOK  %+v", 200, o.Payload)
}

func (o *GetAboutVersionOK) GetPayload() *models.VersionInfo {
	return o.Payload
}

func (o *GetAboutVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAboutVersionDefault creates a GetAboutVersionDefault with default headers values
func NewGetAboutVersionDefault(code int) *GetAboutVersionDefault {
	return &GetAboutVersionDefault{
		_statusCode: code,
	}
}

/*GetAboutVersionDefault handles this case with default header values.

Unexpected Error
*/
type GetAboutVersionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get about version default response
func (o *GetAboutVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetAboutVersionDefault) Error() string {
	return fmt.Sprintf("[GET /about/version][%d] GetAboutVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetAboutVersionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAboutVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
