package logic

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.12.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/NiklasGustafsson/azure-sdk-for-go/Godeps/_workspace/src/github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// WorkflowAccessKeysClient is the client for the WorkflowAccessKeys methods
// of the Logic service.
type WorkflowAccessKeysClient struct {
	ManagementClient
}

// NewWorkflowAccessKeysClient creates an instance of the
// WorkflowAccessKeysClient client.
func NewWorkflowAccessKeysClient(subscriptionID string) WorkflowAccessKeysClient {
	return NewWorkflowAccessKeysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowAccessKeysClientWithBaseURI creates an instance of the
// WorkflowAccessKeysClient client.
func NewWorkflowAccessKeysClientWithBaseURI(baseURI string, subscriptionID string) WorkflowAccessKeysClient {
	return WorkflowAccessKeysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a workflow access key.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. accessKeyName is the workflow access key name. workflowAccesskey is
// the workflow access key.
func (client WorkflowAccessKeysClient) CreateOrUpdate(resourceGroupName string, workflowName string, accessKeyName string, workflowAccesskey WorkflowAccessKey) (result WorkflowAccessKey, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, workflowName, accessKeyName, workflowAccesskey)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WorkflowAccessKeysClient) CreateOrUpdatePreparer(resourceGroupName string, workflowName string, accessKeyName string, workflowAccesskey WorkflowAccessKey) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     url.QueryEscape(accessKeyName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}"),
		autorest.WithJSON(workflowAccesskey),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowAccessKeysClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusCreated)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client WorkflowAccessKeysClient) CreateOrUpdateResponder(resp *http.Response) (result WorkflowAccessKey, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a workflow access key.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. accessKeyName is the workflow access key name.
func (client WorkflowAccessKeysClient) Delete(resourceGroupName string, workflowName string, accessKeyName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, workflowName, accessKeyName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client WorkflowAccessKeysClient) DeletePreparer(resourceGroupName string, workflowName string, accessKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     url.QueryEscape(accessKeyName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowAccessKeysClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WorkflowAccessKeysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a workflow access key.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. accessKeyName is the workflow access key name.
func (client WorkflowAccessKeysClient) Get(resourceGroupName string, workflowName string, accessKeyName string) (result WorkflowAccessKey, ae error) {
	req, err := client.GetPreparer(resourceGroupName, workflowName, accessKeyName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowAccessKeysClient) GetPreparer(resourceGroupName string, workflowName string, accessKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     url.QueryEscape(accessKeyName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowAccessKeysClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowAccessKeysClient) GetResponder(resp *http.Response) (result WorkflowAccessKey, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of workflow access keys.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. top is the number of items to be included in the result.
func (client WorkflowAccessKeysClient) List(resourceGroupName string, workflowName string, top *int) (result WorkflowAccessKeyListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, workflowName, top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkflowAccessKeysClient) ListPreparer(resourceGroupName string, workflowName string, top *int) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = top
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowAccessKeysClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkflowAccessKeysClient) ListResponder(resp *http.Response) (result WorkflowAccessKeyListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client WorkflowAccessKeysClient) ListNextResults(lastResults WorkflowAccessKeyListResult) (result WorkflowAccessKeyListResult, ae error) {
	req, err := lastResults.WorkflowAccessKeyListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "List", "Failure responding to next results request request")
	}

	return
}

// ListSecretKeys lists secret keys.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. accessKeyName is the workflow access key name.
func (client WorkflowAccessKeysClient) ListSecretKeys(resourceGroupName string, workflowName string, accessKeyName string) (result WorkflowSecretKeys, ae error) {
	req, err := client.ListSecretKeysPreparer(resourceGroupName, workflowName, accessKeyName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "ListSecretKeys", "Failure preparing request")
	}

	resp, err := client.ListSecretKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "ListSecretKeys", "Failure sending request")
	}

	result, err = client.ListSecretKeysResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "ListSecretKeys", "Failure responding to request")
	}

	return
}

// ListSecretKeysPreparer prepares the ListSecretKeys request.
func (client WorkflowAccessKeysClient) ListSecretKeysPreparer(resourceGroupName string, workflowName string, accessKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     url.QueryEscape(accessKeyName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}/list"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSecretKeysSender sends the ListSecretKeys request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowAccessKeysClient) ListSecretKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListSecretKeysResponder handles the response to the ListSecretKeys request. The method always
// closes the http.Response Body.
func (client WorkflowAccessKeysClient) ListSecretKeysResponder(resp *http.Response) (result WorkflowSecretKeys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RegenerateSecretKey regenerates secret key.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. accessKeyName is the workflow access key name. parameters is the
// parameters.
func (client WorkflowAccessKeysClient) RegenerateSecretKey(resourceGroupName string, workflowName string, accessKeyName string, parameters RegenerateSecretKeyParameters) (result WorkflowSecretKeys, ae error) {
	req, err := client.RegenerateSecretKeyPreparer(resourceGroupName, workflowName, accessKeyName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "RegenerateSecretKey", "Failure preparing request")
	}

	resp, err := client.RegenerateSecretKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "RegenerateSecretKey", "Failure sending request")
	}

	result, err = client.RegenerateSecretKeyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowAccessKeysClient", "RegenerateSecretKey", "Failure responding to request")
	}

	return
}

// RegenerateSecretKeyPreparer prepares the RegenerateSecretKey request.
func (client WorkflowAccessKeysClient) RegenerateSecretKeyPreparer(resourceGroupName string, workflowName string, accessKeyName string, parameters RegenerateSecretKeyParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     url.QueryEscape(accessKeyName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}/regenerate"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// RegenerateSecretKeySender sends the RegenerateSecretKey request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowAccessKeysClient) RegenerateSecretKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// RegenerateSecretKeyResponder handles the response to the RegenerateSecretKey request. The method always
// closes the http.Response Body.
func (client WorkflowAccessKeysClient) RegenerateSecretKeyResponder(resp *http.Response) (result WorkflowSecretKeys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
