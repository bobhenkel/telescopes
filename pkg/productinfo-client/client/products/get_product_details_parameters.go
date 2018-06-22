// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetProductDetailsParams creates a new GetProductDetailsParams object
// with the default values initialized.
func NewGetProductDetailsParams() *GetProductDetailsParams {
	var ()
	return &GetProductDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductDetailsParamsWithTimeout creates a new GetProductDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductDetailsParamsWithTimeout(timeout time.Duration) *GetProductDetailsParams {
	var ()
	return &GetProductDetailsParams{

		timeout: timeout,
	}
}

// NewGetProductDetailsParamsWithContext creates a new GetProductDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProductDetailsParamsWithContext(ctx context.Context) *GetProductDetailsParams {
	var ()
	return &GetProductDetailsParams{

		Context: ctx,
	}
}

// NewGetProductDetailsParamsWithHTTPClient creates a new GetProductDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProductDetailsParamsWithHTTPClient(client *http.Client) *GetProductDetailsParams {
	var ()
	return &GetProductDetailsParams{
		HTTPClient: client,
	}
}

/*GetProductDetailsParams contains all the parameters to send to the API endpoint
for the get product details operation typically these are written to a http.Request
*/
type GetProductDetailsParams struct {

	/*Provider*/
	Provider string
	/*Region*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get product details params
func (o *GetProductDetailsParams) WithTimeout(timeout time.Duration) *GetProductDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get product details params
func (o *GetProductDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get product details params
func (o *GetProductDetailsParams) WithContext(ctx context.Context) *GetProductDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get product details params
func (o *GetProductDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get product details params
func (o *GetProductDetailsParams) WithHTTPClient(client *http.Client) *GetProductDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get product details params
func (o *GetProductDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProvider adds the provider to the get product details params
func (o *GetProductDetailsParams) WithProvider(provider string) *GetProductDetailsParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the get product details params
func (o *GetProductDetailsParams) SetProvider(provider string) {
	o.Provider = provider
}

// WithRegion adds the region to the get product details params
func (o *GetProductDetailsParams) WithRegion(region string) *GetProductDetailsParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get product details params
func (o *GetProductDetailsParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param provider
	if err := r.SetPathParam("provider", o.Provider); err != nil {
		return err
	}

	// path param region
	if err := r.SetPathParam("region", o.Region); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}