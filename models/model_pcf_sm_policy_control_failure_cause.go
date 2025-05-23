/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.512 V17.11.0; 5G System; Session Management Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PcfSmPolicyControlFailureCause string

// List of PcfSMPolicyControlFailureCause
const (
	PcfSmPolicyControlFailureCause_PCC_RULE_EVENT       PcfSmPolicyControlFailureCause = "PCC_RULE_EVENT"
	PcfSmPolicyControlFailureCause_PCC_QOS_FLOW_EVENT   PcfSmPolicyControlFailureCause = "PCC_QOS_FLOW_EVENT"
	PcfSmPolicyControlFailureCause_RULE_PERMANENT_ERROR PcfSmPolicyControlFailureCause = "RULE_PERMANENT_ERROR"
	PcfSmPolicyControlFailureCause_RULE_TEMPORARY_ERROR PcfSmPolicyControlFailureCause = "RULE_TEMPORARY_ERROR"
	PcfSmPolicyControlFailureCause_POL_DEC_ERROR        PcfSmPolicyControlFailureCause = "POL_DEC_ERROR"
)
