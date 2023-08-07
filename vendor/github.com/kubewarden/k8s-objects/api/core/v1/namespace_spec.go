// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// NamespaceSpec NamespaceSpec describes the attributes on a Namespace.
//
// swagger:model NamespaceSpec
type NamespaceSpec struct {

	// Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
	Finalizers []string `json:"finalizers,omitempty"`
}
