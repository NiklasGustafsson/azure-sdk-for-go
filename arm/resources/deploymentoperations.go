package resources

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
	"github.com/Azure/azure-sdk-for-go/Godeps/_workspace/src/github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// DeploymentOperationsClient is the client for the DeploymentOperations
// methods of the Resources service.
type DeploymentOperationsClient struct {
	ManagementClient
}

// NewDeploymentOperationsClient creates an instance of the
// DeploymentOperationsClient client.
func NewDeploymentOperationsClient(subscriptionID string) DeploymentOperationsClient {
	return NewDeploymentOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDeploymentOperationsClientWithBaseURI creates an instance of the
// DeploymentOperationsClient client.
func NewDeploymentOperationsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentOperationsClient {
	return DeploymentOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a list of deployments operations.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. deploymentName is the name of the deployment. operationID is
// operation Id.
func (client DeploymentOperationsClient) Get(resourceGroupName string, deploymentName string, operationID string) (result DeploymentOperation, ae error) {
	req, err := client.GetPreparer(resourceGroupName, deploymentName, operationID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/DeploymentOperationsClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/DeploymentOperationsClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/DeploymentOperationsClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DeploymentOperationsClient) GetPreparer(resourceGroupName string, deploymentName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    url.QueryEscape(deploymentName),
		"operationId":       url.QueryEscape(operationID),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations/{operationId}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentOperationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeploymentOperationsClient) GetResponder(resp *http.Response) (result DeploymentOperation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of deployments operations.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. deploymentName is the name of the deployment. top is query
// parameters.
func (client DeploymentOperationsClient) List(resourceGroupName string, deploymentName string, top *int) (result DeploymentOperationsListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, deploymentName, top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/DeploymentOperationsClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/DeploymentOperationsClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/DeploymentOperationsClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DeploymentOperationsClient) ListPreparer(resourceGroupName string, deploymentName string, top *int) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    url.QueryEscape(deploymentName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DeploymentOperationsClient) ListResponder(resp *http.Response) (result DeploymentOperationsListResult, err error) {
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
func (client DeploymentOperationsClient) ListNextResults(lastResults DeploymentOperationsListResult) (result DeploymentOperationsListResult, ae error) {
	req, err := lastResults.DeploymentOperationsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/DeploymentOperationsClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/DeploymentOperationsClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/DeploymentOperationsClient", "List", "Failure responding to next results request request")
	}

	return
}
