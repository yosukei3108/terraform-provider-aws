// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package securityhub

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	securityhub_sdkv1 "github.com/aws/aws-sdk-go/service/securityhub"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct {
	config map[string]any
}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccount,
			TypeName: "aws_securityhub_account",
		},
		{
			Factory:  ResourceActionTarget,
			TypeName: "aws_securityhub_action_target",
		},
		{
			Factory:  ResourceFindingAggregator,
			TypeName: "aws_securityhub_finding_aggregator",
		},
		{
			Factory:  ResourceInsight,
			TypeName: "aws_securityhub_insight",
		},
		{
			Factory:  ResourceInviteAccepter,
			TypeName: "aws_securityhub_invite_accepter",
		},
		{
			Factory:  ResourceMember,
			TypeName: "aws_securityhub_member",
		},
		{
			Factory:  ResourceOrganizationAdminAccount,
			TypeName: "aws_securityhub_organization_admin_account",
		},
		{
			Factory:  ResourceOrganizationConfiguration,
			TypeName: "aws_securityhub_organization_configuration",
		},
		{
			Factory:  ResourceProductSubscription,
			TypeName: "aws_securityhub_product_subscription",
		},
		{
			Factory:  ResourceStandardsControl,
			TypeName: "aws_securityhub_standards_control",
		},
		{
			Factory:  ResourceStandardsSubscription,
			TypeName: "aws_securityhub_standards_subscription",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SecurityHub
}

func (p *servicePackage) Configure(ctx context.Context, config map[string]any) {
	p.config = config
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context) (*securityhub_sdkv1.SecurityHub, error) {
	sess := p.config["session"].(*session_sdkv1.Session)

	return securityhub_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
