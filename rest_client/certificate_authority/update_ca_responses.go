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

package certificate_authority

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// UpdateCaReader is a Reader for the UpdateCa structure.
type UpdateCaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateCaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCaOK creates a UpdateCaOK with default headers values
func NewUpdateCaOK() *UpdateCaOK {
	return &UpdateCaOK{}
}

/*UpdateCaOK handles this case with default header values.

The update request was successful and the resource has been altered
*/
type UpdateCaOK struct {
	Payload *rest_model.Empty
}

func (o *UpdateCaOK) Error() string {
	return fmt.Sprintf("[PUT /cas/{id}][%d] updateCaOK  %+v", 200, o.Payload)
}

func (o *UpdateCaOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *UpdateCaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCaBadRequest creates a UpdateCaBadRequest with default headers values
func NewUpdateCaBadRequest() *UpdateCaBadRequest {
	return &UpdateCaBadRequest{}
}

/*UpdateCaBadRequest handles this case with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type UpdateCaBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdateCaBadRequest) Error() string {
	return fmt.Sprintf("[PUT /cas/{id}][%d] updateCaBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCaBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateCaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCaUnauthorized creates a UpdateCaUnauthorized with default headers values
func NewUpdateCaUnauthorized() *UpdateCaUnauthorized {
	return &UpdateCaUnauthorized{}
}

/*UpdateCaUnauthorized handles this case with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type UpdateCaUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdateCaUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /cas/{id}][%d] updateCaUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateCaUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateCaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCaNotFound creates a UpdateCaNotFound with default headers values
func NewUpdateCaNotFound() *UpdateCaNotFound {
	return &UpdateCaNotFound{}
}

/*UpdateCaNotFound handles this case with default header values.

The requested resource does not exist
*/
type UpdateCaNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdateCaNotFound) Error() string {
	return fmt.Sprintf("[PUT /cas/{id}][%d] updateCaNotFound  %+v", 404, o.Payload)
}

func (o *UpdateCaNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateCaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
