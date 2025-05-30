/*
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.122 V17.9.0 T8 reference point for Northbound APIs
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.122/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UserPlaneEvent string

// List of UserPlaneEvent
const (
	UserPlaneEvent_SESSION_TERMINATION             UserPlaneEvent = "SESSION_TERMINATION"
	UserPlaneEvent_LOSS_OF_BEARER                  UserPlaneEvent = "LOSS_OF_BEARER"
	UserPlaneEvent_RECOVERY_OF_BEARER              UserPlaneEvent = "RECOVERY_OF_BEARER"
	UserPlaneEvent_RELEASE_OF_BEARER               UserPlaneEvent = "RELEASE_OF_BEARER"
	UserPlaneEvent_USAGE_REPORT                    UserPlaneEvent = "USAGE_REPORT"
	UserPlaneEvent_FAILED_RESOURCES_ALLOCATION     UserPlaneEvent = "FAILED_RESOURCES_ALLOCATION"
	UserPlaneEvent_QOS_GUARANTEED                  UserPlaneEvent = "QOS_GUARANTEED"
	UserPlaneEvent_QOS_NOT_GUARANTEED              UserPlaneEvent = "QOS_NOT_GUARANTEED"
	UserPlaneEvent_QOS_MONITORING                  UserPlaneEvent = "QOS_MONITORING"
	UserPlaneEvent_SUCCESSFUL_RESOURCES_ALLOCATION UserPlaneEvent = "SUCCESSFUL_RESOURCES_ALLOCATION"
	UserPlaneEvent_ACCESS_TYPE_CHANGE              UserPlaneEvent = "ACCESS_TYPE_CHANGE"
	UserPlaneEvent_PLMN_CHG                        UserPlaneEvent = "PLMN_CHG"
)
