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

// CreateTaskCreatedCode is the HTTP code returned for type CreateTaskCreated
const CreateTaskCreatedCode int = 201

/*CreateTaskCreated Task created

swagger:response createTaskCreated
*/
type CreateTaskCreated struct {
}

// NewCreateTaskCreated creates CreateTaskCreated with default headers values
func NewCreateTaskCreated() *CreateTaskCreated {

	return &CreateTaskCreated{}
}

// WriteResponse to the client
func (o *CreateTaskCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

func (o *CreateTaskCreated) CreateTaskResponder() {}

/*CreateTaskDefault Error response

swagger:response createTaskDefault
*/
type CreateTaskDefault struct {
	_statusCode int
	/*

	 */
	XErrorCode string `json:"X-Error-Code"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTaskDefault creates CreateTaskDefault with default headers values
func NewCreateTaskDefault(code int) *CreateTaskDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateTaskDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create task default response
func (o *CreateTaskDefault) WithStatusCode(code int) *CreateTaskDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create task default response
func (o *CreateTaskDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithXErrorCode adds the xErrorCode to the create task default response
func (o *CreateTaskDefault) WithXErrorCode(xErrorCode string) *CreateTaskDefault {
	o.XErrorCode = xErrorCode
	return o
}

// SetXErrorCode sets the xErrorCode to the create task default response
func (o *CreateTaskDefault) SetXErrorCode(xErrorCode string) {
	o.XErrorCode = xErrorCode
}

// WithPayload adds the payload to the create task default response
func (o *CreateTaskDefault) WithPayload(payload *models.Error) *CreateTaskDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create task default response
func (o *CreateTaskDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTaskDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

func (o *CreateTaskDefault) CreateTaskResponder() {}

type CreateTaskNotImplementedResponder struct {
	middleware.Responder
}

func (*CreateTaskNotImplementedResponder) CreateTaskResponder() {}

func CreateTaskNotImplemented() CreateTaskResponder {
	return &CreateTaskNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.CreateTask has not yet been implemented",
		),
	}
}

type CreateTaskResponder interface {
	middleware.Responder
	CreateTaskResponder()
}
