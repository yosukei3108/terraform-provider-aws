// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package route53recoverycontrolconfig

import (
	"context"

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
			Factory:  ResourceCluster,
			TypeName: "aws_route53recoverycontrolconfig_cluster",
		},
		{
			Factory:  ResourceControlPanel,
			TypeName: "aws_route53recoverycontrolconfig_control_panel",
		},
		{
			Factory:  ResourceRoutingControl,
			TypeName: "aws_route53recoverycontrolconfig_routing_control",
		},
		{
			Factory:  ResourceSafetyRule,
			TypeName: "aws_route53recoverycontrolconfig_safety_rule",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Route53RecoveryControlConfig
}

func (p *servicePackage) Configure(ctx context.Context, config map[string]any) {
	p.config = config
}

var ServicePackage = &servicePackage{}
