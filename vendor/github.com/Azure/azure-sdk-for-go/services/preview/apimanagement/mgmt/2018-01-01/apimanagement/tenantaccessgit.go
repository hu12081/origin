package apimanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// TenantAccessGitClient is the apiManagement Client
type TenantAccessGitClient struct {
	BaseClient
}

// NewTenantAccessGitClient creates an instance of the TenantAccessGitClient client.
func NewTenantAccessGitClient(subscriptionID string) TenantAccessGitClient {
	return NewTenantAccessGitClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTenantAccessGitClientWithBaseURI creates an instance of the TenantAccessGitClient client.
func NewTenantAccessGitClientWithBaseURI(baseURI string, subscriptionID string) TenantAccessGitClient {
	return TenantAccessGitClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the Git access configuration for the tenant.
//
// resourceGroupName is the name of the resource group. serviceName is the name of the API Management service.
func (client TenantAccessGitClient) Get(ctx context.Context, resourceGroupName string, serviceName string) (result AccessInformationContract, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantAccessGitClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessGitClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessGitClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessGitClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client TenantAccessGitClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessName":        autorest.Encode("path", "access"),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/git", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TenantAccessGitClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TenantAccessGitClient) GetResponder(resp *http.Response) (result AccessInformationContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RegeneratePrimaryKey regenerate primary access key for GIT.
//
// resourceGroupName is the name of the resource group. serviceName is the name of the API Management service.
func (client TenantAccessGitClient) RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantAccessGitClient", "RegeneratePrimaryKey", err.Error())
	}

	req, err := client.RegeneratePrimaryKeyPreparer(ctx, resourceGroupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessGitClient", "RegeneratePrimaryKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegeneratePrimaryKeySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessGitClient", "RegeneratePrimaryKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegeneratePrimaryKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessGitClient", "RegeneratePrimaryKey", resp, "Failure responding to request")
	}

	return
}

// RegeneratePrimaryKeyPreparer prepares the RegeneratePrimaryKey request.
func (client TenantAccessGitClient) RegeneratePrimaryKeyPreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessName":        autorest.Encode("path", "access"),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/git/regeneratePrimaryKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegeneratePrimaryKeySender sends the RegeneratePrimaryKey request. The method will close the
// http.Response Body if it receives an error.
func (client TenantAccessGitClient) RegeneratePrimaryKeySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// RegeneratePrimaryKeyResponder handles the response to the RegeneratePrimaryKey request. The method always
// closes the http.Response Body.
func (client TenantAccessGitClient) RegeneratePrimaryKeyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// RegenerateSecondaryKey regenerate secondary access key for GIT.
//
// resourceGroupName is the name of the resource group. serviceName is the name of the API Management service.
func (client TenantAccessGitClient) RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantAccessGitClient", "RegenerateSecondaryKey", err.Error())
	}

	req, err := client.RegenerateSecondaryKeyPreparer(ctx, resourceGroupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessGitClient", "RegenerateSecondaryKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateSecondaryKeySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessGitClient", "RegenerateSecondaryKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateSecondaryKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantAccessGitClient", "RegenerateSecondaryKey", resp, "Failure responding to request")
	}

	return
}

// RegenerateSecondaryKeyPreparer prepares the RegenerateSecondaryKey request.
func (client TenantAccessGitClient) RegenerateSecondaryKeyPreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessName":        autorest.Encode("path", "access"),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/git/regenerateSecondaryKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateSecondaryKeySender sends the RegenerateSecondaryKey request. The method will close the
// http.Response Body if it receives an error.
func (client TenantAccessGitClient) RegenerateSecondaryKeySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// RegenerateSecondaryKeyResponder handles the response to the RegenerateSecondaryKey request. The method always
// closes the http.Response Body.
func (client TenantAccessGitClient) RegenerateSecondaryKeyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
