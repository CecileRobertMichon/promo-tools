/*
Copyright 2020 The Kubernetes Authors.

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

package report

import (
	"context"

	"cloud.google.com/go/errorreporting"
	"k8s.io/klog"
)

// NewGcpErrorReportingClient returns a new Stackdriver Error Reporting client.
func NewGcpErrorReportingClient(
	projectID, serviceName string,
) *errorreporting.Client {

	ctx := context.Background()

	erc, err := errorreporting.NewClient(ctx, projectID, errorreporting.Config{
		ServiceName: serviceName,
		OnError: func(err error) {
			klog.Errorf("Could not log error: %v", err)
		},
	})
	if err != nil {
		klog.Fatalf("Failed to create errorreporting client: %v", err)
	}

	return erc
}
