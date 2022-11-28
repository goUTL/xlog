// Copyright 2022 Vasiliy Vdovin

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xlog

import (
	"log"
)

func isNotNil(v ...any) bool {

	for _, v := range v {
		if v == nil {
			return false
		}
	}
	return true
}

// Fatallf check all values and If all values are non nil - a call to log.Fatalf
func Fatallf(format string, v ...any) {

	is := isNotNil(v...)

	if !is {
		v = nil
	}

	if is {
		log.Fatalf(format, v...)
	}

}

// Fatalln check all values and If all values are non nil - a call to log.Fatalln
func Fatalln(v ...any) {

	is := isNotNil(v...)

	if !is {
		v = nil
	}

	if is {
		log.Fatalln(v...)
	}
}
