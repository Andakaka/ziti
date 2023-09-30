// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
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

package raft

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/ziti/controller/rest_model"
)

// RaftListMembersReader is a Reader for the RaftListMembers structure.
type RaftListMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RaftListMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRaftListMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRaftListMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRaftListMembersOK creates a RaftListMembersOK with default headers values
func NewRaftListMembersOK() *RaftListMembersOK {
	return &RaftListMembersOK{}
}

/* RaftListMembersOK describes a response with status code 200, with default header values.

A response to a raft list-members request
*/
type RaftListMembersOK struct {
	Payload *rest_model.RaftMemberListResponse
}

func (o *RaftListMembersOK) Error() string {
	return fmt.Sprintf("[GET /raft/list-members][%d] raftListMembersOK  %+v", 200, o.Payload)
}
func (o *RaftListMembersOK) GetPayload() *rest_model.RaftMemberListResponse {
	return o.Payload
}

func (o *RaftListMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.RaftMemberListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRaftListMembersUnauthorized creates a RaftListMembersUnauthorized with default headers values
func NewRaftListMembersUnauthorized() *RaftListMembersUnauthorized {
	return &RaftListMembersUnauthorized{}
}

/* RaftListMembersUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type RaftListMembersUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *RaftListMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /raft/list-members][%d] raftListMembersUnauthorized  %+v", 401, o.Payload)
}
func (o *RaftListMembersUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *RaftListMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
