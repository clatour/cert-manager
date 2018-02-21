package servicefabric

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

// ClustersClient is the client for the Clusters methods of the Servicefabric service.
type ClustersClient struct {
	BaseClient
}

// NewClustersClient creates an instance of the ClustersClient client.
func NewClustersClient(subscriptionID string) ClustersClient {
	return NewClustersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClustersClientWithBaseURI creates an instance of the ClustersClient client.
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return ClustersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create cluster resource
//
// resourceGroupName is the name of the resource group to which the resource belongs or get created clusterName is the
// name of the cluster resource parameters is put Request
func (client ClustersClient) Create(ctx context.Context, resourceGroupName string, clusterName string, parameters Cluster) (result ClustersCreateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ClusterProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ClusterProperties.Certificate", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ClusterProperties.Certificate.Thumbprint", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.ClusterProperties.ReverseProxyCertificate", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ClusterProperties.ReverseProxyCertificate.Thumbprint", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.ClusterProperties.ManagementEndpoint", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.ClusterProperties.NodeTypes", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.ClusterProperties.DiagnosticsStorageAccountConfig", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ClusterProperties.DiagnosticsStorageAccountConfig.StorageAccountName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ClusterProperties.DiagnosticsStorageAccountConfig.ProtectedAccountKeyName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ClusterProperties.DiagnosticsStorageAccountConfig.BlobEndpoint", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ClusterProperties.DiagnosticsStorageAccountConfig.QueueEndpoint", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ClusterProperties.DiagnosticsStorageAccountConfig.TableEndpoint", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "parameters.ClusterProperties.UpgradeDescription", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ClusterProperties.UpgradeDescription.UpgradeReplicaSetCheckTimeout", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ClusterProperties.UpgradeDescription.HealthCheckWaitDuration", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ClusterProperties.UpgradeDescription.HealthCheckStableDuration", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ClusterProperties.UpgradeDescription.HealthCheckRetryTimeout", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ClusterProperties.UpgradeDescription.UpgradeTimeout", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ClusterProperties.UpgradeDescription.UpgradeDomainTimeout", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ClusterProperties.UpgradeDescription.HealthPolicy", Name: validation.Null, Rule: true,
								Chain: []validation.Constraint{{Target: "parameters.ClusterProperties.UpgradeDescription.HealthPolicy.MaxPercentUnhealthyNodes", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.ClusterProperties.UpgradeDescription.HealthPolicy.MaxPercentUnhealthyNodes", Name: validation.InclusiveMaximum, Rule: 100, Chain: nil},
										{Target: "parameters.ClusterProperties.UpgradeDescription.HealthPolicy.MaxPercentUnhealthyNodes", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
									}},
									{Target: "parameters.ClusterProperties.UpgradeDescription.HealthPolicy.MaxPercentUnhealthyApplications", Name: validation.Null, Rule: false,
										Chain: []validation.Constraint{{Target: "parameters.ClusterProperties.UpgradeDescription.HealthPolicy.MaxPercentUnhealthyApplications", Name: validation.InclusiveMaximum, Rule: 100, Chain: nil},
											{Target: "parameters.ClusterProperties.UpgradeDescription.HealthPolicy.MaxPercentUnhealthyApplications", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
										}},
								}},
							{Target: "parameters.ClusterProperties.UpgradeDescription.DeltaHealthPolicy", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ClusterProperties.UpgradeDescription.DeltaHealthPolicy.MaxPercentDeltaUnhealthyNodes", Name: validation.Null, Rule: true,
									Chain: []validation.Constraint{{Target: "parameters.ClusterProperties.UpgradeDescription.DeltaHealthPolicy.MaxPercentDeltaUnhealthyNodes", Name: validation.InclusiveMaximum, Rule: 100, Chain: nil},
										{Target: "parameters.ClusterProperties.UpgradeDescription.DeltaHealthPolicy.MaxPercentDeltaUnhealthyNodes", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
									}},
									{Target: "parameters.ClusterProperties.UpgradeDescription.DeltaHealthPolicy.MaxPercentUpgradeDomainDeltaUnhealthyNodes", Name: validation.Null, Rule: true,
										Chain: []validation.Constraint{{Target: "parameters.ClusterProperties.UpgradeDescription.DeltaHealthPolicy.MaxPercentUpgradeDomainDeltaUnhealthyNodes", Name: validation.InclusiveMaximum, Rule: 100, Chain: nil},
											{Target: "parameters.ClusterProperties.UpgradeDescription.DeltaHealthPolicy.MaxPercentUpgradeDomainDeltaUnhealthyNodes", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
										}},
									{Target: "parameters.ClusterProperties.UpgradeDescription.DeltaHealthPolicy.MaxPercentDeltaUnhealthyApplications", Name: validation.Null, Rule: true,
										Chain: []validation.Constraint{{Target: "parameters.ClusterProperties.UpgradeDescription.DeltaHealthPolicy.MaxPercentDeltaUnhealthyApplications", Name: validation.InclusiveMaximum, Rule: 100, Chain: nil},
											{Target: "parameters.ClusterProperties.UpgradeDescription.DeltaHealthPolicy.MaxPercentDeltaUnhealthyApplications", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
										}},
								}},
						}},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servicefabric.ClustersClient", "Create")
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ClustersClient) CreatePreparer(ctx context.Context, resourceGroupName string, clusterName string, parameters Cluster) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) CreateSender(req *http.Request) (future ClustersCreateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ClustersClient) CreateResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete cluster resource
//
// resourceGroupName is the name of the resource group to which the resource belongs or get created clusterName is the
// name of the cluster resource
func (client ClustersClient) Delete(ctx context.Context, resourceGroupName string, clusterName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ClustersClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ClustersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get cluster resource
//
// resourceGroupName is the name of the resource group to which the resource belongs or get created clusterName is the
// name of the cluster resource
func (client ClustersClient) Get(ctx context.Context, resourceGroupName string, clusterName string) (result Cluster, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ClustersClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ClustersClient) GetResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list cluster resource
func (client ClustersClient) List(ctx context.Context) (result ClusterListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "List", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ClustersClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/clusters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListResponder(resp *http.Response) (result ClusterListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ClustersClient) listNextResults(lastResults ClusterListResult) (result ClusterListResult, err error) {
	req, err := lastResults.clusterListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ClustersClient) ListComplete(ctx context.Context) (result ClusterListResultIterator, err error) {
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup list cluster resource by resource group
//
// resourceGroupName is the name of the resource group to which the resource belongs or get created
func (client ClustersClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ClusterListResultPage, err error) {
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ClustersClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListByResourceGroupResponder(resp *http.Response) (result ClusterListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ClustersClient) listByResourceGroupNextResults(lastResults ClusterListResult) (result ClusterListResult, err error) {
	req, err := lastResults.clusterListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ClustersClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ClusterListResultIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Update update cluster configuration
//
// resourceGroupName is the name of the resource group to which the resource belongs or get created clusterName is the
// name of the cluster resource parameters is the parameters which contains the property value and property name which
// used to update the cluster configuration
func (client ClustersClient) Update(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterUpdateParameters) (result ClustersUpdateFuture, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClustersClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ClustersClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) UpdateSender(req *http.Request) (future ClustersUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ClustersClient) UpdateResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
