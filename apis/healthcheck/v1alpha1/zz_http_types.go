/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type HTTPObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HTTPParameters struct {

	// (String) Health check request HTTP body
	// Health check request HTTP body
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// (Set of String) A list of regular expression which will be executed against the response body
	// A list of regular expression which will be executed against the response body
	// +kubebuilder:validation:Optional
	BodyRegexp []*string `json:"bodyRegexp,omitempty" tf:"body_regexp,omitempty"`

	// (String) TLS cacert file to use for the TLS connection
	// TLS cacert file to use for the TLS connection
	// +kubebuilder:validation:Optional
	Cacert *string `json:"cacert,omitempty" tf:"cacert,omitempty"`

	// (String) TLS cert file to use for the TLS connection
	// TLS cert file to use for the TLS connection
	// +kubebuilder:validation:Optional
	Cert *string `json:"cert,omitempty" tf:"cert,omitempty"`

	// (String) Health check description
	// Health check description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Enable the health check on the Appclacks platform
	// Enable the health check on the Appclacks platform
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Map of String) Health check request HTTP headers
	// Health check request HTTP headers
	// +kubebuilder:validation:Optional
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// (String) Health check interval (example: 30s)
	// Health check interval (example: 30s)
	// +kubebuilder:validation:Optional
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// (String) TLS key file to use for the TLS connection
	// TLS key file to use for the TLS connection
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (Map of String) Health check labels
	// Health check labels
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) Health check HTTP method
	// Health check HTTP method
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// (String) Health check name
	// Health check name
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) Health check request HTTP path
	// Health check request HTTP path
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (Number) Health check port
	// Health check port
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// (String) Health check protocol to use (http or https)
	// Health check protocol to use (http or https)
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (Boolean) Follow redirections
	// Follow redirections
	// +kubebuilder:validation:Optional
	Redirect *bool `json:"redirect,omitempty" tf:"redirect,omitempty"`

	// (String) Health check target (can be a domain or an IP address)
	// Health check target (can be a domain or an IP address)
	// +kubebuilder:validation:Required
	Target *string `json:"target" tf:"target,omitempty"`

	// (String) Health check timeout (example: 5s)
	// Health check timeout (example: 5s)
	// +kubebuilder:validation:Optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// (Set of Number) Expected status code(s) for the HTTP response
	// Expected status code(s) for the HTTP response
	// +kubebuilder:validation:Required
	ValidStatus []*float64 `json:"validStatus" tf:"valid_status,omitempty"`
}

// HTTPSpec defines the desired state of HTTP
type HTTPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HTTPParameters `json:"forProvider"`
}

// HTTPStatus defines the observed state of HTTP.
type HTTPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HTTPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HTTP is the Schema for the HTTPs API. Execute an HTTP request on the target
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,appclacks}
type HTTP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HTTPSpec   `json:"spec"`
	Status            HTTPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HTTPList contains a list of HTTPs
type HTTPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTP `json:"items"`
}

// Repository type metadata.
var (
	HTTP_Kind             = "HTTP"
	HTTP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HTTP_Kind}.String()
	HTTP_KindAPIVersion   = HTTP_Kind + "." + CRDGroupVersion.String()
	HTTP_GroupVersionKind = CRDGroupVersion.WithKind(HTTP_Kind)
)

func init() {
	SchemeBuilder.Register(&HTTP{}, &HTTPList{})
}
