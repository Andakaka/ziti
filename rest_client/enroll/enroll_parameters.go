// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package enroll

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewEnrollParams creates a new EnrollParams object
// with the default values initialized.
func NewEnrollParams() *EnrollParams {
	var ()
	return &EnrollParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnrollParamsWithTimeout creates a new EnrollParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnrollParamsWithTimeout(timeout time.Duration) *EnrollParams {
	var ()
	return &EnrollParams{

		timeout: timeout,
	}
}

// NewEnrollParamsWithContext creates a new EnrollParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnrollParamsWithContext(ctx context.Context) *EnrollParams {
	var ()
	return &EnrollParams{

		Context: ctx,
	}
}

// NewEnrollParamsWithHTTPClient creates a new EnrollParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnrollParamsWithHTTPClient(client *http.Client) *EnrollParams {
	var ()
	return &EnrollParams{
		HTTPClient: client,
	}
}

/*EnrollParams contains all the parameters to send to the API endpoint
for the enroll operation typically these are written to a http.Request
*/
type EnrollParams struct {

	/*Token*/
	Token *strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enroll params
func (o *EnrollParams) WithTimeout(timeout time.Duration) *EnrollParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enroll params
func (o *EnrollParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enroll params
func (o *EnrollParams) WithContext(ctx context.Context) *EnrollParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enroll params
func (o *EnrollParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enroll params
func (o *EnrollParams) WithHTTPClient(client *http.Client) *EnrollParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enroll params
func (o *EnrollParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the enroll params
func (o *EnrollParams) WithToken(token *strfmt.UUID) *EnrollParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the enroll params
func (o *EnrollParams) SetToken(token *strfmt.UUID) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *EnrollParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Token != nil {

		// query param token
		var qrToken strfmt.UUID
		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken.String()
		if qToken != "" {
			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
