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

package service_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// ListServicePoliciesOKCode is the HTTP code returned for type ListServicePoliciesOK
const ListServicePoliciesOKCode int = 200

/*ListServicePoliciesOK A list of service policies

swagger:response listServicePoliciesOK
*/
type ListServicePoliciesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListServicePoliciesEnvelope `json:"body,omitempty"`
}

// NewListServicePoliciesOK creates ListServicePoliciesOK with default headers values
func NewListServicePoliciesOK() *ListServicePoliciesOK {

	return &ListServicePoliciesOK{}
}

// WithPayload adds the payload to the list service policies o k response
func (o *ListServicePoliciesOK) WithPayload(payload *rest_model.ListServicePoliciesEnvelope) *ListServicePoliciesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service policies o k response
func (o *ListServicePoliciesOK) SetPayload(payload *rest_model.ListServicePoliciesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicePoliciesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServicePoliciesUnauthorizedCode is the HTTP code returned for type ListServicePoliciesUnauthorized
const ListServicePoliciesUnauthorizedCode int = 401

/*ListServicePoliciesUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listServicePoliciesUnauthorized
*/
type ListServicePoliciesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServicePoliciesUnauthorized creates ListServicePoliciesUnauthorized with default headers values
func NewListServicePoliciesUnauthorized() *ListServicePoliciesUnauthorized {

	return &ListServicePoliciesUnauthorized{}
}

// WithPayload adds the payload to the list service policies unauthorized response
func (o *ListServicePoliciesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServicePoliciesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service policies unauthorized response
func (o *ListServicePoliciesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicePoliciesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
