// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// CSIVolumeSource Represents a source location of a volume to mount, managed by an external CSI driver
//
// swagger:model CSIVolumeSource
type CSIVolumeSource struct {

	// driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.
	// Required: true
	Driver *string `json:"driver"`

	// fsType to mount. Ex. "ext4", "xfs", "ntfs". If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.
	FSType string `json:"fsType,omitempty"`

	// nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.
	NodePublishSecretRef *LocalObjectReference `json:"nodePublishSecretRef,omitempty"`

	// readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).
	ReadOnly bool `json:"readOnly,omitempty"`

	// volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.
	VolumeAttributes map[string]string `json:"volumeAttributes,omitempty"`
}
