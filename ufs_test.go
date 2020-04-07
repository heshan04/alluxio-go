package alluxio

import (
	"reflect"
	"testing"
)

func TestUFSGetMountTable(t *testing.T) {
	output := map[string]MountPointInfo{}
	output["/"] = MountPointInfo{
		UfsType: "local",
		UfsUrl: "/storage/default-volume",
		UfsCapacityBytes: 1000000,
		UfsUsedBytes: 200,
	}
	ufs, cleanup := setupClient(t, jsonHandler(t, join(getMountTable), ufsPrefix, nil, &output, nil))
	defer cleanup()
	if mountTable, err := ufs.GetMountTable(); err != nil {
		t.Fatalf("GetMountTable() failed: %v", err)
	} else if !reflect.DeepEqual(output, mountTable) {
		t.Fatalf("Unexpected output: got %#v, want %#v", mountTable, output)
	}
}
