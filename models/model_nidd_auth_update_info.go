/*
 * Nudm_NIDDAU
 *
 * Nudm NIDD Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.8.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents NIDD authorization update information.
type NiddAuthUpdateInfo struct {
	AuthorizationData *UdmNiddauAuthorizationData `json:"authorizationData" yaml:"authorizationData" bson:"authorizationData,omitempty"`
	InvalidityInd     bool                        `json:"invalidityInd,omitempty" yaml:"invalidityInd" bson:"invalidityInd,omitempty"`
	Snssai            *Snssai                     `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn       string    `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	NiddCause NiddCause `json:"niddCause,omitempty" yaml:"niddCause" bson:"niddCause,omitempty"`
}
