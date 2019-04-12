// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/Djarvur/go-swagger/examples/contributed-templates/stratoscale/models"
)

// NewOrderCreateParams creates a new OrderCreateParams object
// with the default values initialized.
func NewOrderCreateParams() *OrderCreateParams {
	var ()
	return &OrderCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrderCreateParamsWithTimeout creates a new OrderCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrderCreateParamsWithTimeout(timeout time.Duration) *OrderCreateParams {
	var ()
	return &OrderCreateParams{

		timeout: timeout,
	}
}

// NewOrderCreateParamsWithContext creates a new OrderCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrderCreateParamsWithContext(ctx context.Context) *OrderCreateParams {
	var ()
	return &OrderCreateParams{

		Context: ctx,
	}
}

// NewOrderCreateParamsWithHTTPClient creates a new OrderCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrderCreateParamsWithHTTPClient(client *http.Client) *OrderCreateParams {
	var ()
	return &OrderCreateParams{
		HTTPClient: client,
	}
}

/*OrderCreateParams contains all the parameters to send to the API endpoint
for the order create operation typically these are written to a http.Request
*/
type OrderCreateParams struct {

	/*Body
	  order placed for purchasing the pet

	*/
	Body *models.Order

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the order create params
func (o *OrderCreateParams) WithTimeout(timeout time.Duration) *OrderCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order create params
func (o *OrderCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order create params
func (o *OrderCreateParams) WithContext(ctx context.Context) *OrderCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order create params
func (o *OrderCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order create params
func (o *OrderCreateParams) WithHTTPClient(client *http.Client) *OrderCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order create params
func (o *OrderCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the order create params
func (o *OrderCreateParams) WithBody(body *models.Order) *OrderCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the order create params
func (o *OrderCreateParams) SetBody(body *models.Order) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *OrderCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
