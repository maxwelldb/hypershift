//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

// NewAccessReviewDefaultSettingsClient creates a new instance of AccessReviewDefaultSettingsClient.
func (c *ClientFactory) NewAccessReviewDefaultSettingsClient() *AccessReviewDefaultSettingsClient {
	subClient, _ := NewAccessReviewDefaultSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAccessReviewHistoryDefinitionClient creates a new instance of AccessReviewHistoryDefinitionClient.
func (c *ClientFactory) NewAccessReviewHistoryDefinitionClient() *AccessReviewHistoryDefinitionClient {
	subClient, _ := NewAccessReviewHistoryDefinitionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAccessReviewHistoryDefinitionInstanceClient creates a new instance of AccessReviewHistoryDefinitionInstanceClient.
func (c *ClientFactory) NewAccessReviewHistoryDefinitionInstanceClient() *AccessReviewHistoryDefinitionInstanceClient {
	subClient, _ := NewAccessReviewHistoryDefinitionInstanceClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAccessReviewHistoryDefinitionInstancesClient creates a new instance of AccessReviewHistoryDefinitionInstancesClient.
func (c *ClientFactory) NewAccessReviewHistoryDefinitionInstancesClient() *AccessReviewHistoryDefinitionInstancesClient {
	subClient, _ := NewAccessReviewHistoryDefinitionInstancesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAccessReviewHistoryDefinitionsClient creates a new instance of AccessReviewHistoryDefinitionsClient.
func (c *ClientFactory) NewAccessReviewHistoryDefinitionsClient() *AccessReviewHistoryDefinitionsClient {
	subClient, _ := NewAccessReviewHistoryDefinitionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAccessReviewInstanceClient creates a new instance of AccessReviewInstanceClient.
func (c *ClientFactory) NewAccessReviewInstanceClient() *AccessReviewInstanceClient {
	subClient, _ := NewAccessReviewInstanceClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAccessReviewInstanceContactedReviewersClient creates a new instance of AccessReviewInstanceContactedReviewersClient.
func (c *ClientFactory) NewAccessReviewInstanceContactedReviewersClient() *AccessReviewInstanceContactedReviewersClient {
	subClient, _ := NewAccessReviewInstanceContactedReviewersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAccessReviewInstanceDecisionsClient creates a new instance of AccessReviewInstanceDecisionsClient.
func (c *ClientFactory) NewAccessReviewInstanceDecisionsClient() *AccessReviewInstanceDecisionsClient {
	subClient, _ := NewAccessReviewInstanceDecisionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAccessReviewInstanceMyDecisionsClient creates a new instance of AccessReviewInstanceMyDecisionsClient.
func (c *ClientFactory) NewAccessReviewInstanceMyDecisionsClient() *AccessReviewInstanceMyDecisionsClient {
	subClient, _ := NewAccessReviewInstanceMyDecisionsClient(c.credential, c.options)
	return subClient
}

// NewAccessReviewInstancesAssignedForMyApprovalClient creates a new instance of AccessReviewInstancesAssignedForMyApprovalClient.
func (c *ClientFactory) NewAccessReviewInstancesAssignedForMyApprovalClient() *AccessReviewInstancesAssignedForMyApprovalClient {
	subClient, _ := NewAccessReviewInstancesAssignedForMyApprovalClient(c.credential, c.options)
	return subClient
}

// NewAccessReviewInstancesClient creates a new instance of AccessReviewInstancesClient.
func (c *ClientFactory) NewAccessReviewInstancesClient() *AccessReviewInstancesClient {
	subClient, _ := NewAccessReviewInstancesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAccessReviewScheduleDefinitionsAssignedForMyApprovalClient creates a new instance of AccessReviewScheduleDefinitionsAssignedForMyApprovalClient.
func (c *ClientFactory) NewAccessReviewScheduleDefinitionsAssignedForMyApprovalClient() *AccessReviewScheduleDefinitionsAssignedForMyApprovalClient {
	subClient, _ := NewAccessReviewScheduleDefinitionsAssignedForMyApprovalClient(c.credential, c.options)
	return subClient
}

// NewAccessReviewScheduleDefinitionsClient creates a new instance of AccessReviewScheduleDefinitionsClient.
func (c *ClientFactory) NewAccessReviewScheduleDefinitionsClient() *AccessReviewScheduleDefinitionsClient {
	subClient, _ := NewAccessReviewScheduleDefinitionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAlertConfigurationsClient creates a new instance of AlertConfigurationsClient.
func (c *ClientFactory) NewAlertConfigurationsClient() *AlertConfigurationsClient {
	subClient, _ := NewAlertConfigurationsClient(c.credential, c.options)
	return subClient
}

// NewAlertDefinitionsClient creates a new instance of AlertDefinitionsClient.
func (c *ClientFactory) NewAlertDefinitionsClient() *AlertDefinitionsClient {
	subClient, _ := NewAlertDefinitionsClient(c.credential, c.options)
	return subClient
}

// NewAlertIncidentsClient creates a new instance of AlertIncidentsClient.
func (c *ClientFactory) NewAlertIncidentsClient() *AlertIncidentsClient {
	subClient, _ := NewAlertIncidentsClient(c.credential, c.options)
	return subClient
}

// NewAlertOperationClient creates a new instance of AlertOperationClient.
func (c *ClientFactory) NewAlertOperationClient() *AlertOperationClient {
	subClient, _ := NewAlertOperationClient(c.credential, c.options)
	return subClient
}

// NewAlertsClient creates a new instance of AlertsClient.
func (c *ClientFactory) NewAlertsClient() *AlertsClient {
	subClient, _ := NewAlertsClient(c.credential, c.options)
	return subClient
}

// NewClassicAdministratorsClient creates a new instance of ClassicAdministratorsClient.
func (c *ClientFactory) NewClassicAdministratorsClient() *ClassicAdministratorsClient {
	subClient, _ := NewClassicAdministratorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDenyAssignmentsClient creates a new instance of DenyAssignmentsClient.
func (c *ClientFactory) NewDenyAssignmentsClient() *DenyAssignmentsClient {
	subClient, _ := NewDenyAssignmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewEligibleChildResourcesClient creates a new instance of EligibleChildResourcesClient.
func (c *ClientFactory) NewEligibleChildResourcesClient() *EligibleChildResourcesClient {
	subClient, _ := NewEligibleChildResourcesClient(c.credential, c.options)
	return subClient
}

// NewGlobalAdministratorClient creates a new instance of GlobalAdministratorClient.
func (c *ClientFactory) NewGlobalAdministratorClient() *GlobalAdministratorClient {
	subClient, _ := NewGlobalAdministratorClient(c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewPermissionsClient creates a new instance of PermissionsClient.
func (c *ClientFactory) NewPermissionsClient() *PermissionsClient {
	subClient, _ := NewPermissionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewProviderOperationsMetadataClient creates a new instance of ProviderOperationsMetadataClient.
func (c *ClientFactory) NewProviderOperationsMetadataClient() *ProviderOperationsMetadataClient {
	subClient, _ := NewProviderOperationsMetadataClient(c.credential, c.options)
	return subClient
}

// NewRoleAssignmentScheduleInstancesClient creates a new instance of RoleAssignmentScheduleInstancesClient.
func (c *ClientFactory) NewRoleAssignmentScheduleInstancesClient() *RoleAssignmentScheduleInstancesClient {
	subClient, _ := NewRoleAssignmentScheduleInstancesClient(c.credential, c.options)
	return subClient
}

// NewRoleAssignmentScheduleRequestsClient creates a new instance of RoleAssignmentScheduleRequestsClient.
func (c *ClientFactory) NewRoleAssignmentScheduleRequestsClient() *RoleAssignmentScheduleRequestsClient {
	subClient, _ := NewRoleAssignmentScheduleRequestsClient(c.credential, c.options)
	return subClient
}

// NewRoleAssignmentSchedulesClient creates a new instance of RoleAssignmentSchedulesClient.
func (c *ClientFactory) NewRoleAssignmentSchedulesClient() *RoleAssignmentSchedulesClient {
	subClient, _ := NewRoleAssignmentSchedulesClient(c.credential, c.options)
	return subClient
}

// NewRoleAssignmentsClient creates a new instance of RoleAssignmentsClient.
func (c *ClientFactory) NewRoleAssignmentsClient() *RoleAssignmentsClient {
	subClient, _ := NewRoleAssignmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewRoleDefinitionsClient creates a new instance of RoleDefinitionsClient.
func (c *ClientFactory) NewRoleDefinitionsClient() *RoleDefinitionsClient {
	subClient, _ := NewRoleDefinitionsClient(c.credential, c.options)
	return subClient
}

// NewRoleEligibilityScheduleInstancesClient creates a new instance of RoleEligibilityScheduleInstancesClient.
func (c *ClientFactory) NewRoleEligibilityScheduleInstancesClient() *RoleEligibilityScheduleInstancesClient {
	subClient, _ := NewRoleEligibilityScheduleInstancesClient(c.credential, c.options)
	return subClient
}

// NewRoleEligibilityScheduleRequestsClient creates a new instance of RoleEligibilityScheduleRequestsClient.
func (c *ClientFactory) NewRoleEligibilityScheduleRequestsClient() *RoleEligibilityScheduleRequestsClient {
	subClient, _ := NewRoleEligibilityScheduleRequestsClient(c.credential, c.options)
	return subClient
}

// NewRoleEligibilitySchedulesClient creates a new instance of RoleEligibilitySchedulesClient.
func (c *ClientFactory) NewRoleEligibilitySchedulesClient() *RoleEligibilitySchedulesClient {
	subClient, _ := NewRoleEligibilitySchedulesClient(c.credential, c.options)
	return subClient
}

// NewRoleManagementPoliciesClient creates a new instance of RoleManagementPoliciesClient.
func (c *ClientFactory) NewRoleManagementPoliciesClient() *RoleManagementPoliciesClient {
	subClient, _ := NewRoleManagementPoliciesClient(c.credential, c.options)
	return subClient
}

// NewRoleManagementPolicyAssignmentsClient creates a new instance of RoleManagementPolicyAssignmentsClient.
func (c *ClientFactory) NewRoleManagementPolicyAssignmentsClient() *RoleManagementPolicyAssignmentsClient {
	subClient, _ := NewRoleManagementPolicyAssignmentsClient(c.credential, c.options)
	return subClient
}

// NewScopeAccessReviewDefaultSettingsClient creates a new instance of ScopeAccessReviewDefaultSettingsClient.
func (c *ClientFactory) NewScopeAccessReviewDefaultSettingsClient() *ScopeAccessReviewDefaultSettingsClient {
	subClient, _ := NewScopeAccessReviewDefaultSettingsClient(c.credential, c.options)
	return subClient
}

// NewScopeAccessReviewHistoryDefinitionClient creates a new instance of ScopeAccessReviewHistoryDefinitionClient.
func (c *ClientFactory) NewScopeAccessReviewHistoryDefinitionClient() *ScopeAccessReviewHistoryDefinitionClient {
	subClient, _ := NewScopeAccessReviewHistoryDefinitionClient(c.credential, c.options)
	return subClient
}

// NewScopeAccessReviewHistoryDefinitionInstanceClient creates a new instance of ScopeAccessReviewHistoryDefinitionInstanceClient.
func (c *ClientFactory) NewScopeAccessReviewHistoryDefinitionInstanceClient() *ScopeAccessReviewHistoryDefinitionInstanceClient {
	subClient, _ := NewScopeAccessReviewHistoryDefinitionInstanceClient(c.credential, c.options)
	return subClient
}

// NewScopeAccessReviewHistoryDefinitionInstancesClient creates a new instance of ScopeAccessReviewHistoryDefinitionInstancesClient.
func (c *ClientFactory) NewScopeAccessReviewHistoryDefinitionInstancesClient() *ScopeAccessReviewHistoryDefinitionInstancesClient {
	subClient, _ := NewScopeAccessReviewHistoryDefinitionInstancesClient(c.credential, c.options)
	return subClient
}

// NewScopeAccessReviewHistoryDefinitionsClient creates a new instance of ScopeAccessReviewHistoryDefinitionsClient.
func (c *ClientFactory) NewScopeAccessReviewHistoryDefinitionsClient() *ScopeAccessReviewHistoryDefinitionsClient {
	subClient, _ := NewScopeAccessReviewHistoryDefinitionsClient(c.credential, c.options)
	return subClient
}

// NewScopeAccessReviewInstanceClient creates a new instance of ScopeAccessReviewInstanceClient.
func (c *ClientFactory) NewScopeAccessReviewInstanceClient() *ScopeAccessReviewInstanceClient {
	subClient, _ := NewScopeAccessReviewInstanceClient(c.credential, c.options)
	return subClient
}

// NewScopeAccessReviewInstanceContactedReviewersClient creates a new instance of ScopeAccessReviewInstanceContactedReviewersClient.
func (c *ClientFactory) NewScopeAccessReviewInstanceContactedReviewersClient() *ScopeAccessReviewInstanceContactedReviewersClient {
	subClient, _ := NewScopeAccessReviewInstanceContactedReviewersClient(c.credential, c.options)
	return subClient
}

// NewScopeAccessReviewInstanceDecisionsClient creates a new instance of ScopeAccessReviewInstanceDecisionsClient.
func (c *ClientFactory) NewScopeAccessReviewInstanceDecisionsClient() *ScopeAccessReviewInstanceDecisionsClient {
	subClient, _ := NewScopeAccessReviewInstanceDecisionsClient(c.credential, c.options)
	return subClient
}

// NewScopeAccessReviewInstancesClient creates a new instance of ScopeAccessReviewInstancesClient.
func (c *ClientFactory) NewScopeAccessReviewInstancesClient() *ScopeAccessReviewInstancesClient {
	subClient, _ := NewScopeAccessReviewInstancesClient(c.credential, c.options)
	return subClient
}

// NewScopeAccessReviewScheduleDefinitionsClient creates a new instance of ScopeAccessReviewScheduleDefinitionsClient.
func (c *ClientFactory) NewScopeAccessReviewScheduleDefinitionsClient() *ScopeAccessReviewScheduleDefinitionsClient {
	subClient, _ := NewScopeAccessReviewScheduleDefinitionsClient(c.credential, c.options)
	return subClient
}

// NewTenantLevelAccessReviewInstanceContactedReviewersClient creates a new instance of TenantLevelAccessReviewInstanceContactedReviewersClient.
func (c *ClientFactory) NewTenantLevelAccessReviewInstanceContactedReviewersClient() *TenantLevelAccessReviewInstanceContactedReviewersClient {
	subClient, _ := NewTenantLevelAccessReviewInstanceContactedReviewersClient(c.credential, c.options)
	return subClient
}