/*
Copyright 2021 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testingUtil

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()
	if actual == expected {
		return
	}
	t.Fatalf("expected does not equal actual. \nExpected: %v\nActual: %v", expected, actual)
}

func AssertDeepEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()
	if reflect.DeepEqual(actual, expected) {
		return
	}
	t.Fatalf("expected does not deep equal actual. \nExpected: %T %+v\nActual:   %T %+v", expected, expected, actual, actual)
}
