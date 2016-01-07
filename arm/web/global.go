package web

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

// GlobalClient is the use these APIs to manage Azure Websites resources
// through the Azure Resource Manager. All task operations conform to the
// HTTP/1.1 protocol specification and each operation returns an
// x-ms-request-id header that can be used to obtain information about the
// request. You must make sure that requests made to these resources are
// secure. For more information, see <a
// href="https://msdn.microsoft.com/en-us/library/azure/dn790557.aspx">Authenticating
// Azure Resource Manager requests.</a>
type GlobalClient struct {
	ManagementClient
}

// NewGlobalClient creates an instance of the GlobalClient client.
func NewGlobalClient(subscriptionID string) GlobalClient {
	return NewGlobalClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGlobalClientWithBaseURI creates an instance of the GlobalClient client.
func NewGlobalClientWithBaseURI(baseURI string, subscriptionID string) GlobalClient {
	return GlobalClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability sends the check name availability request.
//
// request is name availability request
func (client GlobalClient) CheckNameAvailability(request ResourceNameAvailabilityRequest) (result ResourceNameAvailability, ae error) {
	req, err := client.CheckNameAvailabilityPreparer(request)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "CheckNameAvailability", "Failure preparing request")
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "CheckNameAvailability", "Failure sending request")
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/GlobalClient", "CheckNameAvailability", "Failure responding to request")
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client GlobalClient) CheckNameAvailabilityPreparer(request ResourceNameAvailabilityRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Web/checknameavailability"),
		autorest.WithJSON(request),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client GlobalClient) CheckNameAvailabilityResponder(resp *http.Response) (result ResourceNameAvailability, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAllClassicMobileServices sends the get all classic mobile services
// request.
func (client GlobalClient) GetAllClassicMobileServices() (result ClassicMobileServiceCollection, ae error) {
	req, err := client.GetAllClassicMobileServicesPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllClassicMobileServices", "Failure preparing request")
	}

	resp, err := client.GetAllClassicMobileServicesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllClassicMobileServices", "Failure sending request")
	}

	result, err = client.GetAllClassicMobileServicesResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllClassicMobileServices", "Failure responding to request")
	}

	return
}

// GetAllClassicMobileServicesPreparer prepares the GetAllClassicMobileServices request.
func (client GlobalClient) GetAllClassicMobileServicesPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Web/classicMobileServices"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetAllClassicMobileServicesSender sends the GetAllClassicMobileServices request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalClient) GetAllClassicMobileServicesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetAllClassicMobileServicesResponder handles the response to the GetAllClassicMobileServices request. The method always
// closes the http.Response Body.
func (client GlobalClient) GetAllClassicMobileServicesResponder(resp *http.Response) (result ClassicMobileServiceCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAllHostingEnvironments sends the get all hosting environments request.
func (client GlobalClient) GetAllHostingEnvironments() (result HostingEnvironmentCollection, ae error) {
	req, err := client.GetAllHostingEnvironmentsPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllHostingEnvironments", "Failure preparing request")
	}

	resp, err := client.GetAllHostingEnvironmentsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllHostingEnvironments", "Failure sending request")
	}

	result, err = client.GetAllHostingEnvironmentsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllHostingEnvironments", "Failure responding to request")
	}

	return
}

// GetAllHostingEnvironmentsPreparer prepares the GetAllHostingEnvironments request.
func (client GlobalClient) GetAllHostingEnvironmentsPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Web/hostingEnvironments"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetAllHostingEnvironmentsSender sends the GetAllHostingEnvironments request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalClient) GetAllHostingEnvironmentsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetAllHostingEnvironmentsResponder handles the response to the GetAllHostingEnvironments request. The method always
// closes the http.Response Body.
func (client GlobalClient) GetAllHostingEnvironmentsResponder(resp *http.Response) (result HostingEnvironmentCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAllManagedHostingEnvironments sends the get all managed hosting
// environments request.
func (client GlobalClient) GetAllManagedHostingEnvironments() (result ManagedHostingEnvironmentCollection, ae error) {
	req, err := client.GetAllManagedHostingEnvironmentsPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllManagedHostingEnvironments", "Failure preparing request")
	}

	resp, err := client.GetAllManagedHostingEnvironmentsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllManagedHostingEnvironments", "Failure sending request")
	}

	result, err = client.GetAllManagedHostingEnvironmentsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllManagedHostingEnvironments", "Failure responding to request")
	}

	return
}

