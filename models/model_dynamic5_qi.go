/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V17.11.0; 5G System; Session Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// It indicates the QoS Characteristics for a Non-standardised or not pre-configured 5QI for downlink and uplink.
type Dynamic5Qi struct {
	ResourceType QosResourceType `json:"resourceType" yaml:"resourceType" bson:"resourceType,omitempty"`
	// Unsigned integer indicating the 5QI Priority Level (see clauses 5.7.3.3 and 5.7.4 of 3GPP TS 23.501, within the range 1 to 127.Values are ordered in decreasing order of priority,  i.e. with 1 as the highest priority and 127 as the lowest priority.
	PriorityLevel int32 `json:"priorityLevel" yaml:"priorityLevel" bson:"priorityLevel,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	PacketDelayBudget int32 `json:"packetDelayBudget" yaml:"packetDelayBudget" bson:"packetDelayBudget,omitempty"`
	// String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \"scalar x 10-k\" where the scalar and the exponent k are each encoded as one decimal digit.
	PacketErrRate string `json:"packetErrRate" yaml:"packetErrRate" bson:"packetErrRate,omitempty"`
	// Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	AverWindow int32 `json:"averWindow,omitempty" yaml:"averWindow" bson:"averWindow,omitempty"`
	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.
	MaxDataBurstVol int32 `json:"maxDataBurstVol,omitempty" yaml:"maxDataBurstVol" bson:"maxDataBurstVol,omitempty"`
	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.
	ExtMaxDataBurstVol int32 `json:"extMaxDataBurstVol,omitempty" yaml:"extMaxDataBurstVol" bson:"extMaxDataBurstVol,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501 [8])), expressed in 0.01 milliseconds.
	ExtPacketDelBudget int32 `json:"extPacketDelBudget,omitempty" yaml:"extPacketDelBudget" bson:"extPacketDelBudget,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501 [8])), expressed in 0.01 milliseconds.
	CnPacketDelayBudgetDl int32 `json:"cnPacketDelayBudgetDl,omitempty" yaml:"cnPacketDelayBudgetDl" bson:"cnPacketDelayBudgetDl,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501 [8])), expressed in 0.01 milliseconds.
	CnPacketDelayBudgetUl int32 `json:"cnPacketDelayBudgetUl,omitempty" yaml:"cnPacketDelayBudgetUl" bson:"cnPacketDelayBudgetUl,omitempty"`
}
