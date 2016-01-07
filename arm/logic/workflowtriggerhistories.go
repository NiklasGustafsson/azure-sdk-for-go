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

// WorkflowTriggerHistoriesClient is the client for the
// WorkflowTriggerHistories methods of the Logic service.
type WorkflowTriggerHistoriesClient struct {
	ManagementClient
}

// NewWorkflowTriggerHistoriesClient creates an instance of the
// WorkflowTriggerHistoriesClient client.
func NewWorkflowTriggerHistoriesClient(subscriptionID string) WorkflowTriggerHistoriesClient {
	return NewWorkflowTriggerHistoriesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowTriggerHistoriesClientWithBaseURI creates an instance of the
// WorkflowTriggerHistoriesClient client.
func NewWorkflowTriggerHistoriesClientWithBaseURI(baseURI string, subscriptionID string) WorkflowTriggerHistoriesClient {
	return WorkflowTriggerHistoriesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a workflow trigger history.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. triggerName is the workflow trigger name. historyName is the
// workflow trigger history name.
func (client WorkflowTriggerHistoriesClient) Get(resourceGroupName string, workflowName string, triggerName string, historyName string) (result WorkflowTriggerHistory, ae error) {
	req, err := client.GetPreparer(resourceGroupName, workflowName, triggerName, historyName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggerHistoriesClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggerHistoriesClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowTriggerHistoriesClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowTriggerHistoriesClient) GetPreparer(resourceGroupName string, workflowName string, triggerName string, historyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"historyName":       url.QueryEscape(historyName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"triggerName":       url.QueryEscape(triggerName),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/histories/{historyName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggerHistoriesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowTriggerHistoriesClient) GetResponder(resp *http.Response) (result WorkflowTriggerHistory, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of workflow trigger histories.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. triggerName is the workflow trigger name. top is the number of items
// to be included in the result.
func (client WorkflowTriggerHistoriesClient) List(resourceGroupName string, workflowName string, triggerName string, top *int) (result WorkflowTriggerHistoryListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, workflowName, triggerName, top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggerHistoriesClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggerHistoriesClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowTriggerHistoriesClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkflowTriggerHistoriesClient) ListPreparer(resourceGroupName string, workflowName string, triggerName string, top *int) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"triggerName":       url.QueryEscape(triggerName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/histories"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggerHistoriesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkflowTriggerHistoriesClient) ListResponder(resp *http.Response) (result WorkflowTriggerHistoryListResult, err error) {
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
func (client WorkflowTriggerHistoriesClient) ListNextResults(lastResults WorkflowTriggerHistoryListResult) (result WorkflowTriggerHistoryListResult, ae error) {
	req, err := lastResults.WorkflowTriggerHistoryListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggerHistoriesClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggerHistoriesClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowTriggerHistoriesClient", "List", "Failure responding to next results request request")
	}

	return
}
