/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.10.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Traffic information including UL/DL data rate and/or Traffic volume.
type TrafficInformation struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	UplinkRate string `json:"uplinkRate,omitempty" yaml:"uplinkRate" bson:"uplinkRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	DownlinkRate string `json:"downlinkRate,omitempty" yaml:"downlinkRate" bson:"downlinkRate,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	UplinkVolume int64 `json:"uplinkVolume,omitempty" yaml:"uplinkVolume" bson:"uplinkVolume,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	DownlinkVolume int64 `json:"downlinkVolume,omitempty" yaml:"downlinkVolume" bson:"downlinkVolume,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	TotalVolume int64 `json:"totalVolume,omitempty" yaml:"totalVolume" bson:"totalVolume,omitempty"`
}