// GetAllManagedHostingEnvironmentsPreparer prepares the GetAllManagedHostingEnvironments request.
func (client GlobalClient) GetAllManagedHostingEnvironmentsPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Web/managedHostingEnvironments"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetAllManagedHostingEnvironmentsSender sends the GetAllManagedHostingEnvironments request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalClient) GetAllManagedHostingEnvironmentsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetAllManagedHostingEnvironmentsResponder handles the response to the GetAllManagedHostingEnvironments request. The method always
// closes the http.Response Body.
func (client GlobalClient) GetAllManagedHostingEnvironmentsResponder(resp *http.Response) (result ManagedHostingEnvironmentCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAllServerFarms sends the get all server farms request.
//
// detailed is false to return a subset of App Service Plan properties, true
// to return all of the properties.
// Retrieval of all properties may increase the API latency.
func (client GlobalClient) GetAllServerFarms(detailed *bool) (result ServerFarmCollection, ae error) {
	req, err := client.GetAllServerFarmsPreparer(detailed)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllServerFarms", "Failure preparing request")
	}

	resp, err := client.GetAllServerFarmsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllServerFarms", "Failure sending request")
	}

	result, err = client.GetAllServerFarmsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllServerFarms", "Failure responding to request")
	}

	return
}

// GetAllServerFarmsPreparer prepares the GetAllServerFarms request.
func (client GlobalClient) GetAllServerFarmsPreparer(detailed *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if detailed != nil {
		queryParameters["detailed"] = detailed
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Web/serverfarms"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetAllServerFarmsSender sends the GetAllServerFarms request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalClient) GetAllServerFarmsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetAllServerFarmsResponder handles the response to the GetAllServerFarms request. The method always
// closes the http.Response Body.
func (client GlobalClient) GetAllServerFarmsResponder(resp *http.Response) (result ServerFarmCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAllSites sends the get all sites request.
func (client GlobalClient) GetAllSites() (result SiteCollection, ae error) {
	req, err := client.GetAllSitesPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllSites", "Failure preparing request")
	}

	resp, err := client.GetAllSitesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllSites", "Failure sending request")
	}

	result, err = client.GetAllSitesResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllSites", "Failure responding to request")
	}

	return
}

// GetAllSitesPreparer prepares the GetAllSites request.
func (client GlobalClient) GetAllSitesPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Web/sites"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetAllSitesSender sends the GetAllSites request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalClient) GetAllSitesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetAllSitesResponder handles the response to the GetAllSites request. The method always
// closes the http.Response Body.
func (client GlobalClient) GetAllSitesResponder(resp *http.Response) (result SiteCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAllWebHostingPlans sends the get all web hosting plans request.
//
// detailed is false to return a subset of App Service Plan properties, true
// to return all of the properties.
// Retrieval of all properties may increase the API latency.
func (client GlobalClient) GetAllWebHostingPlans(detailed *bool) (result ServerFarmCollection, ae error) {
	req, err := client.GetAllWebHostingPlansPreparer(detailed)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllWebHostingPlans", "Failure preparing request")
	}

	resp, err := client.GetAllWebHostingPlansSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllWebHostingPlans", "Failure sending request")
	}

	result, err = client.GetAllWebHostingPlansResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/GlobalClient", "GetAllWebHostingPlans", "Failure responding to request")
	}

	return
}

