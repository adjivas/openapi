/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.7.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Represents a subscription to a single event.
type MlEventSubscription struct {
	MLEvent        NwdafEvent                     `json:"mLEvent" yaml:"mLEvent" bson:"mLEvent,omitempty"`
	MLEventFilter  *NwdafAnalyticsInfoEventFilter `json:"mLEventFilter" yaml:"mLEventFilter" bson:"mLEventFilter,omitempty"`
	TgtUe          *TargetUeInformation           `json:"tgtUe,omitempty" yaml:"tgtUe" bson:"tgtUe,omitempty"`
	MLTargetPeriod *TimeWindow                    `json:"mLTargetPeriod,omitempty" yaml:"mLTargetPeriod" bson:"mLTargetPeriod,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	ExpiryTime *time.Time `json:"expiryTime,omitempty" yaml:"expiryTime" bson:"expiryTime,omitempty"`
}
