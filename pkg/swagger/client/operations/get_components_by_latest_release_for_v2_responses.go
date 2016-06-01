package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/deis/workflow-manager/pkg/swagger/models"
)

// GetComponentsByLatestReleaseForV2Reader is a Reader for the GetComponentsByLatestReleaseForV2 structure.
type GetComponentsByLatestReleaseForV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetComponentsByLatestReleaseForV2Reader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetComponentsByLatestReleaseForV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetComponentsByLatestReleaseForV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetComponentsByLatestReleaseForV2OK creates a GetComponentsByLatestReleaseForV2OK with default headers values
func NewGetComponentsByLatestReleaseForV2OK() *GetComponentsByLatestReleaseForV2OK {
	return &GetComponentsByLatestReleaseForV2OK{}
}

/*GetComponentsByLatestReleaseForV2OK handles this case with default header values.

component releases response
*/
type GetComponentsByLatestReleaseForV2OK struct {
	Payload GetComponentsByLatestReleaseForV2OKBodyBody
}

func (o *GetComponentsByLatestReleaseForV2OK) Error() string {
	return fmt.Sprintf("[POST /v2/versions/latest][%d] getComponentsByLatestReleaseForV2OK  %+v", 200, o.Payload)
}

func (o *GetComponentsByLatestReleaseForV2OK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentsByLatestReleaseForV2Default creates a GetComponentsByLatestReleaseForV2Default with default headers values
func NewGetComponentsByLatestReleaseForV2Default(code int) *GetComponentsByLatestReleaseForV2Default {
	return &GetComponentsByLatestReleaseForV2Default{
		_statusCode: code,
	}
}

/*GetComponentsByLatestReleaseForV2Default handles this case with default header values.

unexpected error
*/
type GetComponentsByLatestReleaseForV2Default struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get components by latest release for v2 default response
func (o *GetComponentsByLatestReleaseForV2Default) Code() int {
	return o._statusCode
}

func (o *GetComponentsByLatestReleaseForV2Default) Error() string {
	return fmt.Sprintf("[POST /v2/versions/latest][%d] getComponentsByLatestReleaseForV2 default  %+v", o._statusCode, o.Payload)
}

func (o *GetComponentsByLatestReleaseForV2Default) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetComponentsByLatestReleaseForV2Body get components by latest release for v2 body

swagger:model GetComponentsByLatestReleaseForV2Body
*/
type GetComponentsByLatestReleaseForV2Body struct {

	/* data
	 */
	Data []*models.ComponentVersion `json:"data,omitempty"`
}

// Validate validates this get components by latest release for v2 body
func (o *GetComponentsByLatestReleaseForV2Body) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetComponentsByLatestReleaseForV2Body) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if err := o.Data[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

/*GetComponentsByLatestReleaseForV2OKBodyBody get components by latest release for v2 o k body body

swagger:model GetComponentsByLatestReleaseForV2OKBodyBody
*/
type GetComponentsByLatestReleaseForV2OKBodyBody struct {

	/* data

	Required: true
	*/
	Data []*models.ComponentVersion `json:"data"`
}

// Validate validates this get components by latest release for v2 o k body body
func (o *GetComponentsByLatestReleaseForV2OKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetComponentsByLatestReleaseForV2OKBodyBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getComponentsByLatestReleaseForV2OK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if err := o.Data[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
