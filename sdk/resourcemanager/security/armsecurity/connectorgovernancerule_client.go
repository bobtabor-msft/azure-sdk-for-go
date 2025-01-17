//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ConnectorGovernanceRuleClient contains the methods for the SecurityConnectorGovernanceRule group.
// Don't use this type directly, use NewConnectorGovernanceRuleClient() instead.
type ConnectorGovernanceRuleClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConnectorGovernanceRuleClient creates a new instance of ConnectorGovernanceRuleClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectorGovernanceRuleClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectorGovernanceRuleClient, error) {
	cl, err := arm.NewClient(moduleName+".ConnectorGovernanceRuleClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectorGovernanceRuleClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get a list of all relevant governanceRules over a security connector level scope
//
// Generated from API version 2022-01-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - options - ConnectorGovernanceRuleClientListOptions contains the optional parameters for the ConnectorGovernanceRuleClient.NewListPager
//     method.
func (client *ConnectorGovernanceRuleClient) NewListPager(resourceGroupName string, securityConnectorName string, options *ConnectorGovernanceRuleClientListOptions) *runtime.Pager[ConnectorGovernanceRuleClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectorGovernanceRuleClientListResponse]{
		More: func(page ConnectorGovernanceRuleClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectorGovernanceRuleClientListResponse) (ConnectorGovernanceRuleClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, securityConnectorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConnectorGovernanceRuleClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConnectorGovernanceRuleClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConnectorGovernanceRuleClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ConnectorGovernanceRuleClient) listCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, options *ConnectorGovernanceRuleClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/governanceRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConnectorGovernanceRuleClient) listHandleResponse(resp *http.Response) (ConnectorGovernanceRuleClientListResponse, error) {
	result := ConnectorGovernanceRuleClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GovernanceRuleList); err != nil {
		return ConnectorGovernanceRuleClientListResponse{}, err
	}
	return result, nil
}