// GetAllWebHostingPlansPreparer prepares the GetAllWebHostingPlans request.
func (client GlobalClient) GetAllWebHostingPlansPreparer(detailed *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if detailed != nil {
		queryParameters["detailed"] = detailed
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Web/webhostingplans"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetAllWebHostingPlansSender sends the GetAllWebHostingPlans request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalClient) GetAllWebHostingPlansSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetAllWebHostingPlansResponder handles the response to the GetAllWebHostingPlans request. The method always
// closes the http.Response Body.
func (client GlobalClient) GetAllWebHostingPlansResponder(resp *http.Response) (result ServerFarmCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSubscriptionGeoRegions sends the get subscription geo regions request.
func (client GlobalClient) GetSubscriptionGeoRegions() (result GeoRegionCollection, ae error) {
	req, err := client.GetSubscriptionGeoRegionsPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetSubscriptionGeoRegions", "Failure preparing request")
	}

	resp, err := client.GetSubscriptionGeoRegionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetSubscriptionGeoRegions", "Failure sending request")
	}

	result, err = client.GetSubscriptionGeoRegionsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/GlobalClient", "GetSubscriptionGeoRegions", "Failure responding to request")
	}

	return
}

// GetSubscriptionGeoRegionsPreparer prepares the GetSubscriptionGeoRegions request.
func (client GlobalClient) GetSubscriptionGeoRegionsPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Web/geoRegions"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSubscriptionGeoRegionsSender sends the GetSubscriptionGeoRegions request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalClient) GetSubscriptionGeoRegionsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetSubscriptionGeoRegionsResponder handles the response to the GetSubscriptionGeoRegions request. The method always
// closes the http.Response Body.
func (client GlobalClient) GetSubscriptionGeoRegionsResponder(resp *http.Response) (result GeoRegionCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSubscriptionPublishingCredentials sends the get subscription publishing
// credentials request.
func (client GlobalClient) GetSubscriptionPublishingCredentials() (result User, ae error) {
	req, err := client.GetSubscriptionPublishingCredentialsPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetSubscriptionPublishingCredentials", "Failure preparing request")
	}

	resp, err := client.GetSubscriptionPublishingCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "GetSubscriptionPublishingCredentials", "Failure sending request")
	}

	result, err = client.GetSubscriptionPublishingCredentialsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/GlobalClient", "GetSubscriptionPublishingCredentials", "Failure responding to request")
	}

	return
}

// GetSubscriptionPublishingCredentialsPreparer prepares the GetSubscriptionPublishingCredentials request.
func (client GlobalClient) GetSubscriptionPublishingCredentialsPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Web/publishingCredentials"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSubscriptionPublishingCredentialsSender sends the GetSubscriptionPublishingCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalClient) GetSubscriptionPublishingCredentialsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetSubscriptionPublishingCredentialsResponder handles the response to the GetSubscriptionPublishingCredentials request. The method always
// closes the http.Response Body.
func (client GlobalClient) GetSubscriptionPublishingCredentialsResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// IsHostingEnvironmentNameAvailable sends the is hosting environment name
// available request.
//
// name is hosting environment name
func (client GlobalClient) IsHostingEnvironmentNameAvailable(name string) (result ObjectSet, ae error) {
	req, err := client.IsHostingEnvironmentNameAvailablePreparer(name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "IsHostingEnvironmentNameAvailable", "Failure preparing request")
	}

	resp, err := client.IsHostingEnvironmentNameAvailableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "IsHostingEnvironmentNameAvailable", "Failure sending request")
	}

	result, err = client.IsHostingEnvironmentNameAvailableResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/GlobalClient", "IsHostingEnvironmentNameAvailable", "Failure responding to request")
	}

	return
}

// IsHostingEnvironmentNameAvailablePreparer prepares the IsHostingEnvironmentNameAvailable request.
func (client GlobalClient) IsHostingEnvironmentNameAvailablePreparer(name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"name":        name,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Web/ishostingenvironmentnameavailable"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// IsHostingEnvironmentNameAvailableSender sends the IsHostingEnvironmentNameAvailable request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalClient) IsHostingEnvironmentNameAvailableSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// IsHostingEnvironmentNameAvailableResponder handles the response to the IsHostingEnvironmentNameAvailable request. The method always
// closes the http.Response Body.
func (client GlobalClient) IsHostingEnvironmentNameAvailableResponder(resp *http.Response) (result ObjectSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// IsHostingEnvironmentWithLegacyNameAvailable sends the is hosting
// environment with legacy name available request.
//
// name is hosting environment name
func (client GlobalClient) IsHostingEnvironmentWithLegacyNameAvailable(name string) (result ObjectSet, ae error) {
	req, err := client.IsHostingEnvironmentWithLegacyNameAvailablePreparer(name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "IsHostingEnvironmentWithLegacyNameAvailable", "Failure preparing request")
	}

	resp, err := client.IsHostingEnvironmentWithLegacyNameAvailableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "IsHostingEnvironmentWithLegacyNameAvailable", "Failure sending request")
	}

	result, err = client.IsHostingEnvironmentWithLegacyNameAvailableResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/GlobalClient", "IsHostingEnvironmentWithLegacyNameAvailable", "Failure responding to request")
	}

	return
}

