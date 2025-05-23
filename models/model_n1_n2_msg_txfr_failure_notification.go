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

// Data within a N1/N2 Message Transfer Failure Notification request
type N1N2MsgTxfrFailureNotification struct {
	Cause N1N2MessageTransferCause `json:"cause" yaml:"cause" bson:"cause,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	N1n2MsgDataUri string `json:"n1n2MsgDataUri" yaml:"n1n2MsgDataUri" bson:"n1n2MsgDataUri,omitempty"`
}
