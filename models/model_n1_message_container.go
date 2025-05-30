/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// N1 Message container
type N1MessageContainer struct {
	N1MessageClass   N1MessageClass   `json:"n1MessageClass" yaml:"n1MessageClass" bson:"n1MessageClass,omitempty"`
	N1MessageContent *RefToBinaryData `json:"n1MessageContent" yaml:"n1MessageContent" bson:"n1MessageContent,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfId              string `json:"nfId,omitempty" yaml:"nfId" bson:"nfId,omitempty"`
	ServiceInstanceId string `json:"serviceInstanceId,omitempty" yaml:"serviceInstanceId" bson:"serviceInstanceId,omitempty"`
}
