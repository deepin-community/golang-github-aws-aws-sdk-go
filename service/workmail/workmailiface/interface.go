// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package workmailiface provides an interface to enable mocking the Amazon WorkMail service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package workmailiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/workmail"
)

// WorkMailAPI provides an interface to enable mocking the
// workmail.WorkMail service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon WorkMail.
//	func myFunc(svc workmailiface.WorkMailAPI) bool {
//	    // Make svc.AssociateDelegateToResource request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := workmail.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockWorkMailClient struct {
//	    workmailiface.WorkMailAPI
//	}
//	func (m *mockWorkMailClient) AssociateDelegateToResource(input *workmail.AssociateDelegateToResourceInput) (*workmail.AssociateDelegateToResourceOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockWorkMailClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type WorkMailAPI interface {
	AssociateDelegateToResource(*workmail.AssociateDelegateToResourceInput) (*workmail.AssociateDelegateToResourceOutput, error)
	AssociateDelegateToResourceWithContext(aws.Context, *workmail.AssociateDelegateToResourceInput, ...request.Option) (*workmail.AssociateDelegateToResourceOutput, error)
	AssociateDelegateToResourceRequest(*workmail.AssociateDelegateToResourceInput) (*request.Request, *workmail.AssociateDelegateToResourceOutput)

	AssociateMemberToGroup(*workmail.AssociateMemberToGroupInput) (*workmail.AssociateMemberToGroupOutput, error)
	AssociateMemberToGroupWithContext(aws.Context, *workmail.AssociateMemberToGroupInput, ...request.Option) (*workmail.AssociateMemberToGroupOutput, error)
	AssociateMemberToGroupRequest(*workmail.AssociateMemberToGroupInput) (*request.Request, *workmail.AssociateMemberToGroupOutput)

	AssumeImpersonationRole(*workmail.AssumeImpersonationRoleInput) (*workmail.AssumeImpersonationRoleOutput, error)
	AssumeImpersonationRoleWithContext(aws.Context, *workmail.AssumeImpersonationRoleInput, ...request.Option) (*workmail.AssumeImpersonationRoleOutput, error)
	AssumeImpersonationRoleRequest(*workmail.AssumeImpersonationRoleInput) (*request.Request, *workmail.AssumeImpersonationRoleOutput)

	CancelMailboxExportJob(*workmail.CancelMailboxExportJobInput) (*workmail.CancelMailboxExportJobOutput, error)
	CancelMailboxExportJobWithContext(aws.Context, *workmail.CancelMailboxExportJobInput, ...request.Option) (*workmail.CancelMailboxExportJobOutput, error)
	CancelMailboxExportJobRequest(*workmail.CancelMailboxExportJobInput) (*request.Request, *workmail.CancelMailboxExportJobOutput)

	CreateAlias(*workmail.CreateAliasInput) (*workmail.CreateAliasOutput, error)
	CreateAliasWithContext(aws.Context, *workmail.CreateAliasInput, ...request.Option) (*workmail.CreateAliasOutput, error)
	CreateAliasRequest(*workmail.CreateAliasInput) (*request.Request, *workmail.CreateAliasOutput)

	CreateAvailabilityConfiguration(*workmail.CreateAvailabilityConfigurationInput) (*workmail.CreateAvailabilityConfigurationOutput, error)
	CreateAvailabilityConfigurationWithContext(aws.Context, *workmail.CreateAvailabilityConfigurationInput, ...request.Option) (*workmail.CreateAvailabilityConfigurationOutput, error)
	CreateAvailabilityConfigurationRequest(*workmail.CreateAvailabilityConfigurationInput) (*request.Request, *workmail.CreateAvailabilityConfigurationOutput)

	CreateGroup(*workmail.CreateGroupInput) (*workmail.CreateGroupOutput, error)
	CreateGroupWithContext(aws.Context, *workmail.CreateGroupInput, ...request.Option) (*workmail.CreateGroupOutput, error)
	CreateGroupRequest(*workmail.CreateGroupInput) (*request.Request, *workmail.CreateGroupOutput)

	CreateImpersonationRole(*workmail.CreateImpersonationRoleInput) (*workmail.CreateImpersonationRoleOutput, error)
	CreateImpersonationRoleWithContext(aws.Context, *workmail.CreateImpersonationRoleInput, ...request.Option) (*workmail.CreateImpersonationRoleOutput, error)
	CreateImpersonationRoleRequest(*workmail.CreateImpersonationRoleInput) (*request.Request, *workmail.CreateImpersonationRoleOutput)

	CreateMobileDeviceAccessRule(*workmail.CreateMobileDeviceAccessRuleInput) (*workmail.CreateMobileDeviceAccessRuleOutput, error)
	CreateMobileDeviceAccessRuleWithContext(aws.Context, *workmail.CreateMobileDeviceAccessRuleInput, ...request.Option) (*workmail.CreateMobileDeviceAccessRuleOutput, error)
	CreateMobileDeviceAccessRuleRequest(*workmail.CreateMobileDeviceAccessRuleInput) (*request.Request, *workmail.CreateMobileDeviceAccessRuleOutput)

	CreateOrganization(*workmail.CreateOrganizationInput) (*workmail.CreateOrganizationOutput, error)
	CreateOrganizationWithContext(aws.Context, *workmail.CreateOrganizationInput, ...request.Option) (*workmail.CreateOrganizationOutput, error)
	CreateOrganizationRequest(*workmail.CreateOrganizationInput) (*request.Request, *workmail.CreateOrganizationOutput)

	CreateResource(*workmail.CreateResourceInput) (*workmail.CreateResourceOutput, error)
	CreateResourceWithContext(aws.Context, *workmail.CreateResourceInput, ...request.Option) (*workmail.CreateResourceOutput, error)
	CreateResourceRequest(*workmail.CreateResourceInput) (*request.Request, *workmail.CreateResourceOutput)

	CreateUser(*workmail.CreateUserInput) (*workmail.CreateUserOutput, error)
	CreateUserWithContext(aws.Context, *workmail.CreateUserInput, ...request.Option) (*workmail.CreateUserOutput, error)
	CreateUserRequest(*workmail.CreateUserInput) (*request.Request, *workmail.CreateUserOutput)

	DeleteAccessControlRule(*workmail.DeleteAccessControlRuleInput) (*workmail.DeleteAccessControlRuleOutput, error)
	DeleteAccessControlRuleWithContext(aws.Context, *workmail.DeleteAccessControlRuleInput, ...request.Option) (*workmail.DeleteAccessControlRuleOutput, error)
	DeleteAccessControlRuleRequest(*workmail.DeleteAccessControlRuleInput) (*request.Request, *workmail.DeleteAccessControlRuleOutput)

	DeleteAlias(*workmail.DeleteAliasInput) (*workmail.DeleteAliasOutput, error)
	DeleteAliasWithContext(aws.Context, *workmail.DeleteAliasInput, ...request.Option) (*workmail.DeleteAliasOutput, error)
	DeleteAliasRequest(*workmail.DeleteAliasInput) (*request.Request, *workmail.DeleteAliasOutput)

	DeleteAvailabilityConfiguration(*workmail.DeleteAvailabilityConfigurationInput) (*workmail.DeleteAvailabilityConfigurationOutput, error)
	DeleteAvailabilityConfigurationWithContext(aws.Context, *workmail.DeleteAvailabilityConfigurationInput, ...request.Option) (*workmail.DeleteAvailabilityConfigurationOutput, error)
	DeleteAvailabilityConfigurationRequest(*workmail.DeleteAvailabilityConfigurationInput) (*request.Request, *workmail.DeleteAvailabilityConfigurationOutput)

	DeleteEmailMonitoringConfiguration(*workmail.DeleteEmailMonitoringConfigurationInput) (*workmail.DeleteEmailMonitoringConfigurationOutput, error)
	DeleteEmailMonitoringConfigurationWithContext(aws.Context, *workmail.DeleteEmailMonitoringConfigurationInput, ...request.Option) (*workmail.DeleteEmailMonitoringConfigurationOutput, error)
	DeleteEmailMonitoringConfigurationRequest(*workmail.DeleteEmailMonitoringConfigurationInput) (*request.Request, *workmail.DeleteEmailMonitoringConfigurationOutput)

	DeleteGroup(*workmail.DeleteGroupInput) (*workmail.DeleteGroupOutput, error)
	DeleteGroupWithContext(aws.Context, *workmail.DeleteGroupInput, ...request.Option) (*workmail.DeleteGroupOutput, error)
	DeleteGroupRequest(*workmail.DeleteGroupInput) (*request.Request, *workmail.DeleteGroupOutput)

	DeleteImpersonationRole(*workmail.DeleteImpersonationRoleInput) (*workmail.DeleteImpersonationRoleOutput, error)
	DeleteImpersonationRoleWithContext(aws.Context, *workmail.DeleteImpersonationRoleInput, ...request.Option) (*workmail.DeleteImpersonationRoleOutput, error)
	DeleteImpersonationRoleRequest(*workmail.DeleteImpersonationRoleInput) (*request.Request, *workmail.DeleteImpersonationRoleOutput)

	DeleteMailboxPermissions(*workmail.DeleteMailboxPermissionsInput) (*workmail.DeleteMailboxPermissionsOutput, error)
	DeleteMailboxPermissionsWithContext(aws.Context, *workmail.DeleteMailboxPermissionsInput, ...request.Option) (*workmail.DeleteMailboxPermissionsOutput, error)
	DeleteMailboxPermissionsRequest(*workmail.DeleteMailboxPermissionsInput) (*request.Request, *workmail.DeleteMailboxPermissionsOutput)

	DeleteMobileDeviceAccessOverride(*workmail.DeleteMobileDeviceAccessOverrideInput) (*workmail.DeleteMobileDeviceAccessOverrideOutput, error)
	DeleteMobileDeviceAccessOverrideWithContext(aws.Context, *workmail.DeleteMobileDeviceAccessOverrideInput, ...request.Option) (*workmail.DeleteMobileDeviceAccessOverrideOutput, error)
	DeleteMobileDeviceAccessOverrideRequest(*workmail.DeleteMobileDeviceAccessOverrideInput) (*request.Request, *workmail.DeleteMobileDeviceAccessOverrideOutput)

	DeleteMobileDeviceAccessRule(*workmail.DeleteMobileDeviceAccessRuleInput) (*workmail.DeleteMobileDeviceAccessRuleOutput, error)
	DeleteMobileDeviceAccessRuleWithContext(aws.Context, *workmail.DeleteMobileDeviceAccessRuleInput, ...request.Option) (*workmail.DeleteMobileDeviceAccessRuleOutput, error)
	DeleteMobileDeviceAccessRuleRequest(*workmail.DeleteMobileDeviceAccessRuleInput) (*request.Request, *workmail.DeleteMobileDeviceAccessRuleOutput)

	DeleteOrganization(*workmail.DeleteOrganizationInput) (*workmail.DeleteOrganizationOutput, error)
	DeleteOrganizationWithContext(aws.Context, *workmail.DeleteOrganizationInput, ...request.Option) (*workmail.DeleteOrganizationOutput, error)
	DeleteOrganizationRequest(*workmail.DeleteOrganizationInput) (*request.Request, *workmail.DeleteOrganizationOutput)

	DeleteResource(*workmail.DeleteResourceInput) (*workmail.DeleteResourceOutput, error)
	DeleteResourceWithContext(aws.Context, *workmail.DeleteResourceInput, ...request.Option) (*workmail.DeleteResourceOutput, error)
	DeleteResourceRequest(*workmail.DeleteResourceInput) (*request.Request, *workmail.DeleteResourceOutput)

	DeleteRetentionPolicy(*workmail.DeleteRetentionPolicyInput) (*workmail.DeleteRetentionPolicyOutput, error)
	DeleteRetentionPolicyWithContext(aws.Context, *workmail.DeleteRetentionPolicyInput, ...request.Option) (*workmail.DeleteRetentionPolicyOutput, error)
	DeleteRetentionPolicyRequest(*workmail.DeleteRetentionPolicyInput) (*request.Request, *workmail.DeleteRetentionPolicyOutput)

	DeleteUser(*workmail.DeleteUserInput) (*workmail.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *workmail.DeleteUserInput, ...request.Option) (*workmail.DeleteUserOutput, error)
	DeleteUserRequest(*workmail.DeleteUserInput) (*request.Request, *workmail.DeleteUserOutput)

	DeregisterFromWorkMail(*workmail.DeregisterFromWorkMailInput) (*workmail.DeregisterFromWorkMailOutput, error)
	DeregisterFromWorkMailWithContext(aws.Context, *workmail.DeregisterFromWorkMailInput, ...request.Option) (*workmail.DeregisterFromWorkMailOutput, error)
	DeregisterFromWorkMailRequest(*workmail.DeregisterFromWorkMailInput) (*request.Request, *workmail.DeregisterFromWorkMailOutput)

	DeregisterMailDomain(*workmail.DeregisterMailDomainInput) (*workmail.DeregisterMailDomainOutput, error)
	DeregisterMailDomainWithContext(aws.Context, *workmail.DeregisterMailDomainInput, ...request.Option) (*workmail.DeregisterMailDomainOutput, error)
	DeregisterMailDomainRequest(*workmail.DeregisterMailDomainInput) (*request.Request, *workmail.DeregisterMailDomainOutput)

	DescribeEmailMonitoringConfiguration(*workmail.DescribeEmailMonitoringConfigurationInput) (*workmail.DescribeEmailMonitoringConfigurationOutput, error)
	DescribeEmailMonitoringConfigurationWithContext(aws.Context, *workmail.DescribeEmailMonitoringConfigurationInput, ...request.Option) (*workmail.DescribeEmailMonitoringConfigurationOutput, error)
	DescribeEmailMonitoringConfigurationRequest(*workmail.DescribeEmailMonitoringConfigurationInput) (*request.Request, *workmail.DescribeEmailMonitoringConfigurationOutput)

	DescribeEntity(*workmail.DescribeEntityInput) (*workmail.DescribeEntityOutput, error)
	DescribeEntityWithContext(aws.Context, *workmail.DescribeEntityInput, ...request.Option) (*workmail.DescribeEntityOutput, error)
	DescribeEntityRequest(*workmail.DescribeEntityInput) (*request.Request, *workmail.DescribeEntityOutput)

	DescribeGroup(*workmail.DescribeGroupInput) (*workmail.DescribeGroupOutput, error)
	DescribeGroupWithContext(aws.Context, *workmail.DescribeGroupInput, ...request.Option) (*workmail.DescribeGroupOutput, error)
	DescribeGroupRequest(*workmail.DescribeGroupInput) (*request.Request, *workmail.DescribeGroupOutput)

	DescribeInboundDmarcSettings(*workmail.DescribeInboundDmarcSettingsInput) (*workmail.DescribeInboundDmarcSettingsOutput, error)
	DescribeInboundDmarcSettingsWithContext(aws.Context, *workmail.DescribeInboundDmarcSettingsInput, ...request.Option) (*workmail.DescribeInboundDmarcSettingsOutput, error)
	DescribeInboundDmarcSettingsRequest(*workmail.DescribeInboundDmarcSettingsInput) (*request.Request, *workmail.DescribeInboundDmarcSettingsOutput)

	DescribeMailboxExportJob(*workmail.DescribeMailboxExportJobInput) (*workmail.DescribeMailboxExportJobOutput, error)
	DescribeMailboxExportJobWithContext(aws.Context, *workmail.DescribeMailboxExportJobInput, ...request.Option) (*workmail.DescribeMailboxExportJobOutput, error)
	DescribeMailboxExportJobRequest(*workmail.DescribeMailboxExportJobInput) (*request.Request, *workmail.DescribeMailboxExportJobOutput)

	DescribeOrganization(*workmail.DescribeOrganizationInput) (*workmail.DescribeOrganizationOutput, error)
	DescribeOrganizationWithContext(aws.Context, *workmail.DescribeOrganizationInput, ...request.Option) (*workmail.DescribeOrganizationOutput, error)
	DescribeOrganizationRequest(*workmail.DescribeOrganizationInput) (*request.Request, *workmail.DescribeOrganizationOutput)

	DescribeResource(*workmail.DescribeResourceInput) (*workmail.DescribeResourceOutput, error)
	DescribeResourceWithContext(aws.Context, *workmail.DescribeResourceInput, ...request.Option) (*workmail.DescribeResourceOutput, error)
	DescribeResourceRequest(*workmail.DescribeResourceInput) (*request.Request, *workmail.DescribeResourceOutput)

	DescribeUser(*workmail.DescribeUserInput) (*workmail.DescribeUserOutput, error)
	DescribeUserWithContext(aws.Context, *workmail.DescribeUserInput, ...request.Option) (*workmail.DescribeUserOutput, error)
	DescribeUserRequest(*workmail.DescribeUserInput) (*request.Request, *workmail.DescribeUserOutput)

	DisassociateDelegateFromResource(*workmail.DisassociateDelegateFromResourceInput) (*workmail.DisassociateDelegateFromResourceOutput, error)
	DisassociateDelegateFromResourceWithContext(aws.Context, *workmail.DisassociateDelegateFromResourceInput, ...request.Option) (*workmail.DisassociateDelegateFromResourceOutput, error)
	DisassociateDelegateFromResourceRequest(*workmail.DisassociateDelegateFromResourceInput) (*request.Request, *workmail.DisassociateDelegateFromResourceOutput)

	DisassociateMemberFromGroup(*workmail.DisassociateMemberFromGroupInput) (*workmail.DisassociateMemberFromGroupOutput, error)
	DisassociateMemberFromGroupWithContext(aws.Context, *workmail.DisassociateMemberFromGroupInput, ...request.Option) (*workmail.DisassociateMemberFromGroupOutput, error)
	DisassociateMemberFromGroupRequest(*workmail.DisassociateMemberFromGroupInput) (*request.Request, *workmail.DisassociateMemberFromGroupOutput)

	GetAccessControlEffect(*workmail.GetAccessControlEffectInput) (*workmail.GetAccessControlEffectOutput, error)
	GetAccessControlEffectWithContext(aws.Context, *workmail.GetAccessControlEffectInput, ...request.Option) (*workmail.GetAccessControlEffectOutput, error)
	GetAccessControlEffectRequest(*workmail.GetAccessControlEffectInput) (*request.Request, *workmail.GetAccessControlEffectOutput)

	GetDefaultRetentionPolicy(*workmail.GetDefaultRetentionPolicyInput) (*workmail.GetDefaultRetentionPolicyOutput, error)
	GetDefaultRetentionPolicyWithContext(aws.Context, *workmail.GetDefaultRetentionPolicyInput, ...request.Option) (*workmail.GetDefaultRetentionPolicyOutput, error)
	GetDefaultRetentionPolicyRequest(*workmail.GetDefaultRetentionPolicyInput) (*request.Request, *workmail.GetDefaultRetentionPolicyOutput)

	GetImpersonationRole(*workmail.GetImpersonationRoleInput) (*workmail.GetImpersonationRoleOutput, error)
	GetImpersonationRoleWithContext(aws.Context, *workmail.GetImpersonationRoleInput, ...request.Option) (*workmail.GetImpersonationRoleOutput, error)
	GetImpersonationRoleRequest(*workmail.GetImpersonationRoleInput) (*request.Request, *workmail.GetImpersonationRoleOutput)

	GetImpersonationRoleEffect(*workmail.GetImpersonationRoleEffectInput) (*workmail.GetImpersonationRoleEffectOutput, error)
	GetImpersonationRoleEffectWithContext(aws.Context, *workmail.GetImpersonationRoleEffectInput, ...request.Option) (*workmail.GetImpersonationRoleEffectOutput, error)
	GetImpersonationRoleEffectRequest(*workmail.GetImpersonationRoleEffectInput) (*request.Request, *workmail.GetImpersonationRoleEffectOutput)

	GetMailDomain(*workmail.GetMailDomainInput) (*workmail.GetMailDomainOutput, error)
	GetMailDomainWithContext(aws.Context, *workmail.GetMailDomainInput, ...request.Option) (*workmail.GetMailDomainOutput, error)
	GetMailDomainRequest(*workmail.GetMailDomainInput) (*request.Request, *workmail.GetMailDomainOutput)

	GetMailboxDetails(*workmail.GetMailboxDetailsInput) (*workmail.GetMailboxDetailsOutput, error)
	GetMailboxDetailsWithContext(aws.Context, *workmail.GetMailboxDetailsInput, ...request.Option) (*workmail.GetMailboxDetailsOutput, error)
	GetMailboxDetailsRequest(*workmail.GetMailboxDetailsInput) (*request.Request, *workmail.GetMailboxDetailsOutput)

	GetMobileDeviceAccessEffect(*workmail.GetMobileDeviceAccessEffectInput) (*workmail.GetMobileDeviceAccessEffectOutput, error)
	GetMobileDeviceAccessEffectWithContext(aws.Context, *workmail.GetMobileDeviceAccessEffectInput, ...request.Option) (*workmail.GetMobileDeviceAccessEffectOutput, error)
	GetMobileDeviceAccessEffectRequest(*workmail.GetMobileDeviceAccessEffectInput) (*request.Request, *workmail.GetMobileDeviceAccessEffectOutput)

	GetMobileDeviceAccessOverride(*workmail.GetMobileDeviceAccessOverrideInput) (*workmail.GetMobileDeviceAccessOverrideOutput, error)
	GetMobileDeviceAccessOverrideWithContext(aws.Context, *workmail.GetMobileDeviceAccessOverrideInput, ...request.Option) (*workmail.GetMobileDeviceAccessOverrideOutput, error)
	GetMobileDeviceAccessOverrideRequest(*workmail.GetMobileDeviceAccessOverrideInput) (*request.Request, *workmail.GetMobileDeviceAccessOverrideOutput)

	ListAccessControlRules(*workmail.ListAccessControlRulesInput) (*workmail.ListAccessControlRulesOutput, error)
	ListAccessControlRulesWithContext(aws.Context, *workmail.ListAccessControlRulesInput, ...request.Option) (*workmail.ListAccessControlRulesOutput, error)
	ListAccessControlRulesRequest(*workmail.ListAccessControlRulesInput) (*request.Request, *workmail.ListAccessControlRulesOutput)

	ListAliases(*workmail.ListAliasesInput) (*workmail.ListAliasesOutput, error)
	ListAliasesWithContext(aws.Context, *workmail.ListAliasesInput, ...request.Option) (*workmail.ListAliasesOutput, error)
	ListAliasesRequest(*workmail.ListAliasesInput) (*request.Request, *workmail.ListAliasesOutput)

	ListAliasesPages(*workmail.ListAliasesInput, func(*workmail.ListAliasesOutput, bool) bool) error
	ListAliasesPagesWithContext(aws.Context, *workmail.ListAliasesInput, func(*workmail.ListAliasesOutput, bool) bool, ...request.Option) error

	ListAvailabilityConfigurations(*workmail.ListAvailabilityConfigurationsInput) (*workmail.ListAvailabilityConfigurationsOutput, error)
	ListAvailabilityConfigurationsWithContext(aws.Context, *workmail.ListAvailabilityConfigurationsInput, ...request.Option) (*workmail.ListAvailabilityConfigurationsOutput, error)
	ListAvailabilityConfigurationsRequest(*workmail.ListAvailabilityConfigurationsInput) (*request.Request, *workmail.ListAvailabilityConfigurationsOutput)

	ListAvailabilityConfigurationsPages(*workmail.ListAvailabilityConfigurationsInput, func(*workmail.ListAvailabilityConfigurationsOutput, bool) bool) error
	ListAvailabilityConfigurationsPagesWithContext(aws.Context, *workmail.ListAvailabilityConfigurationsInput, func(*workmail.ListAvailabilityConfigurationsOutput, bool) bool, ...request.Option) error

	ListGroupMembers(*workmail.ListGroupMembersInput) (*workmail.ListGroupMembersOutput, error)
	ListGroupMembersWithContext(aws.Context, *workmail.ListGroupMembersInput, ...request.Option) (*workmail.ListGroupMembersOutput, error)
	ListGroupMembersRequest(*workmail.ListGroupMembersInput) (*request.Request, *workmail.ListGroupMembersOutput)

	ListGroupMembersPages(*workmail.ListGroupMembersInput, func(*workmail.ListGroupMembersOutput, bool) bool) error
	ListGroupMembersPagesWithContext(aws.Context, *workmail.ListGroupMembersInput, func(*workmail.ListGroupMembersOutput, bool) bool, ...request.Option) error

	ListGroups(*workmail.ListGroupsInput) (*workmail.ListGroupsOutput, error)
	ListGroupsWithContext(aws.Context, *workmail.ListGroupsInput, ...request.Option) (*workmail.ListGroupsOutput, error)
	ListGroupsRequest(*workmail.ListGroupsInput) (*request.Request, *workmail.ListGroupsOutput)

	ListGroupsPages(*workmail.ListGroupsInput, func(*workmail.ListGroupsOutput, bool) bool) error
	ListGroupsPagesWithContext(aws.Context, *workmail.ListGroupsInput, func(*workmail.ListGroupsOutput, bool) bool, ...request.Option) error

	ListGroupsForEntity(*workmail.ListGroupsForEntityInput) (*workmail.ListGroupsForEntityOutput, error)
	ListGroupsForEntityWithContext(aws.Context, *workmail.ListGroupsForEntityInput, ...request.Option) (*workmail.ListGroupsForEntityOutput, error)
	ListGroupsForEntityRequest(*workmail.ListGroupsForEntityInput) (*request.Request, *workmail.ListGroupsForEntityOutput)

	ListGroupsForEntityPages(*workmail.ListGroupsForEntityInput, func(*workmail.ListGroupsForEntityOutput, bool) bool) error
	ListGroupsForEntityPagesWithContext(aws.Context, *workmail.ListGroupsForEntityInput, func(*workmail.ListGroupsForEntityOutput, bool) bool, ...request.Option) error

	ListImpersonationRoles(*workmail.ListImpersonationRolesInput) (*workmail.ListImpersonationRolesOutput, error)
	ListImpersonationRolesWithContext(aws.Context, *workmail.ListImpersonationRolesInput, ...request.Option) (*workmail.ListImpersonationRolesOutput, error)
	ListImpersonationRolesRequest(*workmail.ListImpersonationRolesInput) (*request.Request, *workmail.ListImpersonationRolesOutput)

	ListImpersonationRolesPages(*workmail.ListImpersonationRolesInput, func(*workmail.ListImpersonationRolesOutput, bool) bool) error
	ListImpersonationRolesPagesWithContext(aws.Context, *workmail.ListImpersonationRolesInput, func(*workmail.ListImpersonationRolesOutput, bool) bool, ...request.Option) error

	ListMailDomains(*workmail.ListMailDomainsInput) (*workmail.ListMailDomainsOutput, error)
	ListMailDomainsWithContext(aws.Context, *workmail.ListMailDomainsInput, ...request.Option) (*workmail.ListMailDomainsOutput, error)
	ListMailDomainsRequest(*workmail.ListMailDomainsInput) (*request.Request, *workmail.ListMailDomainsOutput)

	ListMailDomainsPages(*workmail.ListMailDomainsInput, func(*workmail.ListMailDomainsOutput, bool) bool) error
	ListMailDomainsPagesWithContext(aws.Context, *workmail.ListMailDomainsInput, func(*workmail.ListMailDomainsOutput, bool) bool, ...request.Option) error

	ListMailboxExportJobs(*workmail.ListMailboxExportJobsInput) (*workmail.ListMailboxExportJobsOutput, error)
	ListMailboxExportJobsWithContext(aws.Context, *workmail.ListMailboxExportJobsInput, ...request.Option) (*workmail.ListMailboxExportJobsOutput, error)
	ListMailboxExportJobsRequest(*workmail.ListMailboxExportJobsInput) (*request.Request, *workmail.ListMailboxExportJobsOutput)

	ListMailboxExportJobsPages(*workmail.ListMailboxExportJobsInput, func(*workmail.ListMailboxExportJobsOutput, bool) bool) error
	ListMailboxExportJobsPagesWithContext(aws.Context, *workmail.ListMailboxExportJobsInput, func(*workmail.ListMailboxExportJobsOutput, bool) bool, ...request.Option) error

	ListMailboxPermissions(*workmail.ListMailboxPermissionsInput) (*workmail.ListMailboxPermissionsOutput, error)
	ListMailboxPermissionsWithContext(aws.Context, *workmail.ListMailboxPermissionsInput, ...request.Option) (*workmail.ListMailboxPermissionsOutput, error)
	ListMailboxPermissionsRequest(*workmail.ListMailboxPermissionsInput) (*request.Request, *workmail.ListMailboxPermissionsOutput)

	ListMailboxPermissionsPages(*workmail.ListMailboxPermissionsInput, func(*workmail.ListMailboxPermissionsOutput, bool) bool) error
	ListMailboxPermissionsPagesWithContext(aws.Context, *workmail.ListMailboxPermissionsInput, func(*workmail.ListMailboxPermissionsOutput, bool) bool, ...request.Option) error

	ListMobileDeviceAccessOverrides(*workmail.ListMobileDeviceAccessOverridesInput) (*workmail.ListMobileDeviceAccessOverridesOutput, error)
	ListMobileDeviceAccessOverridesWithContext(aws.Context, *workmail.ListMobileDeviceAccessOverridesInput, ...request.Option) (*workmail.ListMobileDeviceAccessOverridesOutput, error)
	ListMobileDeviceAccessOverridesRequest(*workmail.ListMobileDeviceAccessOverridesInput) (*request.Request, *workmail.ListMobileDeviceAccessOverridesOutput)

	ListMobileDeviceAccessOverridesPages(*workmail.ListMobileDeviceAccessOverridesInput, func(*workmail.ListMobileDeviceAccessOverridesOutput, bool) bool) error
	ListMobileDeviceAccessOverridesPagesWithContext(aws.Context, *workmail.ListMobileDeviceAccessOverridesInput, func(*workmail.ListMobileDeviceAccessOverridesOutput, bool) bool, ...request.Option) error

	ListMobileDeviceAccessRules(*workmail.ListMobileDeviceAccessRulesInput) (*workmail.ListMobileDeviceAccessRulesOutput, error)
	ListMobileDeviceAccessRulesWithContext(aws.Context, *workmail.ListMobileDeviceAccessRulesInput, ...request.Option) (*workmail.ListMobileDeviceAccessRulesOutput, error)
	ListMobileDeviceAccessRulesRequest(*workmail.ListMobileDeviceAccessRulesInput) (*request.Request, *workmail.ListMobileDeviceAccessRulesOutput)

	ListOrganizations(*workmail.ListOrganizationsInput) (*workmail.ListOrganizationsOutput, error)
	ListOrganizationsWithContext(aws.Context, *workmail.ListOrganizationsInput, ...request.Option) (*workmail.ListOrganizationsOutput, error)
	ListOrganizationsRequest(*workmail.ListOrganizationsInput) (*request.Request, *workmail.ListOrganizationsOutput)

	ListOrganizationsPages(*workmail.ListOrganizationsInput, func(*workmail.ListOrganizationsOutput, bool) bool) error
	ListOrganizationsPagesWithContext(aws.Context, *workmail.ListOrganizationsInput, func(*workmail.ListOrganizationsOutput, bool) bool, ...request.Option) error

	ListResourceDelegates(*workmail.ListResourceDelegatesInput) (*workmail.ListResourceDelegatesOutput, error)
	ListResourceDelegatesWithContext(aws.Context, *workmail.ListResourceDelegatesInput, ...request.Option) (*workmail.ListResourceDelegatesOutput, error)
	ListResourceDelegatesRequest(*workmail.ListResourceDelegatesInput) (*request.Request, *workmail.ListResourceDelegatesOutput)

	ListResourceDelegatesPages(*workmail.ListResourceDelegatesInput, func(*workmail.ListResourceDelegatesOutput, bool) bool) error
	ListResourceDelegatesPagesWithContext(aws.Context, *workmail.ListResourceDelegatesInput, func(*workmail.ListResourceDelegatesOutput, bool) bool, ...request.Option) error

	ListResources(*workmail.ListResourcesInput) (*workmail.ListResourcesOutput, error)
	ListResourcesWithContext(aws.Context, *workmail.ListResourcesInput, ...request.Option) (*workmail.ListResourcesOutput, error)
	ListResourcesRequest(*workmail.ListResourcesInput) (*request.Request, *workmail.ListResourcesOutput)

	ListResourcesPages(*workmail.ListResourcesInput, func(*workmail.ListResourcesOutput, bool) bool) error
	ListResourcesPagesWithContext(aws.Context, *workmail.ListResourcesInput, func(*workmail.ListResourcesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*workmail.ListTagsForResourceInput) (*workmail.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *workmail.ListTagsForResourceInput, ...request.Option) (*workmail.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*workmail.ListTagsForResourceInput) (*request.Request, *workmail.ListTagsForResourceOutput)

	ListUsers(*workmail.ListUsersInput) (*workmail.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *workmail.ListUsersInput, ...request.Option) (*workmail.ListUsersOutput, error)
	ListUsersRequest(*workmail.ListUsersInput) (*request.Request, *workmail.ListUsersOutput)

	ListUsersPages(*workmail.ListUsersInput, func(*workmail.ListUsersOutput, bool) bool) error
	ListUsersPagesWithContext(aws.Context, *workmail.ListUsersInput, func(*workmail.ListUsersOutput, bool) bool, ...request.Option) error

	PutAccessControlRule(*workmail.PutAccessControlRuleInput) (*workmail.PutAccessControlRuleOutput, error)
	PutAccessControlRuleWithContext(aws.Context, *workmail.PutAccessControlRuleInput, ...request.Option) (*workmail.PutAccessControlRuleOutput, error)
	PutAccessControlRuleRequest(*workmail.PutAccessControlRuleInput) (*request.Request, *workmail.PutAccessControlRuleOutput)

	PutEmailMonitoringConfiguration(*workmail.PutEmailMonitoringConfigurationInput) (*workmail.PutEmailMonitoringConfigurationOutput, error)
	PutEmailMonitoringConfigurationWithContext(aws.Context, *workmail.PutEmailMonitoringConfigurationInput, ...request.Option) (*workmail.PutEmailMonitoringConfigurationOutput, error)
	PutEmailMonitoringConfigurationRequest(*workmail.PutEmailMonitoringConfigurationInput) (*request.Request, *workmail.PutEmailMonitoringConfigurationOutput)

	PutInboundDmarcSettings(*workmail.PutInboundDmarcSettingsInput) (*workmail.PutInboundDmarcSettingsOutput, error)
	PutInboundDmarcSettingsWithContext(aws.Context, *workmail.PutInboundDmarcSettingsInput, ...request.Option) (*workmail.PutInboundDmarcSettingsOutput, error)
	PutInboundDmarcSettingsRequest(*workmail.PutInboundDmarcSettingsInput) (*request.Request, *workmail.PutInboundDmarcSettingsOutput)

	PutMailboxPermissions(*workmail.PutMailboxPermissionsInput) (*workmail.PutMailboxPermissionsOutput, error)
	PutMailboxPermissionsWithContext(aws.Context, *workmail.PutMailboxPermissionsInput, ...request.Option) (*workmail.PutMailboxPermissionsOutput, error)
	PutMailboxPermissionsRequest(*workmail.PutMailboxPermissionsInput) (*request.Request, *workmail.PutMailboxPermissionsOutput)

	PutMobileDeviceAccessOverride(*workmail.PutMobileDeviceAccessOverrideInput) (*workmail.PutMobileDeviceAccessOverrideOutput, error)
	PutMobileDeviceAccessOverrideWithContext(aws.Context, *workmail.PutMobileDeviceAccessOverrideInput, ...request.Option) (*workmail.PutMobileDeviceAccessOverrideOutput, error)
	PutMobileDeviceAccessOverrideRequest(*workmail.PutMobileDeviceAccessOverrideInput) (*request.Request, *workmail.PutMobileDeviceAccessOverrideOutput)

	PutRetentionPolicy(*workmail.PutRetentionPolicyInput) (*workmail.PutRetentionPolicyOutput, error)
	PutRetentionPolicyWithContext(aws.Context, *workmail.PutRetentionPolicyInput, ...request.Option) (*workmail.PutRetentionPolicyOutput, error)
	PutRetentionPolicyRequest(*workmail.PutRetentionPolicyInput) (*request.Request, *workmail.PutRetentionPolicyOutput)

	RegisterMailDomain(*workmail.RegisterMailDomainInput) (*workmail.RegisterMailDomainOutput, error)
	RegisterMailDomainWithContext(aws.Context, *workmail.RegisterMailDomainInput, ...request.Option) (*workmail.RegisterMailDomainOutput, error)
	RegisterMailDomainRequest(*workmail.RegisterMailDomainInput) (*request.Request, *workmail.RegisterMailDomainOutput)

	RegisterToWorkMail(*workmail.RegisterToWorkMailInput) (*workmail.RegisterToWorkMailOutput, error)
	RegisterToWorkMailWithContext(aws.Context, *workmail.RegisterToWorkMailInput, ...request.Option) (*workmail.RegisterToWorkMailOutput, error)
	RegisterToWorkMailRequest(*workmail.RegisterToWorkMailInput) (*request.Request, *workmail.RegisterToWorkMailOutput)

	ResetPassword(*workmail.ResetPasswordInput) (*workmail.ResetPasswordOutput, error)
	ResetPasswordWithContext(aws.Context, *workmail.ResetPasswordInput, ...request.Option) (*workmail.ResetPasswordOutput, error)
	ResetPasswordRequest(*workmail.ResetPasswordInput) (*request.Request, *workmail.ResetPasswordOutput)

	StartMailboxExportJob(*workmail.StartMailboxExportJobInput) (*workmail.StartMailboxExportJobOutput, error)
	StartMailboxExportJobWithContext(aws.Context, *workmail.StartMailboxExportJobInput, ...request.Option) (*workmail.StartMailboxExportJobOutput, error)
	StartMailboxExportJobRequest(*workmail.StartMailboxExportJobInput) (*request.Request, *workmail.StartMailboxExportJobOutput)

	TagResource(*workmail.TagResourceInput) (*workmail.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *workmail.TagResourceInput, ...request.Option) (*workmail.TagResourceOutput, error)
	TagResourceRequest(*workmail.TagResourceInput) (*request.Request, *workmail.TagResourceOutput)

	TestAvailabilityConfiguration(*workmail.TestAvailabilityConfigurationInput) (*workmail.TestAvailabilityConfigurationOutput, error)
	TestAvailabilityConfigurationWithContext(aws.Context, *workmail.TestAvailabilityConfigurationInput, ...request.Option) (*workmail.TestAvailabilityConfigurationOutput, error)
	TestAvailabilityConfigurationRequest(*workmail.TestAvailabilityConfigurationInput) (*request.Request, *workmail.TestAvailabilityConfigurationOutput)

	UntagResource(*workmail.UntagResourceInput) (*workmail.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *workmail.UntagResourceInput, ...request.Option) (*workmail.UntagResourceOutput, error)
	UntagResourceRequest(*workmail.UntagResourceInput) (*request.Request, *workmail.UntagResourceOutput)

	UpdateAvailabilityConfiguration(*workmail.UpdateAvailabilityConfigurationInput) (*workmail.UpdateAvailabilityConfigurationOutput, error)
	UpdateAvailabilityConfigurationWithContext(aws.Context, *workmail.UpdateAvailabilityConfigurationInput, ...request.Option) (*workmail.UpdateAvailabilityConfigurationOutput, error)
	UpdateAvailabilityConfigurationRequest(*workmail.UpdateAvailabilityConfigurationInput) (*request.Request, *workmail.UpdateAvailabilityConfigurationOutput)

	UpdateDefaultMailDomain(*workmail.UpdateDefaultMailDomainInput) (*workmail.UpdateDefaultMailDomainOutput, error)
	UpdateDefaultMailDomainWithContext(aws.Context, *workmail.UpdateDefaultMailDomainInput, ...request.Option) (*workmail.UpdateDefaultMailDomainOutput, error)
	UpdateDefaultMailDomainRequest(*workmail.UpdateDefaultMailDomainInput) (*request.Request, *workmail.UpdateDefaultMailDomainOutput)

	UpdateGroup(*workmail.UpdateGroupInput) (*workmail.UpdateGroupOutput, error)
	UpdateGroupWithContext(aws.Context, *workmail.UpdateGroupInput, ...request.Option) (*workmail.UpdateGroupOutput, error)
	UpdateGroupRequest(*workmail.UpdateGroupInput) (*request.Request, *workmail.UpdateGroupOutput)

	UpdateImpersonationRole(*workmail.UpdateImpersonationRoleInput) (*workmail.UpdateImpersonationRoleOutput, error)
	UpdateImpersonationRoleWithContext(aws.Context, *workmail.UpdateImpersonationRoleInput, ...request.Option) (*workmail.UpdateImpersonationRoleOutput, error)
	UpdateImpersonationRoleRequest(*workmail.UpdateImpersonationRoleInput) (*request.Request, *workmail.UpdateImpersonationRoleOutput)

	UpdateMailboxQuota(*workmail.UpdateMailboxQuotaInput) (*workmail.UpdateMailboxQuotaOutput, error)
	UpdateMailboxQuotaWithContext(aws.Context, *workmail.UpdateMailboxQuotaInput, ...request.Option) (*workmail.UpdateMailboxQuotaOutput, error)
	UpdateMailboxQuotaRequest(*workmail.UpdateMailboxQuotaInput) (*request.Request, *workmail.UpdateMailboxQuotaOutput)

	UpdateMobileDeviceAccessRule(*workmail.UpdateMobileDeviceAccessRuleInput) (*workmail.UpdateMobileDeviceAccessRuleOutput, error)
	UpdateMobileDeviceAccessRuleWithContext(aws.Context, *workmail.UpdateMobileDeviceAccessRuleInput, ...request.Option) (*workmail.UpdateMobileDeviceAccessRuleOutput, error)
	UpdateMobileDeviceAccessRuleRequest(*workmail.UpdateMobileDeviceAccessRuleInput) (*request.Request, *workmail.UpdateMobileDeviceAccessRuleOutput)

	UpdatePrimaryEmailAddress(*workmail.UpdatePrimaryEmailAddressInput) (*workmail.UpdatePrimaryEmailAddressOutput, error)
	UpdatePrimaryEmailAddressWithContext(aws.Context, *workmail.UpdatePrimaryEmailAddressInput, ...request.Option) (*workmail.UpdatePrimaryEmailAddressOutput, error)
	UpdatePrimaryEmailAddressRequest(*workmail.UpdatePrimaryEmailAddressInput) (*request.Request, *workmail.UpdatePrimaryEmailAddressOutput)

	UpdateResource(*workmail.UpdateResourceInput) (*workmail.UpdateResourceOutput, error)
	UpdateResourceWithContext(aws.Context, *workmail.UpdateResourceInput, ...request.Option) (*workmail.UpdateResourceOutput, error)
	UpdateResourceRequest(*workmail.UpdateResourceInput) (*request.Request, *workmail.UpdateResourceOutput)

	UpdateUser(*workmail.UpdateUserInput) (*workmail.UpdateUserOutput, error)
	UpdateUserWithContext(aws.Context, *workmail.UpdateUserInput, ...request.Option) (*workmail.UpdateUserOutput, error)
	UpdateUserRequest(*workmail.UpdateUserInput) (*request.Request, *workmail.UpdateUserOutput)
}

var _ WorkMailAPI = (*workmail.WorkMail)(nil)
