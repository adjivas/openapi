/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Subscription to a set of NF Instances (DCCFs), identified by NF types, NF Set Id(s) or DCCF Serving Area information, i.e. list of TAIs served by the DCCF
type DccfCond struct {
	ConditionType      string                  `json:"conditionType" yaml:"conditionType" bson:"conditionType,omitempty"`
	TaiList            []Tai                   `json:"taiList,omitempty" yaml:"taiList" bson:"taiList,omitempty"`
	TaiRangeList       []TaiRange              `json:"taiRangeList,omitempty" yaml:"taiRangeList" bson:"taiRangeList,omitempty"`
	ServingNfTypeList  []NrfNfManagementNfType `json:"servingNfTypeList,omitempty" yaml:"servingNfTypeList" bson:"servingNfTypeList,omitempty"`
	ServingNfSetIdList []string                `json:"servingNfSetIdList,omitempty" yaml:"servingNfSetIdList" bson:"servingNfSetIdList,omitempty"`
}
