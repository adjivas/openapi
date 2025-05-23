/*
 * Namf_Location
 *
 * AMF Location Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Data within Provide Positioning Information Request
type RequestPosInfo struct {
	LcsClientType ExternalClientType      `json:"lcsClientType" yaml:"lcsClientType" bson:"lcsClientType,omitempty"`
	LcsLocation   AmfLocationLocationType `json:"lcsLocation" yaml:"lcsLocation" bson:"lcsLocation,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi string `json:"supi,omitempty" yaml:"supi" bson:"supi,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi                       string               `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	Priority                   LcsPriority          `json:"priority,omitempty" yaml:"priority" bson:"priority,omitempty"`
	LcsQoS                     *LocationQoS         `json:"lcsQoS,omitempty" yaml:"lcsQoS" bson:"lcsQoS,omitempty"`
	VelocityRequested          VelocityRequested    `json:"velocityRequested,omitempty" yaml:"velocityRequested" bson:"velocityRequested,omitempty"`
	LcsSupportedGADShapes      SupportedGadShapes   `json:"lcsSupportedGADShapes,omitempty" yaml:"lcsSupportedGADShapes" bson:"lcsSupportedGADShapes,omitempty"`
	AdditionalLcsSuppGADShapes []SupportedGadShapes `json:"additionalLcsSuppGADShapes,omitempty" yaml:"additionalLcsSuppGADShapes" bson:"additionalLcsSuppGADShapes,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	LocationNotificationUri string `json:"locationNotificationUri,omitempty" yaml:"locationNotificationUri" bson:"locationNotificationUri,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	OldGuami          *Guami `json:"oldGuami,omitempty" yaml:"oldGuami" bson:"oldGuami,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.
	Pei string `json:"pei,omitempty" yaml:"pei" bson:"pei,omitempty"`
	// LCS service type.
	LcsServiceType int32   `json:"lcsServiceType,omitempty" yaml:"lcsServiceType" bson:"lcsServiceType,omitempty"`
	LdrType        LdrType `json:"ldrType,omitempty" yaml:"ldrType" bson:"ldrType,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	HgmlcCallBackURI string `json:"hgmlcCallBackURI,omitempty" yaml:"hgmlcCallBackURI" bson:"hgmlcCallBackURI,omitempty"`
	// LDR Reference.
	LdrReference      string             `json:"ldrReference,omitempty" yaml:"ldrReference" bson:"ldrReference,omitempty"`
	PeriodicEventInfo *PeriodicEventInfo `json:"periodicEventInfo,omitempty" yaml:"periodicEventInfo" bson:"periodicEventInfo,omitempty"`
	AreaEventInfo     *AreaEventInfo     `json:"areaEventInfo,omitempty" yaml:"areaEventInfo" bson:"areaEventInfo,omitempty"`
	MotionEventInfo   *MotionEventInfo   `json:"motionEventInfo,omitempty" yaml:"motionEventInfo" bson:"motionEventInfo,omitempty"`
	// Contains the external client identification
	ExternalClientIdentification string `json:"externalClientIdentification,omitempty" yaml:"externalClientIdentification" bson:"externalClientIdentification,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	AfID string `json:"afID,omitempty" yaml:"afID" bson:"afID,omitempty"`
	// Contains the codeword
	CodeWord              string                 `json:"codeWord,omitempty" yaml:"codeWord" bson:"codeWord,omitempty"`
	UePrivacyRequirements *UePrivacyRequirements `json:"uePrivacyRequirements,omitempty" yaml:"uePrivacyRequirements" bson:"uePrivacyRequirements,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ScheduledLocTime      *time.Time             `json:"scheduledLocTime,omitempty" yaml:"scheduledLocTime" bson:"scheduledLocTime,omitempty"`
	ReliableLocReq        bool                   `json:"reliableLocReq,omitempty" yaml:"reliableLocReq" bson:"reliableLocReq,omitempty"`
	IntegrityRequirements *IntegrityRequirements `json:"integrityRequirements,omitempty" yaml:"integrityRequirements" bson:"integrityRequirements,omitempty"`
}
