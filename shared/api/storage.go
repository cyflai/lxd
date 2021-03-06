package api

// StoragePool represents the fields of a LXD storage pool.
//
// API extension: storage
type StoragePool struct {
	StoragePoolPut `yaml:",inline"`

	Name   string   `json:"name" yaml:"name"`
	Driver string   `json:"driver" yaml:"driver"`
	UsedBy []string `json:"used_by" yaml:"used_by"`
}

// StoragePoolPut represents the modifiable fields of a LXD storage pool.
//
// API extension: storage
type StoragePoolPut struct {
	Config map[string]string `json:"config" yaml:"config"`
}

// StorageVolume represents the fields of a LXD storage volume.
//
// API extension: storage
type StorageVolume struct {
	StorageVolumePut `yaml:",inline"`

	Type   string   `json:"type" yaml:"type"`
	UsedBy []string `json:"used_by" yaml:"used_by"`
}

// StorageVolumePut represents the modifiable fields of a LXD storage volume.
//
// API extension: storage
type StorageVolumePut struct {
	Name   string            `json:"name" yaml:"name"`
	Config map[string]string `json:"config" yaml:"config"`
}

// Writable converts a full StoragePool struct into a StoragePoolPut struct
// (filters read-only fields).
func (storagePool *StoragePool) Writable() StoragePoolPut {
	return storagePool.StoragePoolPut
}

// Writable converts a full StorageVolume struct into a StorageVolumePut struct
// (filters read-only fields).
func (storageVolume *StorageVolume) Writable() StorageVolumePut {
	return storageVolume.StorageVolumePut
}
