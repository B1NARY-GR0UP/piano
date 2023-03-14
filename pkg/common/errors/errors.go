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

package errors

import "errors"

var _ error = (*Error)(nil)

type Error struct {
	Err  error
	Type ErrorType
	Meta any
}

func (e *Error) Error() string {
	return e.Err.Error()
}

type ErrorType uint64

const (
	PrivateType ErrorType = iota
	PublicType
	OtherType
)

type ErrorChain []*Error

func New(err error, t ErrorType, meta any) *Error {
	return &Error{
		Err:  err,
		Type: t,
		Meta: meta,
	}
}

func (e *Error) SetErr(err error) {
	e.Err = err
}

func (e *Error) SetType(t ErrorType) {
	e.Type = t
}

func (e *Error) SetMeta(meta any) {
	e.Meta = meta
}

func NewPublic(err string) *Error {
	return New(errors.New(err), PublicType, nil)
}

func NewPrivate(err string) *Error {
	return New(errors.New(err), PrivateType, nil)
}

func NewOther(err string) *Error {
	return New(errors.New(err), OtherType, nil)
}

func (e *Error) IsType(t ErrorType) bool {
	return e.Type == t
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (ec ErrorChain) Errors() []string {
	if len(ec) == 0 {
		return nil
	}
	errs := make([]string, 0)
	for _, e := range ec {
		errs = append(errs, e.Error())
	}
	return errs
}
