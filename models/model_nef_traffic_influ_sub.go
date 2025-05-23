/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.522 V17.7.0; 5G System; Network Exposure Function Northbound APIs.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.522/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents a traffic influence subscription.
type NefTrafficInfluSub struct {
	// Identifies a service on behalf of which the AF is issuing the request.
	AfServiceId string `json:"afServiceId,omitempty" yaml:"afServiceId" bson:"afServiceId,omitempty"`
	// Identifies an application.
	AfAppId string `json:"afAppId,omitempty" yaml:"afAppId" bson:"afAppId,omitempty"`
	// Identifies an NEF Northbound interface transaction, generated by the AF.
	AfTransId string `json:"afTransId,omitempty" yaml:"afTransId" bson:"afTransId,omitempty"`
	// Identifies whether an application can be relocated once a location of the application has been selected.
	AppReloInd bool `json:"appReloInd,omitempty" yaml:"appReloInd" bson:"appReloInd,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn    string  `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExternalGroupId string `json:"externalGroupId,omitempty" yaml:"externalGroupId" bson:"externalGroupId,omitempty"`
	// Identifies whether the AF request applies to any UE. This attribute shall set to \"true\" if applicable for any UE, otherwise, set to \"false\".
	AnyUeInd bool `json:"anyUeInd,omitempty" yaml:"anyUeInd" bson:"anyUeInd,omitempty"`
	// Identifies the requirement to be notified of the event(s).
	SubscribedEvents []SubscribedEvent `json:"subscribedEvents,omitempty" yaml:"subscribedEvents" bson:"subscribedEvents,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi string `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	Ipv4Addr string `json:"ipv4Addr,omitempty" yaml:"ipv4Addr" bson:"ipv4Addr,omitempty"`
	IpDomain string `json:"ipDomain,omitempty" yaml:"ipDomain" bson:"ipDomain,omitempty"`
	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	Ipv6Addr string `json:"ipv6Addr,omitempty" yaml:"ipv6Addr" bson:"ipv6Addr,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	MacAddr     string         `json:"macAddr,omitempty" yaml:"macAddr" bson:"macAddr,omitempty"`
	DnaiChgType DnaiChangeType `json:"dnaiChgType,omitempty" yaml:"dnaiChgType" bson:"dnaiChgType,omitempty"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination,omitempty" yaml:"notificationDestination" bson:"notificationDestination,omitempty"`
	// Set to true by the SCS/AS to request the NEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise.
	RequestTestNotification bool                `json:"requestTestNotification,omitempty" yaml:"requestTestNotification" bson:"requestTestNotification,omitempty"`
	WebsockNotifConfig      *WebsockNotifConfig `json:"websockNotifConfig,omitempty" yaml:"websockNotifConfig" bson:"websockNotifConfig,omitempty"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self string `json:"self,omitempty" yaml:"self" bson:"self,omitempty"`
	// Identifies IP packet filters.
	TrafficFilters []FlowInfo `json:"trafficFilters,omitempty" yaml:"trafficFilters" bson:"trafficFilters,omitempty"`
	// Identifies Ethernet packet filters.
	EthTrafficFilters []EthFlowDescription `json:"ethTrafficFilters,omitempty" yaml:"ethTrafficFilters" bson:"ethTrafficFilters,omitempty"`
	// Identifies the N6 traffic routing requirement.
	TrafficRoutes  []*RouteToLocation `json:"trafficRoutes,omitempty" yaml:"trafficRoutes" bson:"trafficRoutes,omitempty"`
	TfcCorrInd     bool               `json:"tfcCorrInd,omitempty" yaml:"tfcCorrInd" bson:"tfcCorrInd,omitempty"`
	TempValidities []TemporalValidity `json:"tempValidities,omitempty" yaml:"tempValidities" bson:"tempValidities,omitempty"`
	// Identifies a geographic zone that the AF request applies only to the traffic of UE(s) located in this specific zone.
	ValidGeoZoneIds []string `json:"validGeoZoneIds,omitempty" yaml:"validGeoZoneIds" bson:"validGeoZoneIds,omitempty"`
	// Identifies geographical areas within which the AF request applies.
	GeoAreas      []GeographicalArea `json:"geoAreas,omitempty" yaml:"geoAreas" bson:"geoAreas,omitempty"`
	AfAckInd      bool               `json:"afAckInd,omitempty" yaml:"afAckInd" bson:"afAckInd,omitempty"`
	AddrPreserInd bool               `json:"addrPreserInd,omitempty" yaml:"addrPreserInd" bson:"addrPreserInd,omitempty"`
	// Indicates whether simultaneous connectivity should be temporarily maintained for the source and target PSA.
	SimConnInd bool `json:"simConnInd,omitempty" yaml:"simConnInd" bson:"simConnInd,omitempty"`
	// indicating a time in seconds.
	SimConnTerm int32 `json:"simConnTerm,omitempty" yaml:"simConnTerm" bson:"simConnTerm,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxAllowedUpLat int32 `json:"maxAllowedUpLat,omitempty" yaml:"maxAllowedUpLat" bson:"maxAllowedUpLat,omitempty"`
	// Contains EAS IP replacement information.
	EasIpReplaceInfos []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty" yaml:"easIpReplaceInfos" bson:"easIpReplaceInfos,omitempty"`
	// Indicates the EAS rediscovery is required for the application if it is included and set to \"true\".
	EasRedisInd  bool                  `json:"easRedisInd,omitempty" yaml:"easRedisInd" bson:"easRedisInd,omitempty"`
	EventReq     *ReportingInformation `json:"eventReq,omitempty" yaml:"eventReq" bson:"eventReq,omitempty"`
	EventReports []EventNotification   `json:"eventReports,omitempty" yaml:"eventReports" bson:"eventReports,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat string `json:"suppFeat,omitempty" yaml:"suppFeat" bson:"suppFeat,omitempty"`
}
