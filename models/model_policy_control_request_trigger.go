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

type PolicyControlRequestTrigger string

// List of PolicyControlRequestTrigger
const (
	PolicyControlRequestTrigger_PLMN_CH                          PolicyControlRequestTrigger = "PLMN_CH"
	PolicyControlRequestTrigger_RES_MO_RE                        PolicyControlRequestTrigger = "RES_MO_RE"
	PolicyControlRequestTrigger_AC_TY_CH                         PolicyControlRequestTrigger = "AC_TY_CH"
	PolicyControlRequestTrigger_UE_IP_CH                         PolicyControlRequestTrigger = "UE_IP_CH"
	PolicyControlRequestTrigger_UE_MAC_CH                        PolicyControlRequestTrigger = "UE_MAC_CH"
	PolicyControlRequestTrigger_AN_CH_COR                        PolicyControlRequestTrigger = "AN_CH_COR"
	PolicyControlRequestTrigger_US_RE                            PolicyControlRequestTrigger = "US_RE"
	PolicyControlRequestTrigger_APP_STA                          PolicyControlRequestTrigger = "APP_STA"
	PolicyControlRequestTrigger_APP_STO                          PolicyControlRequestTrigger = "APP_STO"
	PolicyControlRequestTrigger_AN_INFO                          PolicyControlRequestTrigger = "AN_INFO"
	PolicyControlRequestTrigger_CM_SES_FAIL                      PolicyControlRequestTrigger = "CM_SES_FAIL"
	PolicyControlRequestTrigger_PS_DA_OFF                        PolicyControlRequestTrigger = "PS_DA_OFF"
	PolicyControlRequestTrigger_DEF_QOS_CH                       PolicyControlRequestTrigger = "DEF_QOS_CH"
	PolicyControlRequestTrigger_SE_AMBR_CH                       PolicyControlRequestTrigger = "SE_AMBR_CH"
	PolicyControlRequestTrigger_QOS_NOTIF                        PolicyControlRequestTrigger = "QOS_NOTIF"
	PolicyControlRequestTrigger_NO_CREDIT                        PolicyControlRequestTrigger = "NO_CREDIT"
	PolicyControlRequestTrigger_REALLO_OF_CREDIT                 PolicyControlRequestTrigger = "REALLO_OF_CREDIT"
	PolicyControlRequestTrigger_PRA_CH                           PolicyControlRequestTrigger = "PRA_CH"
	PolicyControlRequestTrigger_SAREA_CH                         PolicyControlRequestTrigger = "SAREA_CH"
	PolicyControlRequestTrigger_SCNN_CH                          PolicyControlRequestTrigger = "SCNN_CH"
	PolicyControlRequestTrigger_RE_TIMEOUT                       PolicyControlRequestTrigger = "RE_TIMEOUT"
	PolicyControlRequestTrigger_RES_RELEASE                      PolicyControlRequestTrigger = "RES_RELEASE"
	PolicyControlRequestTrigger_SUCC_RES_ALLO                    PolicyControlRequestTrigger = "SUCC_RES_ALLO"
	PolicyControlRequestTrigger_RAI_CH                           PolicyControlRequestTrigger = "RAI_CH"
	PolicyControlRequestTrigger_RAT_TY_CH                        PolicyControlRequestTrigger = "RAT_TY_CH"
	PolicyControlRequestTrigger_REF_QOS_IND_CH                   PolicyControlRequestTrigger = "REF_QOS_IND_CH"
	PolicyControlRequestTrigger_NUM_OF_PACKET_FILTER             PolicyControlRequestTrigger = "NUM_OF_PACKET_FILTER"
	PolicyControlRequestTrigger_UE_STATUS_RESUME                 PolicyControlRequestTrigger = "UE_STATUS_RESUME"
	PolicyControlRequestTrigger_UE_TZ_CH                         PolicyControlRequestTrigger = "UE_TZ_CH"
	PolicyControlRequestTrigger_AUTH_PROF_CH                     PolicyControlRequestTrigger = "AUTH_PROF_CH"
	PolicyControlRequestTrigger_QOS_MONITORING                   PolicyControlRequestTrigger = "QOS_MONITORING"
	PolicyControlRequestTrigger_SCELL_CH                         PolicyControlRequestTrigger = "SCELL_CH"
	PolicyControlRequestTrigger_USER_LOCATION_CH                 PolicyControlRequestTrigger = "USER_LOCATION_CH"
	PolicyControlRequestTrigger_EPS_FALLBACK                     PolicyControlRequestTrigger = "EPS_FALLBACK"
	PolicyControlRequestTrigger_MA_PDU                           PolicyControlRequestTrigger = "MA_PDU"
	PolicyControlRequestTrigger_TSN_BRIDGE_INFO                  PolicyControlRequestTrigger = "TSN_BRIDGE_INFO"
	PolicyControlRequestTrigger__5_G_RG_JOIN                     PolicyControlRequestTrigger = "5G_RG_JOIN"
	PolicyControlRequestTrigger__5_G_RG_LEAVE                    PolicyControlRequestTrigger = "5G_RG_LEAVE"
	PolicyControlRequestTrigger_DDN_FAILURE                      PolicyControlRequestTrigger = "DDN_FAILURE"
	PolicyControlRequestTrigger_DDN_DELIVERY_STATUS              PolicyControlRequestTrigger = "DDN_DELIVERY_STATUS"
	PolicyControlRequestTrigger_GROUP_ID_LIST_CHG                PolicyControlRequestTrigger = "GROUP_ID_LIST_CHG"
	PolicyControlRequestTrigger_DDN_FAILURE_CANCELLATION         PolicyControlRequestTrigger = "DDN_FAILURE_CANCELLATION"
	PolicyControlRequestTrigger_DDN_DELIVERY_STATUS_CANCELLATION PolicyControlRequestTrigger = "DDN_DELIVERY_STATUS_CANCELLATION"
	PolicyControlRequestTrigger_VPLMN_QOS_CH                     PolicyControlRequestTrigger = "VPLMN_QOS_CH"
	PolicyControlRequestTrigger_SUCC_QOS_UPDATE                  PolicyControlRequestTrigger = "SUCC_QOS_UPDATE"
	PolicyControlRequestTrigger_SAT_CATEGORY_CHG                 PolicyControlRequestTrigger = "SAT_CATEGORY_CHG"
	PolicyControlRequestTrigger_PCF_UE_NOTIF_IND                 PolicyControlRequestTrigger = "PCF_UE_NOTIF_IND"
	PolicyControlRequestTrigger_NWDAF_DATA_CHG                   PolicyControlRequestTrigger = "NWDAF_DATA_CHG"
)
