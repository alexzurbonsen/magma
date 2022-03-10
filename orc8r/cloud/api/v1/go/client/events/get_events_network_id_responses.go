// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetEventsNetworkIDReader is a Reader for the GetEventsNetworkID structure.
type GetEventsNetworkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventsNetworkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventsNetworkIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEventsNetworkIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEventsNetworkIDOK creates a GetEventsNetworkIDOK with default headers values
func NewGetEventsNetworkIDOK() *GetEventsNetworkIDOK {
	return &GetEventsNetworkIDOK{}
}

/*GetEventsNetworkIDOK handles this case with default header values.

Success
*/
type GetEventsNetworkIDOK struct {
	Payload *GetEventsNetworkIDOKBodyTuple0
}

func (o *GetEventsNetworkIDOK) Error() string {
	return fmt.Sprintf("[GET /events/{network_id}][%d] getEventsNetworkIdOK  %+v", 200, o.Payload)
}

func (o *GetEventsNetworkIDOK) GetPayload() *GetEventsNetworkIDOKBodyTuple0 {
	return o.Payload
}

func (o *GetEventsNetworkIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEventsNetworkIDOKBodyTuple0)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsNetworkIDDefault creates a GetEventsNetworkIDDefault with default headers values
func NewGetEventsNetworkIDDefault(code int) *GetEventsNetworkIDDefault {
	return &GetEventsNetworkIDDefault{
		_statusCode: code,
	}
}

/*GetEventsNetworkIDDefault handles this case with default header values.

Unexpected Error
*/
type GetEventsNetworkIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get events network ID default response
func (o *GetEventsNetworkIDDefault) Code() int {
	return o._statusCode
}

func (o *GetEventsNetworkIDDefault) Error() string {
	return fmt.Sprintf("[GET /events/{network_id}][%d] GetEventsNetworkID default  %+v", o._statusCode, o.Payload)
}

func (o *GetEventsNetworkIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEventsNetworkIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetEventsNetworkIDOKBodyTuple0 GetEventsNetworkIDOKBodyTuple0 a representation of an anonymous Tuple type
swagger:model GetEventsNetworkIDOKBodyTuple0
*/
type GetEventsNetworkIDOKBodyTuple0 struct {

	// p0
	// Required: true
	P0 *models.Event `json:"-"` // custom serializer

}

// UnmarshalJSON unmarshals this tuple type from a JSON array
func (o *GetEventsNetworkIDOKBodyTuple0) UnmarshalJSON(raw []byte) error {
	// stage 1, get the array but just the array
	var stage1 []json.RawMessage
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&stage1); err != nil {
		return err
	}

	// stage 2: hydrates struct members with tuple elements
	if len(stage1) > 0 {
		var dataP0 models.Event
		buf = bytes.NewBuffer(stage1[0])
		dec := json.NewDecoder(buf)
		dec.UseNumber()
		if err := dec.Decode(&dataP0); err != nil {
			return err
		}
		o.P0 = &dataP0

	}

	return nil
}

// MarshalJSON marshals this tuple type into a JSON array
func (o GetEventsNetworkIDOKBodyTuple0) MarshalJSON() ([]byte, error) {
	data := []interface{}{
		o.P0,
	}

	return json.Marshal(data)
}

// Validate validates this get events network ID o k body tuple0
func (o *GetEventsNetworkIDOKBodyTuple0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateP0(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEventsNetworkIDOKBodyTuple0) validateP0(formats strfmt.Registry) error {

	if err := validate.Required("P0", "body", o.P0); err != nil {
		return err
	}

	if o.P0 != nil {
		if err := o.P0.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("P0")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEventsNetworkIDOKBodyTuple0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEventsNetworkIDOKBodyTuple0) UnmarshalBinary(b []byte) error {
	var res GetEventsNetworkIDOKBodyTuple0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
