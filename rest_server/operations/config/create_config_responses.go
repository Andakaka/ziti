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

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// CreateConfigOKCode is the HTTP code returned for type CreateConfigOK
const CreateConfigOKCode int = 200

/*CreateConfigOK The create request was successful and the resource has been added at the following location

swagger:response createConfigOK
*/
type CreateConfigOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.CreateEnvelope `json:"body,omitempty"`
}

// NewCreateConfigOK creates CreateConfigOK with default headers values
func NewCreateConfigOK() *CreateConfigOK {

	return &CreateConfigOK{}
}

// WithPayload adds the payload to the create config o k response
func (o *CreateConfigOK) WithPayload(payload *rest_model.CreateEnvelope) *CreateConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create config o k response
func (o *CreateConfigOK) SetPayload(payload *rest_model.CreateEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateConfigBadRequestCode is the HTTP code returned for type CreateConfigBadRequest
const CreateConfigBadRequestCode int = 400

/*CreateConfigBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response createConfigBadRequest
*/
type CreateConfigBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateConfigBadRequest creates CreateConfigBadRequest with default headers values
func NewCreateConfigBadRequest() *CreateConfigBadRequest {

	return &CreateConfigBadRequest{}
}

// WithPayload adds the payload to the create config bad request response
func (o *CreateConfigBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateConfigBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create config bad request response
func (o *CreateConfigBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConfigBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateConfigUnauthorizedCode is the HTTP code returned for type CreateConfigUnauthorized
const CreateConfigUnauthorizedCode int = 401

/*CreateConfigUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response createConfigUnauthorized
*/
type CreateConfigUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateConfigUnauthorized creates CreateConfigUnauthorized with default headers values
func NewCreateConfigUnauthorized() *CreateConfigUnauthorized {

	return &CreateConfigUnauthorized{}
}

// WithPayload adds the payload to the create config unauthorized response
func (o *CreateConfigUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateConfigUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create config unauthorized response
func (o *CreateConfigUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConfigUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
