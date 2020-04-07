package alluxio

var ufsPrefix = "ufs"
var getMountTable = "mountTable"

type MountPointInfo struct {
	UfsUrl           string            `json:"ufsUrl"`
	UfsType          string            `json:"ufsType"`
	UfsCapacityBytes int64             `json:"ufsCapacityBytes"`
	UfsUsedBytes     int64             `json:"ufsUsedBytes"`
	ReadOnly         bool              `json:"readOnly"`
	Shared           bool              `json:"shared"`
	Properties       map[string]string `json:"properties"`
}

// CreateDirectory creates a directory.
func (client *Client) GetMountTable() (map[string]MountPointInfo, error) {
	result := make(map[string]MountPointInfo)
	err := client.get(join(ufsPrefix, getMountTable), noParams, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
