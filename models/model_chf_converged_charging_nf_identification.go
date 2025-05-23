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

type ChfConvergedChargingNfIdentification struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NFName string `json:"nFName,omitempty" yaml:"nFName" bson:"nFName,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	NFIPv4Address     string                                `json:"nFIPv4Address,omitempty" yaml:"nFIPv4Address" bson:"nFIPv4Address,omitempty"`
	NFIPv6Address     string                                `json:"nFIPv6Address,omitempty" yaml:"nFIPv6Address" bson:"nFIPv6Address,omitempty"`
	NFPLMNID          *PlmnId                               `json:"nFPLMNID,omitempty" yaml:"nFPLMNID" bson:"nFPLMNID,omitempty"`
	NodeFunctionality ChfConvergedChargingNodeFunctionality `json:"nodeFunctionality" yaml:"nodeFunctionality" bson:"nodeFunctionality,omitempty"`
	NFFqdn            string                                `json:"nFFqdn,omitempty" yaml:"nFFqdn" bson:"nFFqdn,omitempty"`
}
