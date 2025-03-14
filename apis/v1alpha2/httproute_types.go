/*
Copyright 2020 The Kubernetes Authors.

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

package v1alpha2

import (
	"sigs.k8s.io/gateway-api/apis/v1beta1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=gateway-api
// +kubebuilder:subresource:status
// +kubebuilder:deprecatedversion:warning="The v1alpha2 version of HTTPRoute has been deprecated and will be removed in a future release of the API. Please upgrade to v1beta1."
// +kubebuilder:printcolumn:name="Hostnames",type=string,JSONPath=`.spec.hostnames`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// HTTPRoute provides a way to route HTTP requests. This includes the capability
// to match requests by hostname, path, header, or query param. Filters can be
// used to specify additional processing steps. Backends specify where matching
// requests should be routed.
type HTTPRoute v1beta1.HTTPRoute

// +kubebuilder:object:root=true

// HTTPRouteList contains a list of HTTPRoute.
type HTTPRouteList v1beta1.HTTPRouteList

// HTTPRouteSpec defines the desired state of HTTPRoute
// +k8s:deepcopy-gen=false
type HTTPRouteSpec = v1beta1.HTTPRouteSpec

// HTTPRouteRule defines semantics for matching an HTTP request based on
// conditions (matches), processing it (filters), and forwarding the request to
// an API object (backendRefs).
// +k8s:deepcopy-gen=false
type HTTPRouteRule = v1beta1.HTTPRouteRule

// PathMatchType specifies the semantics of how HTTP paths should be compared.
// Valid PathMatchType values are:
//
// * "Exact"
// * "PathPrefix"
// * "RegularExpression"
//
// PathPrefix and Exact paths must be syntactically valid:
//
// - Must begin with the `/` character
// - Must not contain consecutive `/` characters (e.g. `/foo///`, `//`).
//
// Note that values may be added to this enum, implementations
// must ensure that unknown values will not cause a crash.
//
// Unknown values here must result in the implementation setting the
// Accepted Condition for the Route to `status: False`, with a
// Reason of `UnsupportedValue`.
//
// +kubebuilder:validation:Enum=Exact;PathPrefix;RegularExpression
// +k8s:deepcopy-gen=false
type PathMatchType = v1beta1.PathMatchType

const (
	// Matches the URL path exactly and with case sensitivity.
	PathMatchExact PathMatchType = "Exact"

	// Matches based on a URL path prefix split by `/`. Matching is
	// case sensitive and done on a path element by element basis. A
	// path element refers to the list of labels in the path split by
	// the `/` separator. When specified, a trailing `/` is ignored.
	//
	// For example, the paths `/abc`, `/abc/`, and `/abc/def` would all match
	// the prefix `/abc`, but the path `/abcd` would not.
	//
	// "PathPrefix" is semantically equivalent to the "Prefix" path type in the
	// Kubernetes Ingress API.
	PathMatchPathPrefix PathMatchType = "PathPrefix"

	// Matches if the URL path matches the given regular expression with
	// case sensitivity.
	//
	// Since `"RegularExpression"` has custom conformance, implementations
	// can support POSIX, PCRE, RE2 or any other regular expression dialect.
	// Please read the implementation's documentation to determine the supported
	// dialect.
	PathMatchRegularExpression PathMatchType = "RegularExpression"
)

// HTTPPathMatch describes how to select a HTTP route by matching the HTTP request path.
// +k8s:deepcopy-gen=false
type HTTPPathMatch = v1beta1.HTTPPathMatch

// HeaderMatchType specifies the semantics of how HTTP header values should be
// compared. Valid HeaderMatchType values are:
//
// * "Exact"
// * "RegularExpression"
//
// Note that values may be added to this enum, implementations
// must ensure that unknown values will not cause a crash.
//
// Unknown values here must result in the implementation setting the
// Accepted Condition for the Route to `status: False`, with a
// Reason of `UnsupportedValue`.
//
// +kubebuilder:validation:Enum=Exact;RegularExpression
// +k8s:deepcopy-gen=false
type HeaderMatchType = v1beta1.HeaderMatchType

// HeaderMatchType constants.
const (
	HeaderMatchExact             = v1beta1.HeaderMatchExact
	HeaderMatchRegularExpression = v1beta1.HeaderMatchRegularExpression
)

// HTTPHeaderName is the name of an HTTP header.
//
// Valid values include:
// * "Authorization"
// * "Set-Cookie"
//
// Invalid values include:
//
//   - ":method" - ":" is an invalid character. This means that HTTP/2 pseudo
//     headers are not currently supported by this type.
//
// * "/invalid" - "/" is an invalid character
//
// +kubebuilder:validation:MinLength=1
// +kubebuilder:validation:MaxLength=256
// +kubebuilder:validation:Pattern=`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`
// +k8s:deepcopy-gen=false
type HTTPHeaderName = v1beta1.HTTPHeaderName

// HTTPHeaderMatch describes how to select a HTTP route by matching HTTP request
// headers.
// +k8s:deepcopy-gen=false
type HTTPHeaderMatch = v1beta1.HTTPHeaderMatch

// QueryParamMatchType specifies the semantics of how HTTP query parameter
// values should be compared. Valid QueryParamMatchType values are:
//
// * "Exact"
// * "RegularExpression"
//
// Note that values may be added to this enum, implementations
// must ensure that unknown values will not cause a crash.
//
// Unknown values here must result in the implementation setting the
// Accepted Condition for the Route to `status: False`, with a
// Reason of `UnsupportedValue`.
//
// +kubebuilder:validation:Enum=Exact;RegularExpression
// +k8s:deepcopy-gen=false
type QueryParamMatchType = v1beta1.QueryParamMatchType

// QueryParamMatchType constants.
const (
	QueryParamMatchExact             QueryParamMatchType = "Exact"
	QueryParamMatchRegularExpression QueryParamMatchType = "RegularExpression"
)

// HTTPQueryParamMatch describes how to select a HTTP route by matching HTTP
// query parameters.
// +k8s:deepcopy-gen=false
type HTTPQueryParamMatch = v1beta1.HTTPQueryParamMatch

// HTTPMethod describes how to select a HTTP route by matching the HTTP
// method as defined by
// [RFC 7231](https://datatracker.ietf.org/doc/html/rfc7231#section-4) and
// [RFC 5789](https://datatracker.ietf.org/doc/html/rfc5789#section-2).
// The value is expected in upper case.
//
// Note that values may be added to this enum, implementations
// must ensure that unknown values will not cause a crash.
//
// Unknown values here must result in the implementation setting the
// Accepted Condition for the Route to `status: False`, with a
// Reason of `UnsupportedValue`.
//
// +kubebuilder:validation:Enum=GET;HEAD;POST;PUT;DELETE;CONNECT;OPTIONS;TRACE;PATCH
// +k8s:deepcopy-gen=false
type HTTPMethod = v1beta1.HTTPMethod

const (
	HTTPMethodGet     HTTPMethod = "GET"
	HTTPMethodHead    HTTPMethod = "HEAD"
	HTTPMethodPost    HTTPMethod = "POST"
	HTTPMethodPut     HTTPMethod = "PUT"
	HTTPMethodDelete  HTTPMethod = "DELETE"
	HTTPMethodConnect HTTPMethod = "CONNECT"
	HTTPMethodOptions HTTPMethod = "OPTIONS"
	HTTPMethodTrace   HTTPMethod = "TRACE"
	HTTPMethodPatch   HTTPMethod = "PATCH"
)

// HTTPRouteMatch defines the predicate used to match requests to a given
// action. Multiple match types are ANDed together, i.e. the match will
// evaluate to true only if all conditions are satisfied.
//
// For example, the match below will match a HTTP request only if its path
// starts with `/foo` AND it contains the `version: v1` header:
//
// ```
// match:
//
//	path:
//	  value: "/foo"
//	headers:
//	- name: "version"
//	  value "v1"
//
// ```
// +k8s:deepcopy-gen=false
type HTTPRouteMatch = v1beta1.HTTPRouteMatch

// HTTPRouteFilter defines processing steps that must be completed during the
// request or response lifecycle. HTTPRouteFilters are meant as an extension
// point to express processing that may be done in Gateway implementations. Some
// examples include request or response modification, implementing
// authentication strategies, rate-limiting, and traffic shaping. API
// guarantee/conformance is defined based on the type of the filter.
// +k8s:deepcopy-gen=false
type HTTPRouteFilter = v1beta1.HTTPRouteFilter

// HTTPRouteFilterType identifies a type of HTTPRoute filter.
// +k8s:deepcopy-gen=false
type HTTPRouteFilterType = v1beta1.HTTPRouteFilterType

const (
	// HTTPRouteFilterRequestHeaderModifier can be used to add or remove an HTTP
	// header from an HTTP request before it is sent to the upstream target.
	//
	// Support in HTTPRouteRule: Core
	//
	// Support in HTTPBackendRef: Extended
	HTTPRouteFilterRequestHeaderModifier HTTPRouteFilterType = "RequestHeaderModifier"

	// HTTPRouteFilterRequestRedirect can be used to redirect a request to
	// another location. This filter can also be used for HTTP to HTTPS
	// redirects. This may not be used on the same Route rule or BackendRef as a
	// URLRewrite filter.
	//
	// Support in HTTPRouteRule: Core
	//
	// Support in HTTPBackendRef: Extended
	HTTPRouteFilterRequestRedirect HTTPRouteFilterType = "RequestRedirect"

	// HTTPRouteFilterURLRewrite can be used to modify a request during
	// forwarding. At most one of these filters may be used on a Route rule.
	// This may not be used on the same Route rule or BackendRef as a
	// RequestRedirect filter.
	//
	// Support in HTTPRouteRule: Extended
	//
	// Support in HTTPBackendRef: Extended
	//
	// <gateway:experimental>
	HTTPRouteFilterURLRewrite HTTPRouteFilterType = "URLRewrite"

	// HTTPRouteFilterRequestMirror can be used to mirror HTTP requests to a
	// different backend. The responses from this backend MUST be ignored by
	// the Gateway.
	//
	// Support in HTTPRouteRule: Extended
	//
	// Support in HTTPBackendRef: Extended
	HTTPRouteFilterRequestMirror HTTPRouteFilterType = "RequestMirror"

	// HTTPRouteFilterExtensionRef should be used for configuring custom
	// HTTP filters.
	//
	// Support in HTTPRouteRule: Custom
	//
	// Support in HTTPBackendRef: Custom
	HTTPRouteFilterExtensionRef HTTPRouteFilterType = "ExtensionRef"
)

// HTTPHeader represents an HTTP Header name and value as defined by RFC 7230.
// +k8s:deepcopy-gen=false
type HTTPHeader = v1beta1.HTTPHeader

// HTTPHeaderFilter defines a filter that modifies the headers of an HTTP request
// or response.
// +k8s:deepcopy-gen=false
type HTTPHeaderFilter = v1beta1.HTTPHeaderFilter

// HTTPPathModifierType defines the type of path redirect or rewrite.
// +k8s:deepcopy-gen=false
type HTTPPathModifierType = v1beta1.HTTPPathModifierType

const (
	// This type of modifier indicates that the full path will be replaced
	// by the specified value.
	FullPathHTTPPathModifier HTTPPathModifierType = "ReplaceFullPath"

	// This type of modifier indicates that any prefix path matches will be
	// replaced by the substitution value. For example, a path with a prefix
	// match of "/foo" and a ReplacePrefixMatch substitution of "/bar" will have
	// the "/foo" prefix replaced with "/bar" in matching requests.
	//
	// Note that this matches the behavior of the PathPrefix match type. This
	// matches full path elements. A path element refers to the list of labels
	// in the path split by the `/` separator. When specified, a trailing `/` is
	// ignored. For example, the paths `/abc`, `/abc/`, and `/abc/def` would all
	// match the prefix `/abc`, but the path `/abcd` would not.
	PrefixMatchHTTPPathModifier HTTPPathModifierType = "ReplacePrefixMatch"
)

// HTTPPathModifier defines configuration for path modifiers.
// <gateway:experimental>
// +k8s:deepcopy-gen=false
type HTTPPathModifier = v1beta1.HTTPPathModifier

// HTTPRequestRedirect defines a filter that redirects a request. This filter
// MUST NOT be used on the same Route rule as a HTTPURLRewrite filter.
// +k8s:deepcopy-gen=false
type HTTPRequestRedirectFilter = v1beta1.HTTPRequestRedirectFilter

// HTTPURLRewriteFilter defines a filter that modifies a request during
// forwarding. At most one of these filters may be used on a Route rule. This
// MUST NOT be used on the same Route rule as a HTTPRequestRedirect filter.
//
// Support: Extended
//
// <gateway:experimental>
// +k8s:deepcopy-gen=false
type HTTPURLRewriteFilter = v1beta1.HTTPURLRewriteFilter

// HTTPRequestMirrorFilter defines configuration for the RequestMirror filter.
// +k8s:deepcopy-gen=false
type HTTPRequestMirrorFilter = v1beta1.HTTPRequestMirrorFilter

// HTTPBackendRef defines how a HTTPRoute should forward an HTTP request.
// +k8s:deepcopy-gen=false
type HTTPBackendRef = v1beta1.HTTPBackendRef

// HTTPRouteStatus defines the observed state of HTTPRoute.
// +k8s:deepcopy-gen=false
type HTTPRouteStatus = v1beta1.HTTPRouteStatus
