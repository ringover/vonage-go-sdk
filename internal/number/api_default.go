/*
 * Numbers API
 *
 * The Numbers API enables you to manage your existing numbers and buy new virtual numbers for use with Nexmo's APIs. Further information is here: <https://developer.nexmo.com/numbers/overview>
 *
 * API version: 1.0.18
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package number

import (
	_context "context"
	"io"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

// BuyANumberOpts Optional parameters for the method 'BuyANumber'
type BuyANumberOpts struct {
	TargetApiKey optional.String
}

/*
BuyANumber Buy a number
Request to purchase a specific inbound number.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param country The two character country code in ISO 3166-1 alpha-2 format
  - @param msisdn An available inbound virtual number.
  - @param optional nil or *BuyANumberOpts - Optional Parameters:
  - @param "TargetApiKey" (optional.String) -  If you’d like to perform an action on a subaccount, provide the `api_key` of that account here. If you’d like to perform an action on your own account, you do not need to provide this field.

@return Response
*/
func (a *DefaultApiService) BuyANumber(ctx _context.Context, country string, msisdn string, localVarOptionals *BuyANumberOpts) (Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/number/buy"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(country) < 2 {
		return localVarReturnValue, nil, reportError("country must have at least 2 elements")
	}
	if strlen(country) > 2 {
		return localVarReturnValue, nil, reportError("country must have less than 2 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("country", parameterToString(country, ""))
	localVarFormParams.Add("msisdn", parameterToString(msisdn, ""))
	if localVarOptionals != nil && localVarOptionals.TargetApiKey.IsSet() {
		localVarFormParams.Add("target_api_key", parameterToString(localVarOptionals.TargetApiKey.Value(), ""))
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("api_key", key)
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("api_secret", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// CancelANumberOpts Optional parameters for the method 'CancelANumber'
type CancelANumberOpts struct {
	TargetApiKey optional.String
}

/*
CancelANumber Cancel a number
Cancel your subscription for a specific inbound number.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param country The two character country code in ISO 3166-1 alpha-2 format
  - @param msisdn An available inbound virtual number.
  - @param optional nil or *CancelANumberOpts - Optional Parameters:
  - @param "TargetApiKey" (optional.String) -  If you’d like to perform an action on a subaccount, provide the `api_key` of that account here. If you’d like to perform an action on your own account, you do not need to provide this field.

@return Response
*/
func (a *DefaultApiService) CancelANumber(ctx _context.Context, country string, msisdn string, localVarOptionals *CancelANumberOpts) (Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/number/cancel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(country) < 2 {
		return localVarReturnValue, nil, reportError("country must have at least 2 elements")
	}
	if strlen(country) > 2 {
		return localVarReturnValue, nil, reportError("country must have less than 2 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("country", parameterToString(country, ""))
	localVarFormParams.Add("msisdn", parameterToString(msisdn, ""))
	if localVarOptionals != nil && localVarOptionals.TargetApiKey.IsSet() {
		localVarFormParams.Add("target_api_key", parameterToString(localVarOptionals.TargetApiKey.Value(), ""))
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("api_key", key)
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("api_secret", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetAvailableNumbersOpts Optional parameters for the method 'GetAvailableNumbers'
type GetAvailableNumbersOpts struct {
	Type_         optional.String
	Pattern       optional.String
	SearchPattern optional.Int32
	Features      optional.String
	Size          optional.Int32
	Index         optional.Int32
}

/*
GetAvailableNumbers Search available numbers
Retrieve inbound numbers that are available for the specified country.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param country The two character country code to filter on (in ISO 3166-1 alpha-2 format)
  - @param optional nil or *GetAvailableNumbersOpts - Optional Parameters:
  - @param "Type_" (optional.String) -  Set this parameter to filter the type of number, such as mobile or landline
  - @param "Pattern" (optional.String) -  The number pattern you want to search for. Use in conjunction with `search_pattern`.
  - @param "SearchPattern" (optional.Int32) -  The strategy you want to use for matching:   * `0` - Search for numbers that start with `pattern` (Note: all numbers are in E.164 format, so the starting pattern includes the country code, such as 1 for USA) * `1` - Search for numbers that contain `pattern` * `2` - Search for numbers that end with `pattern`
  - @param "Features" (optional.String) -  Available features are `SMS`, `VOICE` and `MMS`. To look for numbers that support multiple features, use a comma-separated value: `SMS,MMS,VOICE`.
  - @param "Size" (optional.Int32) -  Page size
  - @param "Index" (optional.Int32) -  Page index

@return AvailableNumbers
*/
func (a *DefaultApiService) GetAvailableNumbers(ctx _context.Context, country string, localVarOptionals *GetAvailableNumbersOpts) (AvailableNumbers, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AvailableNumbers
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/number/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(country) < 2 {
		return localVarReturnValue, nil, reportError("country must have at least 2 elements")
	}
	if strlen(country) > 2 {
		return localVarReturnValue, nil, reportError("country must have less than 2 elements")
	}

	localVarQueryParams.Add("country", parameterToString(country, ""))
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Pattern.IsSet() {
		localVarQueryParams.Add("pattern", parameterToString(localVarOptionals.Pattern.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchPattern.IsSet() {
		localVarQueryParams.Add("search_pattern", parameterToString(localVarOptionals.SearchPattern.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Features.IsSet() {
		localVarQueryParams.Add("features", parameterToString(localVarOptionals.Features.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Index.IsSet() {
		localVarQueryParams.Add("index", parameterToString(localVarOptionals.Index.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("api_key", key)
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("api_secret", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v AvailableNumbers
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetOwnedNumbersOpts Optional parameters for the method 'GetOwnedNumbers'
type GetOwnedNumbersOpts struct {
	ApplicationId  optional.String
	HasApplication optional.Bool
	Country        optional.String
	Pattern        optional.String
	SearchPattern  optional.Int32
	Size           optional.Int32
	Index          optional.Int32
}

/*
GetOwnedNumbers List the numbers you own
Retrieve all the inbound numbers associated with your Nexmo account.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param optional nil or *GetOwnedNumbersOpts - Optional Parameters:
  - @param "ApplicationId" (optional.String) -  The Application that you want to return the numbers for.
  - @param "HasApplication" (optional.Bool) -  Set this optional field to `true` to restrict your results to numbers associated with an Application (any Application). Set to `false` to find all numbers not associated with any Application. Omit the field to avoid filtering on whether or not the number is assigned to an Application.
  - @param "Country" (optional.String) -
  - @param "Pattern" (optional.String) -  The number pattern you want to search for. Use in conjunction with `search_pattern`.
  - @param "SearchPattern" (optional.Int32) -  The strategy you want to use for matching:   * `0` - Search for numbers that start with `pattern` (Note: all numbers are in E.164 format, so the starting pattern includes the country code, such as 1 for USA) * `1` - Search for numbers that contain `pattern` * `2` - Search for numbers that end with `pattern`
  - @param "Size" (optional.Int32) -  Page size
  - @param "Index" (optional.Int32) -  Page index

@return InboundNumbers
*/
func (a *DefaultApiService) GetOwnedNumbers(ctx _context.Context, localVarOptionals *GetOwnedNumbersOpts) (InboundNumbers, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InboundNumbers
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account/numbers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.ApplicationId.IsSet() {
		localVarQueryParams.Add("application_id", parameterToString(localVarOptionals.ApplicationId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HasApplication.IsSet() {
		localVarQueryParams.Add("has_application", parameterToString(localVarOptionals.HasApplication.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarQueryParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Pattern.IsSet() {
		localVarQueryParams.Add("pattern", parameterToString(localVarOptionals.Pattern.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchPattern.IsSet() {
		localVarQueryParams.Add("search_pattern", parameterToString(localVarOptionals.SearchPattern.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Index.IsSet() {
		localVarQueryParams.Add("index", parameterToString(localVarOptionals.Index.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("api_key", key)
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("api_secret", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v InboundNumbers
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// UpdateANumberOpts Optional parameters for the method 'UpdateANumber'
type UpdateANumberOpts struct {
	AppId                 optional.String
	MoHttpUrl             optional.String
	MoSmppSysType         optional.String
	VoiceCallbackType     optional.String
	VoiceCallbackValue    optional.String
	VoiceStatusCallback   optional.String
	MessagesCallbackType  optional.String
	MessagesCallbackValue optional.String
}

/*
UpdateANumber Update a number
Change the behaviour of a number that you own.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param country The two character country code in ISO 3166-1 alpha-2 format
  - @param msisdn An available inbound virtual number.
  - @param optional nil or *UpdateANumberOpts - Optional Parameters:
  - @param "AppId" (optional.String) -  The Application that will handle inbound traffic to this number.
  - @param "MoHttpUrl" (optional.String) -  An URL-encoded URI to the webhook endpoint that handles inbound messages. Your webhook endpoint must be active before you make this request. Nexmo makes a `GET` request to the endpoint and checks that it returns a `200 OK` response. Set this parameter's value to an empty string to remove the webhook.
  - @param "MoSmppSysType" (optional.String) -  The associated system type for your SMPP client
  - @param "VoiceCallbackType" (optional.String) -  Specify whether inbound voice calls on your number are forwarded to a SIP or a telephone number.  This must be used with the `voiceCallbackValue` parameter. If set, `sip` or `tel` are prioritized over the Voice capability in your Application.  *Note: The `app` value is deprecated and will be removed in future.*
  - @param "VoiceCallbackValue" (optional.String) -  A SIP URI or telephone number. Must be used with the `voiceCallbackType` parameter.
  - @param "VoiceStatusCallback" (optional.String) -  A webhook URI for Nexmo to send a request to when a call ends
  - @param "MessagesCallbackType" (optional.String) -  <strong>DEPRECATED</strong> - We recommend that you use `app_id` instead.  Specifies the Messages webhook type (always `app`) associated with this number and must be used with the `messagesCallbackValue` parameter.
  - @param "MessagesCallbackValue" (optional.String) -  <strong>DEPRECATED</strong> - We recommend that you use `app_id` instead.  Specifies the Application ID of your Messages application. It must be used with the `messagesCallbackType` parameter.

@return Response
*/
func (a *DefaultApiService) UpdateANumber(ctx _context.Context, country string, msisdn string, localVarOptionals *UpdateANumberOpts) (Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/number/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(country) < 2 {
		return localVarReturnValue, nil, reportError("country must have at least 2 elements")
	}
	if strlen(country) > 2 {
		return localVarReturnValue, nil, reportError("country must have less than 2 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("country", parameterToString(country, ""))
	localVarFormParams.Add("msisdn", parameterToString(msisdn, ""))
	if localVarOptionals != nil && localVarOptionals.AppId.IsSet() {
		localVarFormParams.Add("app_id", parameterToString(localVarOptionals.AppId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MoHttpUrl.IsSet() {
		localVarFormParams.Add("moHttpUrl", parameterToString(localVarOptionals.MoHttpUrl.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MoSmppSysType.IsSet() {
		localVarFormParams.Add("moSmppSysType", parameterToString(localVarOptionals.MoSmppSysType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.VoiceCallbackType.IsSet() {
		localVarFormParams.Add("voiceCallbackType", parameterToString(localVarOptionals.VoiceCallbackType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.VoiceCallbackValue.IsSet() {
		localVarFormParams.Add("voiceCallbackValue", parameterToString(localVarOptionals.VoiceCallbackValue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.VoiceStatusCallback.IsSet() {
		localVarFormParams.Add("voiceStatusCallback", parameterToString(localVarOptionals.VoiceStatusCallback.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MessagesCallbackType.IsSet() {
		localVarFormParams.Add("messagesCallbackType", parameterToString(localVarOptionals.MessagesCallbackType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MessagesCallbackValue.IsSet() {
		localVarFormParams.Add("messagesCallbackValue", parameterToString(localVarOptionals.MessagesCallbackValue.Value(), ""))
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("api_key", key)
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("api_secret", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
