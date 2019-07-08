/*
Copyright 2019 The Kubernetes Authors.
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

package v1alpha1

import (
	"encoding/json"
)

const (
	// CallGetDiskNumberWithId system call type
	CallGetDiskNumberWithId = 2000 + iota
	CallFormatAndMountDisk
	CallMkDir
	CallMount
)

// Operation is the enumerator defining the underlay system call
type Operation uint32

// Message is the basic type wrapper around a request/response from client
type Message struct {
	Type Operation
	Body json.RawMessage
}

// GetDiskNumberWithIdReq body of the request for GetDiskDeviceWithId
type GetDiskNumberWithIdReq struct {
	DeviceId string
}

// GetDiskNumberWithIdRsp body of the response for GetDiskDeviceWithId
type GetDiskNumberWithIdRsp struct {
	DiskNumber int
	Err        string
}

// FormatAndMountDiskReq body of request for FormatAndMountDisk
type FormatAndMountDiskReq struct {
	DiskNumber int
	FsType     string
	MountPath  string
}

// FormatAndMountDiskRsp body of response for FormatAndMountDisk
type FormatAndMountDiskRsp struct {
	Err string
}

// MkdirReq body of request for MkDir
type MkDirReq struct {
	Path string
}

// MkDirRsp body of response for MkDir
type MkDirRsp struct {
	Err string
}

// MountReq body of request for Mount
type MountReq struct {
	SourcePath string
	TargetPath string
	FsType     string
	Options    []string
}

// MountRsp body of response for Mount
type MountRsp struct {
	Err string
}
