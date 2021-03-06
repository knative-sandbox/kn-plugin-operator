//go:build beta
// +build beta

/*
Copyright 2022 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"testing"

	"knative.dev/kn-plugin-operator/test/resources"
	"knative.dev/operator/test/client"
)

// TestOperatorBeta verifies the installation of Knative Operator 1.3 or later
func TestOperatorBeta(t *testing.T) {
	clients := client.Setup(t)
	resources.VerifyOperatorInstallationBeta(t, clients)
}
