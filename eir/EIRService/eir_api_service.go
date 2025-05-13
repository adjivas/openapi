package EIRSelection

import (
	"context"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/free5gc/amf/internal/logger"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
)

// Linger please
var (
	_ context.Context
)

type EIRApiService service

type EIREquipementStatusGetResponse struct {
	Status string `json:"status"`
}

func (a *EIRApiService) EIREquipementStatusGet(ctx context.Context, imei string) (*EIREquipementStatusGetResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EIREquipementStatusGetResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/equipement-status?pei=" + imei

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	var requestProblemDetails models.ProblemDetails

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 400:
		err = openapi.Deserialize(&requestProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = requestProblemDetails
		return nil, apiError
	case 401:
		err = openapi.Deserialize(&requestProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = requestProblemDetails
		return nil, apiError
	case 404:
		err = openapi.Deserialize(&requestProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = requestProblemDetails
		return nil, apiError
	case 414:
		err = openapi.Deserialize(&requestProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = requestProblemDetails
		return nil, apiError
	case 429:
		err = openapi.Deserialize(&requestProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = requestProblemDetails
		return nil, apiError
	case 500:
		err = openapi.Deserialize(&requestProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = requestProblemDetails
		return nil, apiError
	case 503:
		err = openapi.Deserialize(&requestProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = requestProblemDetails
		return nil, apiError
	default:
		return nil, apiError
	}
}
