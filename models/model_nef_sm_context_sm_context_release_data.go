/*
 * Nnef_SMContext
 *
 * Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.541 V17.3.0; 5G System; Session Management Services for Non-IP Data Delivery (NIDD).
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.541/
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Representation of the information to release the Individual SM context.
type NefSmContextSmContextReleaseData struct {
	Cause NefSmContextReleaseCause `json:"cause" yaml:"cause" bson:"cause,omitempty"`
}
