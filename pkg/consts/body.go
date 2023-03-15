// Copyright 2023 BINARY Members
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package consts

const (
	BodyOK       = "200 OK"
	BodyCreated  = "201 Created"
	BodyAccepted = "202 Accepted"

	BodyMovedPermanently = "301 Moved Permanently"
	BodyFound            = "302 Found"
	BodySeeOther         = "303 See Other"
	BodyNotModified      = "304 Not Modified"

	BodyBadRequest       = "400 bad request"
	BodyForbidden        = "403 Forbidden"
	BodyNotFound         = "404 page not found"
	BodyMethodNotAllowed = "405 method not allowed"

	BodyInternalServerError = "500 Internal Server Error"
	BodyNotImplemented      = "501 Not Implemented"
	BodyBadGateway          = "502 Bad Gateway"
	BodyServiceUnavailable  = "503 Service Unavailable"
)
