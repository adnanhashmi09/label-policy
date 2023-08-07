// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// NodeConfigSource NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil. This API is deprecated since 1.22
//
// swagger:model NodeConfigSource
type NodeConfigSource struct {

	// ConfigMap is a reference to a Node's ConfigMap
	ConfigMap *ConfigMapNodeConfigSource `json:"configMap,omitempty"`
}
