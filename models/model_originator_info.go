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

type OriginatorInfo struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	OriginatorSUPI string `json:"originatorSUPI,omitempty" yaml:"originatorSUPI" bson:"originatorSUPI,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	OriginatorGPSI            string         `json:"originatorGPSI,omitempty" yaml:"originatorGPSI" bson:"originatorGPSI,omitempty"`
	OriginatorOtherAddress    *SmAddressInfo `json:"originatorOtherAddress,omitempty" yaml:"originatorOtherAddress" bson:"originatorOtherAddress,omitempty"`
	OriginatorReceivedAddress *SmAddressInfo `json:"originatorReceivedAddress,omitempty" yaml:"originatorReceivedAddress" bson:"originatorReceivedAddress,omitempty"`
	OriginatorSCCPAddress     string         `json:"originatorSCCPAddress,omitempty" yaml:"originatorSCCPAddress" bson:"originatorSCCPAddress,omitempty"`
	SMOriginatorInterface     *SmInterface   `json:"sMOriginatorInterface,omitempty" yaml:"sMOriginatorInterface" bson:"sMOriginatorInterface,omitempty"`
	SMOriginatorProtocolId    string         `json:"sMOriginatorProtocolId,omitempty" yaml:"sMOriginatorProtocolId" bson:"sMOriginatorProtocolId,omitempty"`
}
