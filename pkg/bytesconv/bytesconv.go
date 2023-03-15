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

package bytesconv

import (
	"reflect"
	"unsafe"
)

// B2S converts byte slice to a string without memory allocation
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ
//
// Note that it may break if string and/or slice header will change
// in the future go versions
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// S2B converts string to a byte slice without memory allocation
//
// Note that it may break if string and/or slice header will change
// in the future go versions
func S2B(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

// AppendUint appends n to dst and returns the extended dst
func AppendUint(dst []byte, n uint) []byte {
	var b [20]byte
	buf := b[:]
	i := len(buf)
	var q uint
	for n >= 10 {
		i--
		q = n / 10
		buf[i] = '0' + byte(n-q*10)
		n = q
	}
	i--
	buf[i] = '0' + byte(n)
	dst = append(dst, buf[i:]...)
	return dst
}
