/*
Copyright 2018 The Kubernetes Authors.
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

// +build windows

package gceGCEDriver

import (
        "fmt"
        "google.golang.org/grpc/codes"
        "google.golang.org/grpc/status"
	winutils "github.com/ddebroy/csi-proxy/pkg/api/v1alpha1"
)

func GetDevicePath(ns *GCENodeServer, deviceName string, partition string, volumeKey string) (string, error) {
	diskNumber, err := winutils.GetDiskNumberWithID(deviceName)
	if err != nil {
		return "", status.Error(codes.Internal, fmt.Sprintf("Unable to find disk device with ID:", deviceName))
	}
	devicePath := fmt.Sprintf("%d", *diskNumber)
	return devicePath, nil
}