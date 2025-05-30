/*
 * Nnef_EventExposure
 *
 * NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.591 V17.7.0; 5G System; Network Exposure Function Southbound Services; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.591/
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type MediaStreamingAccessRecord struct {
	// string with format 'date-time' as defined in OpenAPI.
	Timestamp                         *time.Time             `json:"timestamp" yaml:"timestamp" bson:"timestamp,omitempty"`
	MediaStreamHandlerEndpointAddress *EndpointAddress       `json:"mediaStreamHandlerEndpointAddress" yaml:"mediaStreamHandlerEndpointAddress" bson:"mediaStreamHandlerEndpointAddress,omitempty"`
	ApplicationServerEndpointAddress  *EndpointAddress       `json:"applicationServerEndpointAddress" yaml:"applicationServerEndpointAddress" bson:"applicationServerEndpointAddress,omitempty"`
	SessionIdentifier                 string                 `json:"sessionIdentifier,omitempty" yaml:"sessionIdentifier" bson:"sessionIdentifier,omitempty"`
	RequestMessage                    map[string]interface{} `json:"requestMessage" yaml:"requestMessage" bson:"requestMessage,omitempty"`
	CacheStatus                       CacheStatus            `json:"cacheStatus,omitempty" yaml:"cacheStatus" bson:"cacheStatus,omitempty"`
	ResponseMessage                   map[string]interface{} `json:"responseMessage" yaml:"responseMessage" bson:"responseMessage,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	ProcessingLatency float32                `json:"processingLatency" yaml:"processingLatency" bson:"processingLatency,omitempty"`
	ConnectionMetrics map[string]interface{} `json:"connectionMetrics,omitempty" yaml:"connectionMetrics" bson:"connectionMetrics,omitempty"`
}
