/*
 * Nhss_imsUEAU
 *
 * Nhss UE Authentication Service for IMS. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * Source file: 3GPP TS 29.562 HSS Services for Interworking with IMS, version 16.5.0
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.562/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models



type DigestAuthentication struct {
	DigestRealm string `json:"digestRealm" yaml:"digestRealm" bson:"digestRealm"`
	DigestAlgorithm SipDigestAlgorithm `json:"digestAlgorithm" yaml:"digestAlgorithm" bson:"digestAlgorithm"`
	DigestQop SipDigestQop `json:"digestQop" yaml:"digestQop" bson:"digestQop"`
	Ha1 string `json:"ha1" yaml:"ha1" bson:"ha1"`
}
