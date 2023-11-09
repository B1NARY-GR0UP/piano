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

package core

// Param is a single URL parameter
type Param struct {
	Key   string
	Value string
}

type Params []Param

// Get returns the value of the first Param which key matches the given key.
// If no matching Param is found, an empty string is returned.
func (ps *Params) Get(key string) string {
	for _, v := range *ps {
		if v.Key == key {
			return v.Value
		}
	}
	return ""
}

// Set will set key-value pair into Params
func (ps *Params) Set(key, value string) {
	*ps = append(*ps, Param{
		Key:   key,
		Value: value,
	})
}
