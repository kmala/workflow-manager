package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/deis/workflow-manager/pkg/swagger/models"
)

// PublishComponentReleaseReader is a Reader for the PublishComponentRelease structure.
type PublishComponentReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PublishComponentReleaseReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPublishComponentReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPublishComponentReleaseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPublishComponentReleaseOK creates a PublishComponentReleaseOK with default headers values
func NewPublishComponentReleaseOK() *PublishComponentReleaseOK {
	return &PublishComponentReleaseOK{}
}

/*PublishComponentReleaseOK handles this case with default header values.

publish component release response
*/
type PublishComponentReleaseOK struct {
	Payload *models.ComponentDetail
}

func (o *PublishComponentReleaseOK) Error() string {
	return fmt.Sprintf("[POST /versions/{train}/{component}/{release}][%d] publishComponentReleaseOK  %+v", 200, o.Payload)
}

func (o *PublishComponentReleaseOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublishComponentReleaseDefault creates a PublishComponentReleaseDefault with default headers values
func NewPublishComponentReleaseDefault(code int) *PublishComponentReleaseDefault {
	return &PublishComponentReleaseDefault{
		_statusCode: code,
	}
}

/*PublishComponentReleaseDefault handles this case with default header values.

unexpected error
*/
type PublishComponentReleaseDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the publish component release default response
func (o *PublishComponentReleaseDefault) Code() int {
	return o._statusCode
}

func (o *PublishComponentReleaseDefault) Error() string {
	return fmt.Sprintf("[POST /versions/{train}/{component}/{release}][%d] publishComponentRelease default  %+v", o._statusCode, o.Payload)
}

func (o *PublishComponentReleaseDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
