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

package sanitytest

import (
	"fmt"
	"os"
	"path"
	"testing"

	sanity "github.com/kubernetes-csi/csi-test/pkg/sanity"
	compute "google.golang.org/api/compute/v1"
	gce "sigs.k8s.io/gcp-compute-persistent-disk-csi-driver/pkg/gce-cloud-provider/compute"
	metadataservice "sigs.k8s.io/gcp-compute-persistent-disk-csi-driver/pkg/gce-cloud-provider/metadata"
	driver "sigs.k8s.io/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver"
	mountmanager "sigs.k8s.io/gcp-compute-persistent-disk-csi-driver/pkg/mount-manager"
)

func TestSanity(t *testing.T) {
	// Set up variables
	driverName := "test-driver"
	project := "test-project"
	zone := "test-zone"
	vendorVersion := "test-version"
	tmpDir := "/tmp/csi"
	endpoint := fmt.Sprintf("unix:%s/csi.sock", tmpDir)
	mountPath := path.Join(tmpDir, "mount")
	stagePath := path.Join(tmpDir, "stage")
	// Set up driver and env
	gceDriver := driver.GetGCEDriver()

	cloudProvider, err := gce.CreateFakeCloudProvider(project, zone, nil)
	if err != nil {
		t.Fatalf("Failed to get cloud provider: %v", err)
	}

	mounter := mountmanager.NewFakeSafeMounter()
	deviceUtils := mountmanager.NewFakeDeviceUtils()

	//Initialize GCE Driver
	err = gceDriver.SetupGCEDriver(cloudProvider, mounter, deviceUtils, metadataservice.NewFakeService(), driverName, vendorVersion)
	if err != nil {
		t.Fatalf("Failed to initialize GCE CSI Driver: %v", err)
	}

	instance := &compute.Instance{
		Name:  "test-name",
		Disks: []*compute.AttachedDisk{},
	}
	cloudProvider.InsertInstance(instance, "test-location", "test-name")

	err = os.MkdirAll(tmpDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create sanity temp working dir %s: %v", tmpDir, err)
	}

	defer func() {
		// Clean up tmp dir
		if err = os.RemoveAll(tmpDir); err != nil {
			t.Fatalf("Failed to clean up sanity temp working dir %s: %v", tmpDir, err)
		}
	}()

	go func() {
		gceDriver.Run(endpoint)
	}()

	// Run test
	config := &sanity.Config{
		TargetPath:  mountPath,
		StagingPath: stagePath,
		Address:     endpoint,
	}
	sanity.Test(t, config)
}
