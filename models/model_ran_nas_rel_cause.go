/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291  V17.0.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the RAN/NAS release cause.
type RanNasRelCause struct {
	NgApCause *NgApCause `json:"ngApCause,omitempty" yaml:"ngApCause" bson:"ngApCause,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gMmCause int32 `json:"5gMmCause,omitempty" yaml:"5gMmCause" bson:"5gMmCause,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gSmCause int32 `json:"5gSmCause,omitempty" yaml:"5gSmCause" bson:"5gSmCause,omitempty"`
	// Defines the EPS RAN/NAS release cause.
	EpsCause string `json:"epsCause,omitempty" yaml:"epsCause" bson:"epsCause,omitempty"`
}
