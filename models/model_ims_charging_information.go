/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291 V17.9.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 3.1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type ImsChargingInformation struct {
	EventType            *SipEventType                        `json:"eventType,omitempty" yaml:"eventType" bson:"eventType,omitempty"`
	IMSNodeFunctionality ImsNodeFunctionality                 `json:"iMSNodeFunctionality,omitempty" yaml:"iMSNodeFunctionality" bson:"iMSNodeFunctionality,omitempty"`
	RoleOfNode           RoleOfImsNode                        `json:"roleOfNode,omitempty" yaml:"roleOfNode" bson:"roleOfNode,omitempty"`
	UserInformation      *ChfConvergedChargingUserInformation `json:"userInformation,omitempty" yaml:"userInformation" bson:"userInformation,omitempty"`
	UserLocationInfo     *UserLocation                        `json:"userLocationInfo,omitempty" yaml:"userLocationInfo" bson:"userLocationInfo,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	UeTimeZone                          string                   `json:"ueTimeZone,omitempty" yaml:"ueTimeZone" bson:"ueTimeZone,omitempty"`
	Var3gppPSDataOffStatus              Model3GpppsDataOffStatus `json:"3gppPSDataOffStatus,omitempty" yaml:"3gppPSDataOffStatus" bson:"3gppPSDataOffStatus,omitempty"`
	IsupCause                           *IsupCause               `json:"isupCause,omitempty" yaml:"isupCause" bson:"isupCause,omitempty"`
	ControlPlaneAddress                 *ImsAddress              `json:"controlPlaneAddress,omitempty" yaml:"controlPlaneAddress" bson:"controlPlaneAddress,omitempty"`
	VlrNumber                           string                   `json:"vlrNumber,omitempty" yaml:"vlrNumber" bson:"vlrNumber,omitempty"`
	MscAddress                          string                   `json:"mscAddress,omitempty" yaml:"mscAddress" bson:"mscAddress,omitempty"`
	UserSessionID                       string                   `json:"userSessionID,omitempty" yaml:"userSessionID" bson:"userSessionID,omitempty"`
	OutgoingSessionID                   string                   `json:"outgoingSessionID,omitempty" yaml:"outgoingSessionID" bson:"outgoingSessionID,omitempty"`
	SessionPriority                     ImsSessionPriority       `json:"sessionPriority,omitempty" yaml:"sessionPriority" bson:"sessionPriority,omitempty"`
	CallingPartyAddresses               []string                 `json:"callingPartyAddresses,omitempty" yaml:"callingPartyAddresses" bson:"callingPartyAddresses,omitempty"`
	CalledPartyAddress                  string                   `json:"calledPartyAddress,omitempty" yaml:"calledPartyAddress" bson:"calledPartyAddress,omitempty"`
	NumberPortabilityRoutinginformation string                   `json:"numberPortabilityRoutinginformation,omitempty" yaml:"numberPortabilityRoutinginformation" bson:"numberPortabilityRoutinginformation,omitempty"`
	CarrierSelectRoutingInformation     string                   `json:"carrierSelectRoutingInformation,omitempty" yaml:"carrierSelectRoutingInformation" bson:"carrierSelectRoutingInformation,omitempty"`
	AlternateChargedPartyAddress        string                   `json:"alternateChargedPartyAddress,omitempty" yaml:"alternateChargedPartyAddress" bson:"alternateChargedPartyAddress,omitempty"`
	RequestedPartyAddress               []string                 `json:"requestedPartyAddress,omitempty" yaml:"requestedPartyAddress" bson:"requestedPartyAddress,omitempty"`
	CalledAssertedIdentities            []string                 `json:"calledAssertedIdentities,omitempty" yaml:"calledAssertedIdentities" bson:"calledAssertedIdentities,omitempty"`
	CalledIdentityChanges               []CalledIdentityChange   `json:"calledIdentityChanges,omitempty" yaml:"calledIdentityChanges" bson:"calledIdentityChanges,omitempty"`
	AssociatedURI                       []string                 `json:"associatedURI,omitempty" yaml:"associatedURI" bson:"associatedURI,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamps                         *time.Time                  `json:"timeStamps,omitempty" yaml:"timeStamps" bson:"timeStamps,omitempty"`
	ApplicationServerInformation       []string                    `json:"applicationServerInformation,omitempty" yaml:"applicationServerInformation" bson:"applicationServerInformation,omitempty"`
	InterOperatorIdentifier            []InterOperatorIdentifier   `json:"interOperatorIdentifier,omitempty" yaml:"interOperatorIdentifier" bson:"interOperatorIdentifier,omitempty"`
	ImsChargingIdentifier              string                      `json:"imsChargingIdentifier,omitempty" yaml:"imsChargingIdentifier" bson:"imsChargingIdentifier,omitempty"`
	RelatedICID                        string                      `json:"relatedICID,omitempty" yaml:"relatedICID" bson:"relatedICID,omitempty"`
	RelatedICIDGenerationNode          string                      `json:"relatedICIDGenerationNode,omitempty" yaml:"relatedICIDGenerationNode" bson:"relatedICIDGenerationNode,omitempty"`
	TransitIOIList                     []string                    `json:"transitIOIList,omitempty" yaml:"transitIOIList" bson:"transitIOIList,omitempty"`
	EarlyMediaDescription              []EarlyMediaDescription     `json:"earlyMediaDescription,omitempty" yaml:"earlyMediaDescription" bson:"earlyMediaDescription,omitempty"`
	SdpSessionDescription              []string                    `json:"sdpSessionDescription,omitempty" yaml:"sdpSessionDescription" bson:"sdpSessionDescription,omitempty"`
	SdpMediaComponent                  []SdpMediaComponent         `json:"sdpMediaComponent,omitempty" yaml:"sdpMediaComponent" bson:"sdpMediaComponent,omitempty"`
	ServedPartyIPAddress               *ImsAddress                 `json:"servedPartyIPAddress,omitempty" yaml:"servedPartyIPAddress" bson:"servedPartyIPAddress,omitempty"`
	ServerCapabilities                 *ServerCapabilities         `json:"serverCapabilities,omitempty" yaml:"serverCapabilities" bson:"serverCapabilities,omitempty"`
	TrunkGroupID                       *TrunkGroupId               `json:"trunkGroupID,omitempty" yaml:"trunkGroupID" bson:"trunkGroupID,omitempty"`
	BearerService                      string                      `json:"bearerService,omitempty" yaml:"bearerService" bson:"bearerService,omitempty"`
	ImsServiceId                       string                      `json:"imsServiceId,omitempty" yaml:"imsServiceId" bson:"imsServiceId,omitempty"`
	MessageBodies                      []MessageBody               `json:"messageBodies,omitempty" yaml:"messageBodies" bson:"messageBodies,omitempty"`
	AccessNetworkInformation           []string                    `json:"accessNetworkInformation,omitempty" yaml:"accessNetworkInformation" bson:"accessNetworkInformation,omitempty"`
	AdditionalAccessNetworkInformation string                      `json:"additionalAccessNetworkInformation,omitempty" yaml:"additionalAccessNetworkInformation" bson:"additionalAccessNetworkInformation,omitempty"`
	CellularNetworkInformation         string                      `json:"cellularNetworkInformation,omitempty" yaml:"cellularNetworkInformation" bson:"cellularNetworkInformation,omitempty"`
	AccessTransferInformation          []AccessTransferInformation `json:"accessTransferInformation,omitempty" yaml:"accessTransferInformation" bson:"accessTransferInformation,omitempty"`
	AccessNetworkInfoChange            []AccessNetworkInfoChange   `json:"accessNetworkInfoChange,omitempty" yaml:"accessNetworkInfoChange" bson:"accessNetworkInfoChange,omitempty"`
	ImsCommunicationServiceID          string                      `json:"imsCommunicationServiceID,omitempty" yaml:"imsCommunicationServiceID" bson:"imsCommunicationServiceID,omitempty"`
	ImsApplicationReferenceID          string                      `json:"imsApplicationReferenceID,omitempty" yaml:"imsApplicationReferenceID" bson:"imsApplicationReferenceID,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	CauseCode                    int32            `json:"causeCode,omitempty" yaml:"causeCode" bson:"causeCode,omitempty"`
	ReasonHeader                 []string         `json:"reasonHeader,omitempty" yaml:"reasonHeader" bson:"reasonHeader,omitempty"`
	InitialIMSChargingIdentifier string           `json:"initialIMSChargingIdentifier,omitempty" yaml:"initialIMSChargingIdentifier" bson:"initialIMSChargingIdentifier,omitempty"`
	NniInformation               []NniInformation `json:"nniInformation,omitempty" yaml:"nniInformation" bson:"nniInformation,omitempty"`
	FromAddress                  string           `json:"fromAddress,omitempty" yaml:"fromAddress" bson:"fromAddress,omitempty"`
	ImsEmergencyIndication       bool             `json:"imsEmergencyIndication,omitempty" yaml:"imsEmergencyIndication" bson:"imsEmergencyIndication,omitempty"`
	ImsVisitedNetworkIdentifier  string           `json:"imsVisitedNetworkIdentifier,omitempty" yaml:"imsVisitedNetworkIdentifier" bson:"imsVisitedNetworkIdentifier,omitempty"`
	SipRouteHeaderReceived       string           `json:"sipRouteHeaderReceived,omitempty" yaml:"sipRouteHeaderReceived" bson:"sipRouteHeaderReceived,omitempty"`
	SipRouteHeaderTransmitted    string           `json:"sipRouteHeaderTransmitted,omitempty" yaml:"sipRouteHeaderTransmitted" bson:"sipRouteHeaderTransmitted,omitempty"`
	TadIdentifier                TadIdentifier    `json:"tadIdentifier,omitempty" yaml:"tadIdentifier" bson:"tadIdentifier,omitempty"`
	FeIdentifierList             string           `json:"feIdentifierList,omitempty" yaml:"feIdentifierList" bson:"feIdentifierList,omitempty"`
}
