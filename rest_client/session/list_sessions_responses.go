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

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// ListSessionsReader is a Reader for the ListSessions structure.
type ListSessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListSessionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSessionsOK creates a ListSessionsOK with default headers values
func NewListSessionsOK() *ListSessionsOK {
	return &ListSessionsOK{}
}

/*ListSessionsOK handles this case with default header values.

A list of sessions
*/
type ListSessionsOK struct {
	Payload *rest_model.ListSessionsEnvelope
}

func (o *ListSessionsOK) Error() string {
	return fmt.Sprintf("[GET /sessions][%d] listSessionsOK  %+v", 200, o.Payload)
}

func (o *ListSessionsOK) GetPayload() *rest_model.ListSessionsEnvelope {
	return o.Payload
}

func (o *ListSessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListSessionsEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSessionsUnauthorized creates a ListSessionsUnauthorized with default headers values
func NewListSessionsUnauthorized() *ListSessionsUnauthorized {
	return &ListSessionsUnauthorized{}
}

/*ListSessionsUnauthorized handles this case with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListSessionsUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListSessionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sessions][%d] listSessionsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListSessionsUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListSessionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
