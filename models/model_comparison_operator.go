/*
 * Nudsf_Timer
 *
 * Nudsf Timer Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.598 UDSF Services, V17.6.0.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.598/
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ComparisonOperator string

// List of ComparisonOperator
const (
	ComparisonOperator_EQ  ComparisonOperator = "EQ"
	ComparisonOperator_NEQ ComparisonOperator = "NEQ"
	ComparisonOperator_GT  ComparisonOperator = "GT"
	ComparisonOperator_GTE ComparisonOperator = "GTE"
	ComparisonOperator_LT  ComparisonOperator = "LT"
	ComparisonOperator_LTE ComparisonOperator = "LTE"
)
