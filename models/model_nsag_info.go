/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.531 V17.7.0; 5G System; Network Slice Selection Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.531/
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the association of NSAGs and S-NSSAI(s) along with the TA(s) within which the association is valid.
type NsagInfo struct {
	NsagIds      []int32    `json:"nsagIds" yaml:"nsagIds" bson:"nsagIds,omitempty"`
	SnssaiList   []Snssai   `json:"snssaiList" yaml:"snssaiList" bson:"snssaiList,omitempty"`
	TaiList      []Tai      `json:"taiList,omitempty" yaml:"taiList" bson:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty" yaml:"taiRangeList" bson:"taiRangeList,omitempty"`
}
