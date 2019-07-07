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

package mountmanager

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ddebroy/csi-proxy/pkg/api/v1alpha1"
	"k8s.io/kubernetes/pkg/util/mount"
)

var _ mount.Interface = &ProxyMounter{}

type ProxyMounter struct {
	dummy int
}

func (mounter *ProxyMounter) formatAndMount(source string, target string, fstype string, options []string) error {
	diskNumber, err := strconv.Atoi(source)
	if err != nil {
		return err
	}
	err = v1alpha1.FormatAndMountDisk(diskNumber, "NTFS", target)
	return err
}

func (mounter *ProxyMounter) Mount(source string, target string, fstype string, options []string) error {
	return fmt.Errorf("Mount not implementend for ProxyMounter")
}

func (mounter *ProxyMounter) Unmount(target string) error {
	return fmt.Errorf("Unmount not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) List() ([]mount.MountPoint, error) {
	return []mount.MountPoint{}, nil
}

func (mounter *ProxyMounter) IsMountPointMatch(mp mount.MountPoint, dir string) bool {
	return mp.Path == dir
}

func (mounter *ProxyMounter) IsLikelyNotMountPoint(file string) (bool, error) {
	return false, fmt.Errorf("IsLikelyNotMountPoint not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) GetMountRefs(pathname string) ([]string, error) {
	return []string{}, fmt.Errorf("GetMountRefs not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) GetDeviceNameFromMount(mountPath, pluginMountDir string) (string, error) {
	return "", fmt.Errorf("GetDeviceNameFromMount not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) CleanSubPaths(podDir string, volumeName string) error {
	return fmt.Errorf("CleanSubPaths not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) DeviceOpened(pathname string) (bool, error) {
	return false, fmt.Errorf("DeviceOpened not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) MakeDir(pathname string) error {
	return fmt.Errorf("MakeDir not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) MakeFile(pathname string) error {
	return fmt.Errorf("MakeFile not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) ExistsPath(pathname string) (bool, error) {
	return false, fmt.Errorf("ExistsPath not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) GetFSGroup(pathname string) (int64, error) {
	return -1, fmt.Errorf("GetFSGroup not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) GetSELinuxSupport(pathname string) (bool, error) {
	return false, fmt.Errorf("GetSELinuxSupport not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) GetMode(pathname string) (os.FileMode, error) {
	return 0, fmt.Errorf("GetMode not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) GetFileType(pathname string) (mount.FileType, error) {
	return mount.FileType("fake"), fmt.Errorf("GetFileType not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) IsNotMountPoint(dir string) (bool, error) {
	return false, fmt.Errorf("IsNotMountPoint not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) MakeRShared(path string) error {
	return fmt.Errorf("MakeRShared not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) PathIsDevice(pathname string) (bool, error) {
	return false, fmt.Errorf("PathIsDevice not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) PrepareSafeSubpath(subPath mount.Subpath) (newHostPath string, cleanupAction func(), err error) {
	return subPath.Path, nil, fmt.Errorf("PrepareSafeSubpath  not implemented for ProxyMounter")
}

func (mounter *ProxyMounter) SafeMakeDir(pathname string, base string, perm os.FileMode) error {
	return fmt.Errorf("SafeMakeDir  not implemented for ProxyMounter")
}

func NewSafeMounter() *mount.SafeFormatAndMount {
	proxyMounter := &ProxyMounter{0}
	return &mount.SafeFormatAndMount{
		Interface: proxyMounter,
	}
}
