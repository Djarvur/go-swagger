// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/Djarvur/go-swagger/examples/task-tracker-strict-buildapi/models"
)

// GetTaskCommentsOKCode is the HTTP code returned for type GetTaskCommentsOK
const GetTaskCommentsOKCode int = 200

/*GetTaskCommentsOK The list of comments

swagger:response getTaskCommentsOK
*/
type GetTaskCommentsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Comment `json:"body,omitempty"`
}

// NewGetTaskCommentsOK creates GetTaskCommentsOK with default headers values
func NewGetTaskCommentsOK() *GetTaskCommentsOK {

	return &GetTaskCommentsOK{}
}

// WithPayload adds the payload to the get task comments o k response
func (o *GetTaskCommentsOK) WithPayload(payload []*models.Comment) *GetTaskCommentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get task comments o k response
func (o *GetTaskCommentsOK) SetPayload(payload []*models.Comment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTaskCommentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Comment, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *GetTaskCommentsOK) GetTaskCommentsResponder() {}

/*GetTaskCommentsDefault Error response

swagger:response getTaskCommentsDefault
*/
type GetTaskCommentsDefault struct {
	_statusCode int
	/*

	 */
	XErrorCode string `json:"X-Error-Code"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTaskCommentsDefault creates GetTaskCommentsDefault with default headers values
func NewGetTaskCommentsDefault(code int) *GetTaskCommentsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTaskCommentsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get task comments default response
func (o *GetTaskCommentsDefault) WithStatusCode(code int) *GetTaskCommentsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get task comments default response
func (o *GetTaskCommentsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithXErrorCode adds the xErrorCode to the get task comments default response
func (o *GetTaskCommentsDefault) WithXErrorCode(xErrorCode string) *GetTaskCommentsDefault {
	o.XErrorCode = xErrorCode
	return o
}

// SetXErrorCode sets the xErrorCode to the get task comments default response
func (o *GetTaskCommentsDefault) SetXErrorCode(xErrorCode string) {
	o.XErrorCode = xErrorCode
}

// WithPayload adds the payload to the get task comments default response
func (o *GetTaskCommentsDefault) WithPayload(payload *models.Error) *GetTaskCommentsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get task comments default response
func (o *GetTaskCommentsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTaskCommentsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Error-Code

	xErrorCode := o.XErrorCode
	if xErrorCode != "" {
		rw.Header().Set("X-Error-Code", xErrorCode)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetTaskCommentsDefault) GetTaskCommentsResponder() {}

type GetTaskCommentsNotImplementedResponder struct {
	middleware.Responder
}

func (*GetTaskCommentsNotImplementedResponder) GetTaskCommentsResponder() {}

func GetTaskCommentsNotImplemented() GetTaskCommentsResponder {
	return &GetTaskCommentsNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetTaskComments has not yet been implemented",
		),
	}
}

type GetTaskCommentsResponder interface {
	middleware.Responder
	GetTaskCommentsResponder()
}
