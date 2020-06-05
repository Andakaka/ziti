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

package transit_router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// CreateTransitRouterReader is a Reader for the CreateTransitRouter structure.
type CreateTransitRouterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTransitRouterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTransitRouterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTransitRouterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateTransitRouterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateTransitRouterOK creates a CreateTransitRouterOK with default headers values
func NewCreateTransitRouterOK() *CreateTransitRouterOK {
	return &CreateTransitRouterOK{}
}

/*CreateTransitRouterOK handles this case with default header values.

The create request was successful and the resource has been added at the following location
*/
type CreateTransitRouterOK struct {
	Payload *rest_model.CreateEnvelope
}

func (o *CreateTransitRouterOK) Error() string {
	return fmt.Sprintf("[POST /transit-routers][%d] createTransitRouterOK  %+v", 200, o.Payload)
}

func (o *CreateTransitRouterOK) GetPayload() *rest_model.CreateEnvelope {
	return o.Payload
}

func (o *CreateTransitRouterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.CreateEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransitRouterBadRequest creates a CreateTransitRouterBadRequest with default headers values
func NewCreateTransitRouterBadRequest() *CreateTransitRouterBadRequest {
	return &CreateTransitRouterBadRequest{}
}

/*CreateTransitRouterBadRequest handles this case with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type CreateTransitRouterBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateTransitRouterBadRequest) Error() string {
	return fmt.Sprintf("[POST /transit-routers][%d] createTransitRouterBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTransitRouterBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateTransitRouterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransitRouterUnauthorized creates a CreateTransitRouterUnauthorized with default headers values
func NewCreateTransitRouterUnauthorized() *CreateTransitRouterUnauthorized {
	return &CreateTransitRouterUnauthorized{}
}

/*CreateTransitRouterUnauthorized handles this case with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type CreateTransitRouterUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateTransitRouterUnauthorized) Error() string {
	return fmt.Sprintf("[POST /transit-routers][%d] createTransitRouterUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateTransitRouterUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateTransitRouterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
