/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains parameters that can be used to guide the URSP.
type UrspRuleRequest struct {
	TrafficDesc *TrafficDescriptorComponents `json:"trafficDesc,omitempty" yaml:"trafficDesc" bson:"trafficDesc,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RelatPrecedence int32 `json:"relatPrecedence,omitempty" yaml:"relatPrecedence" bson:"relatPrecedence,omitempty"`
	// Sets of parameters that may be used to guide the Route Selection Descriptors of the  URSP.
	RouteSelParamSets []RouteSelectionParameterSet `json:"routeSelParamSets,omitempty" yaml:"routeSelParamSets" bson:"routeSelParamSets,omitempty"`
}
