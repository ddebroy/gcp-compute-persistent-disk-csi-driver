// +build windows

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
	"fmt"
	"os"

	winio "github.com/Microsoft/go-winio"
)

const (
	pipeEnv = "CSI_PROXY_PIPE"
)

// GetDiskDeviceWithID gets disk device with ID
func GetDiskNumberWithID(storageId string) (*int, error) {
	r, w, err := NewClientPipe()
	if err != nil {
		return nil, err
	}

	if err = w.Write(CallGetDiskNumberWithId, &GetDiskNumberWithIdReq{storageId}); err != nil {
		return nil, err
	}

	var response GetDiskNumberWithIdRsp
	_, err = r.ReadWithBody(&response)
	if err != nil {
		return nil, err
	}

	if response.Err != "" {
		return nil, fmt.Errorf(response.Err)
	}
	return &response.DiskNumber, nil
}

// FormatAndMountDisk formats and mounts the disk with number
func FormatAndMountDisk(diskNumber int, fsType string, mountPath string) (error) {
	r, w, err := NewClientPipe()
	if err != nil {
		return err
	}
	
	if err = w.Write(CallFormatAndMountDisk, &FormatAndMountDiskReq{diskNumber, fsType, mountPath}); err != nil {
		return err
	}
	
	var response FormatAndMountDiskRsp
	_, err = r.ReadWithBody(&response)
	if err != nil {
		return err
	}
	
	if response.Err != "" {
		return fmt.Errorf(response.Err)
	}
	return nil
} 

// NewClientPipe opens the pipe with the name passed
func NewClientPipe() (*Reader, *Writer, error) {
	name := os.Getenv(pipeEnv)
	if name == "" {
		return nil, nil, fmt.Errorf("unable to find the pipe name")
	}
	c, err := winio.DialPipe(name, nil)
	if err != nil {
		return nil, nil, err
	}
	return NewReader(c), NewWriter(c), nil
}