// IsHostingEnvironmentWithLegacyNameAvailablePreparer prepares the IsHostingEnvironmentWithLegacyNameAvailable request.
func (client GlobalClient) IsHostingEnvironmentWithLegacyNameAvailablePreparer(name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":           url.QueryEscape(name),
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Web/ishostingenvironmentnameavailable/{name}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// IsHostingEnvironmentWithLegacyNameAvailableSender sends the IsHostingEnvironmentWithLegacyNameAvailable request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalClient) IsHostingEnvironmentWithLegacyNameAvailableSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// IsHostingEnvironmentWithLegacyNameAvailableResponder handles the response to the IsHostingEnvironmentWithLegacyNameAvailable request. The method always
// closes the http.Response Body.
func (client GlobalClient) IsHostingEnvironmentWithLegacyNameAvailableResponder(resp *http.Response) (result ObjectSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListPremierAddOnOffers sends the list premier add on offers request.
func (client GlobalClient) ListPremierAddOnOffers() (result ObjectSet, ae error) {
	req, err := client.ListPremierAddOnOffersPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "ListPremierAddOnOffers", "Failure preparing request")
	}

	resp, err := client.ListPremierAddOnOffersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "ListPremierAddOnOffers", "Failure sending request")
	}

	result, err = client.ListPremierAddOnOffersResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/GlobalClient", "ListPremierAddOnOffers", "Failure responding to request")
	}

	return
}

// ListPremierAddOnOffersPreparer prepares the ListPremierAddOnOffers request.
func (client GlobalClient) ListPremierAddOnOffersPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Web/premieraddonoffers"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListPremierAddOnOffersSender sends the ListPremierAddOnOffers request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalClient) ListPremierAddOnOffersSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListPremierAddOnOffersResponder handles the response to the ListPremierAddOnOffers request. The method always
// closes the http.Response Body.
func (client GlobalClient) ListPremierAddOnOffersResponder(resp *http.Response) (result ObjectSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateSubscriptionPublishingCredentials sends the update subscription
// publishing credentials request.
//
// requestMessage is requestMessage with new publishing credentials
func (client GlobalClient) UpdateSubscriptionPublishingCredentials(requestMessage User) (result User, ae error) {
	req, err := client.UpdateSubscriptionPublishingCredentialsPreparer(requestMessage)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "UpdateSubscriptionPublishingCredentials", "Failure preparing request")
	}

	resp, err := client.UpdateSubscriptionPublishingCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/GlobalClient", "UpdateSubscriptionPublishingCredentials", "Failure sending request")
	}

	result, err = client.UpdateSubscriptionPublishingCredentialsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/GlobalClient", "UpdateSubscriptionPublishingCredentials", "Failure responding to request")
	}

	return
}

// UpdateSubscriptionPublishingCredentialsPreparer prepares the UpdateSubscriptionPublishingCredentials request.
func (client GlobalClient) UpdateSubscriptionPublishingCredentialsPreparer(requestMessage User) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Web/publishingCredentials"),
		autorest.WithJSON(requestMessage),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// UpdateSubscriptionPublishingCredentialsSender sends the UpdateSubscriptionPublishingCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalClient) UpdateSubscriptionPublishingCredentialsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// UpdateSubscriptionPublishingCredentialsResponder handles the response to the UpdateSubscriptionPublishingCredentials request. The method always
// closes the http.Response Body.
func (client GlobalClient) UpdateSubscriptionPublishingCredentialsResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
