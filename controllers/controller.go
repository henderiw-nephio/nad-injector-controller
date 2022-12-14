/*
Copyright 2022 Nokia.

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

package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/henderiw-nephio/nad-injector-controller/controllers/injector"
	"github.com/henderiw-nephio/nad-injector-controller/pkg/shared"
)

// Setup package controllers.
func Setup(mgr ctrl.Manager, opts *shared.Options) error {
	for _, setup := range []func(ctrl.Manager, *shared.Options) error{
		injector.Setup,
	} {
		if err := setup(mgr, opts); err != nil {
			return err
		}
	}

	return nil
}
