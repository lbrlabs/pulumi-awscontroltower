// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awscontroltower

import (
	"context"
	"reflect"

	"errors"
	"github.com/lbrlabs/pulumi-awscontroltower/sdk/go/awscontroltower/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type ControlTowerAwsAccount struct {
	pulumi.CustomResourceState

	// ID of the AWS account.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// If enabled, this will close the AWS account on resource deletion, beginning the 90-day suspension period. Otherwise, the
	// account will just be unenrolled from Control Tower.
	CloseAccountOnDelete pulumi.BoolPtrOutput `pulumi:"closeAccountOnDelete"`
	// Root email of the account.
	Email pulumi.StringOutput `pulumi:"email"`
	// Name of the account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of the Organizational Unit under which the account resides.
	OrganizationalUnit pulumi.StringOutput `pulumi:"organizationalUnit"`
	// ID of the Organizational Unit to which the account should be moved when the resource is deleted. If no value is
	// provided, the account will not be moved.
	OrganizationalUnitIdOnDelete pulumi.StringPtrOutput `pulumi:"organizationalUnitIdOnDelete"`
	// Name of the path identifier of the product. This value is optional if the product has a default path, and required if
	// the product has more than one path. To list the paths for a product, use ListLaunchPaths.
	PathId pulumi.StringOutput `pulumi:"pathId"`
	// Name of the service catalog product that is provisioned. Defaults to a slugified version of the account name.
	ProvisionedProductName pulumi.StringOutput `pulumi:"provisionedProductName"`
	// Assigned SSO user settings.
	Sso ControlTowerAwsAccountSsoOutput `pulumi:"sso"`
	// Key-value map of resource tags for the account.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewControlTowerAwsAccount registers a new resource with the given unique name, arguments, and options.
func NewControlTowerAwsAccount(ctx *pulumi.Context,
	name string, args *ControlTowerAwsAccountArgs, opts ...pulumi.ResourceOption) (*ControlTowerAwsAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.OrganizationalUnit == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationalUnit'")
	}
	if args.Sso == nil {
		return nil, errors.New("invalid value for required argument 'Sso'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ControlTowerAwsAccount
	err := ctx.RegisterResource("awscontroltower:index/controlTowerAwsAccount:controlTowerAwsAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetControlTowerAwsAccount gets an existing ControlTowerAwsAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetControlTowerAwsAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ControlTowerAwsAccountState, opts ...pulumi.ResourceOption) (*ControlTowerAwsAccount, error) {
	var resource ControlTowerAwsAccount
	err := ctx.ReadResource("awscontroltower:index/controlTowerAwsAccount:controlTowerAwsAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ControlTowerAwsAccount resources.
type controlTowerAwsAccountState struct {
	// ID of the AWS account.
	AccountId *string `pulumi:"accountId"`
	// If enabled, this will close the AWS account on resource deletion, beginning the 90-day suspension period. Otherwise, the
	// account will just be unenrolled from Control Tower.
	CloseAccountOnDelete *bool `pulumi:"closeAccountOnDelete"`
	// Root email of the account.
	Email *string `pulumi:"email"`
	// Name of the account.
	Name *string `pulumi:"name"`
	// Name of the Organizational Unit under which the account resides.
	OrganizationalUnit *string `pulumi:"organizationalUnit"`
	// ID of the Organizational Unit to which the account should be moved when the resource is deleted. If no value is
	// provided, the account will not be moved.
	OrganizationalUnitIdOnDelete *string `pulumi:"organizationalUnitIdOnDelete"`
	// Name of the path identifier of the product. This value is optional if the product has a default path, and required if
	// the product has more than one path. To list the paths for a product, use ListLaunchPaths.
	PathId *string `pulumi:"pathId"`
	// Name of the service catalog product that is provisioned. Defaults to a slugified version of the account name.
	ProvisionedProductName *string `pulumi:"provisionedProductName"`
	// Assigned SSO user settings.
	Sso *ControlTowerAwsAccountSso `pulumi:"sso"`
	// Key-value map of resource tags for the account.
	Tags map[string]string `pulumi:"tags"`
}

type ControlTowerAwsAccountState struct {
	// ID of the AWS account.
	AccountId pulumi.StringPtrInput
	// If enabled, this will close the AWS account on resource deletion, beginning the 90-day suspension period. Otherwise, the
	// account will just be unenrolled from Control Tower.
	CloseAccountOnDelete pulumi.BoolPtrInput
	// Root email of the account.
	Email pulumi.StringPtrInput
	// Name of the account.
	Name pulumi.StringPtrInput
	// Name of the Organizational Unit under which the account resides.
	OrganizationalUnit pulumi.StringPtrInput
	// ID of the Organizational Unit to which the account should be moved when the resource is deleted. If no value is
	// provided, the account will not be moved.
	OrganizationalUnitIdOnDelete pulumi.StringPtrInput
	// Name of the path identifier of the product. This value is optional if the product has a default path, and required if
	// the product has more than one path. To list the paths for a product, use ListLaunchPaths.
	PathId pulumi.StringPtrInput
	// Name of the service catalog product that is provisioned. Defaults to a slugified version of the account name.
	ProvisionedProductName pulumi.StringPtrInput
	// Assigned SSO user settings.
	Sso ControlTowerAwsAccountSsoPtrInput
	// Key-value map of resource tags for the account.
	Tags pulumi.StringMapInput
}

func (ControlTowerAwsAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*controlTowerAwsAccountState)(nil)).Elem()
}

type controlTowerAwsAccountArgs struct {
	// If enabled, this will close the AWS account on resource deletion, beginning the 90-day suspension period. Otherwise, the
	// account will just be unenrolled from Control Tower.
	CloseAccountOnDelete *bool `pulumi:"closeAccountOnDelete"`
	// Root email of the account.
	Email string `pulumi:"email"`
	// Name of the account.
	Name *string `pulumi:"name"`
	// Name of the Organizational Unit under which the account resides.
	OrganizationalUnit string `pulumi:"organizationalUnit"`
	// ID of the Organizational Unit to which the account should be moved when the resource is deleted. If no value is
	// provided, the account will not be moved.
	OrganizationalUnitIdOnDelete *string `pulumi:"organizationalUnitIdOnDelete"`
	// Name of the path identifier of the product. This value is optional if the product has a default path, and required if
	// the product has more than one path. To list the paths for a product, use ListLaunchPaths.
	PathId *string `pulumi:"pathId"`
	// Name of the service catalog product that is provisioned. Defaults to a slugified version of the account name.
	ProvisionedProductName *string `pulumi:"provisionedProductName"`
	// Assigned SSO user settings.
	Sso ControlTowerAwsAccountSso `pulumi:"sso"`
	// Key-value map of resource tags for the account.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ControlTowerAwsAccount resource.
type ControlTowerAwsAccountArgs struct {
	// If enabled, this will close the AWS account on resource deletion, beginning the 90-day suspension period. Otherwise, the
	// account will just be unenrolled from Control Tower.
	CloseAccountOnDelete pulumi.BoolPtrInput
	// Root email of the account.
	Email pulumi.StringInput
	// Name of the account.
	Name pulumi.StringPtrInput
	// Name of the Organizational Unit under which the account resides.
	OrganizationalUnit pulumi.StringInput
	// ID of the Organizational Unit to which the account should be moved when the resource is deleted. If no value is
	// provided, the account will not be moved.
	OrganizationalUnitIdOnDelete pulumi.StringPtrInput
	// Name of the path identifier of the product. This value is optional if the product has a default path, and required if
	// the product has more than one path. To list the paths for a product, use ListLaunchPaths.
	PathId pulumi.StringPtrInput
	// Name of the service catalog product that is provisioned. Defaults to a slugified version of the account name.
	ProvisionedProductName pulumi.StringPtrInput
	// Assigned SSO user settings.
	Sso ControlTowerAwsAccountSsoInput
	// Key-value map of resource tags for the account.
	Tags pulumi.StringMapInput
}

func (ControlTowerAwsAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*controlTowerAwsAccountArgs)(nil)).Elem()
}

type ControlTowerAwsAccountInput interface {
	pulumi.Input

	ToControlTowerAwsAccountOutput() ControlTowerAwsAccountOutput
	ToControlTowerAwsAccountOutputWithContext(ctx context.Context) ControlTowerAwsAccountOutput
}

func (*ControlTowerAwsAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlTowerAwsAccount)(nil)).Elem()
}

func (i *ControlTowerAwsAccount) ToControlTowerAwsAccountOutput() ControlTowerAwsAccountOutput {
	return i.ToControlTowerAwsAccountOutputWithContext(context.Background())
}

func (i *ControlTowerAwsAccount) ToControlTowerAwsAccountOutputWithContext(ctx context.Context) ControlTowerAwsAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlTowerAwsAccountOutput)
}

func (i *ControlTowerAwsAccount) ToOutput(ctx context.Context) pulumix.Output[*ControlTowerAwsAccount] {
	return pulumix.Output[*ControlTowerAwsAccount]{
		OutputState: i.ToControlTowerAwsAccountOutputWithContext(ctx).OutputState,
	}
}

// ControlTowerAwsAccountArrayInput is an input type that accepts ControlTowerAwsAccountArray and ControlTowerAwsAccountArrayOutput values.
// You can construct a concrete instance of `ControlTowerAwsAccountArrayInput` via:
//
//	ControlTowerAwsAccountArray{ ControlTowerAwsAccountArgs{...} }
type ControlTowerAwsAccountArrayInput interface {
	pulumi.Input

	ToControlTowerAwsAccountArrayOutput() ControlTowerAwsAccountArrayOutput
	ToControlTowerAwsAccountArrayOutputWithContext(context.Context) ControlTowerAwsAccountArrayOutput
}

type ControlTowerAwsAccountArray []ControlTowerAwsAccountInput

func (ControlTowerAwsAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ControlTowerAwsAccount)(nil)).Elem()
}

func (i ControlTowerAwsAccountArray) ToControlTowerAwsAccountArrayOutput() ControlTowerAwsAccountArrayOutput {
	return i.ToControlTowerAwsAccountArrayOutputWithContext(context.Background())
}

func (i ControlTowerAwsAccountArray) ToControlTowerAwsAccountArrayOutputWithContext(ctx context.Context) ControlTowerAwsAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlTowerAwsAccountArrayOutput)
}

func (i ControlTowerAwsAccountArray) ToOutput(ctx context.Context) pulumix.Output[[]*ControlTowerAwsAccount] {
	return pulumix.Output[[]*ControlTowerAwsAccount]{
		OutputState: i.ToControlTowerAwsAccountArrayOutputWithContext(ctx).OutputState,
	}
}

// ControlTowerAwsAccountMapInput is an input type that accepts ControlTowerAwsAccountMap and ControlTowerAwsAccountMapOutput values.
// You can construct a concrete instance of `ControlTowerAwsAccountMapInput` via:
//
//	ControlTowerAwsAccountMap{ "key": ControlTowerAwsAccountArgs{...} }
type ControlTowerAwsAccountMapInput interface {
	pulumi.Input

	ToControlTowerAwsAccountMapOutput() ControlTowerAwsAccountMapOutput
	ToControlTowerAwsAccountMapOutputWithContext(context.Context) ControlTowerAwsAccountMapOutput
}

type ControlTowerAwsAccountMap map[string]ControlTowerAwsAccountInput

func (ControlTowerAwsAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ControlTowerAwsAccount)(nil)).Elem()
}

func (i ControlTowerAwsAccountMap) ToControlTowerAwsAccountMapOutput() ControlTowerAwsAccountMapOutput {
	return i.ToControlTowerAwsAccountMapOutputWithContext(context.Background())
}

func (i ControlTowerAwsAccountMap) ToControlTowerAwsAccountMapOutputWithContext(ctx context.Context) ControlTowerAwsAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlTowerAwsAccountMapOutput)
}

func (i ControlTowerAwsAccountMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ControlTowerAwsAccount] {
	return pulumix.Output[map[string]*ControlTowerAwsAccount]{
		OutputState: i.ToControlTowerAwsAccountMapOutputWithContext(ctx).OutputState,
	}
}

type ControlTowerAwsAccountOutput struct{ *pulumi.OutputState }

func (ControlTowerAwsAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlTowerAwsAccount)(nil)).Elem()
}

func (o ControlTowerAwsAccountOutput) ToControlTowerAwsAccountOutput() ControlTowerAwsAccountOutput {
	return o
}

func (o ControlTowerAwsAccountOutput) ToControlTowerAwsAccountOutputWithContext(ctx context.Context) ControlTowerAwsAccountOutput {
	return o
}

func (o ControlTowerAwsAccountOutput) ToOutput(ctx context.Context) pulumix.Output[*ControlTowerAwsAccount] {
	return pulumix.Output[*ControlTowerAwsAccount]{
		OutputState: o.OutputState,
	}
}

// ID of the AWS account.
func (o ControlTowerAwsAccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlTowerAwsAccount) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// If enabled, this will close the AWS account on resource deletion, beginning the 90-day suspension period. Otherwise, the
// account will just be unenrolled from Control Tower.
func (o ControlTowerAwsAccountOutput) CloseAccountOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ControlTowerAwsAccount) pulumi.BoolPtrOutput { return v.CloseAccountOnDelete }).(pulumi.BoolPtrOutput)
}

// Root email of the account.
func (o ControlTowerAwsAccountOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlTowerAwsAccount) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Name of the account.
func (o ControlTowerAwsAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlTowerAwsAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name of the Organizational Unit under which the account resides.
func (o ControlTowerAwsAccountOutput) OrganizationalUnit() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlTowerAwsAccount) pulumi.StringOutput { return v.OrganizationalUnit }).(pulumi.StringOutput)
}

// ID of the Organizational Unit to which the account should be moved when the resource is deleted. If no value is
// provided, the account will not be moved.
func (o ControlTowerAwsAccountOutput) OrganizationalUnitIdOnDelete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlTowerAwsAccount) pulumi.StringPtrOutput { return v.OrganizationalUnitIdOnDelete }).(pulumi.StringPtrOutput)
}

// Name of the path identifier of the product. This value is optional if the product has a default path, and required if
// the product has more than one path. To list the paths for a product, use ListLaunchPaths.
func (o ControlTowerAwsAccountOutput) PathId() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlTowerAwsAccount) pulumi.StringOutput { return v.PathId }).(pulumi.StringOutput)
}

// Name of the service catalog product that is provisioned. Defaults to a slugified version of the account name.
func (o ControlTowerAwsAccountOutput) ProvisionedProductName() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlTowerAwsAccount) pulumi.StringOutput { return v.ProvisionedProductName }).(pulumi.StringOutput)
}

// Assigned SSO user settings.
func (o ControlTowerAwsAccountOutput) Sso() ControlTowerAwsAccountSsoOutput {
	return o.ApplyT(func(v *ControlTowerAwsAccount) ControlTowerAwsAccountSsoOutput { return v.Sso }).(ControlTowerAwsAccountSsoOutput)
}

// Key-value map of resource tags for the account.
func (o ControlTowerAwsAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ControlTowerAwsAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

type ControlTowerAwsAccountArrayOutput struct{ *pulumi.OutputState }

func (ControlTowerAwsAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ControlTowerAwsAccount)(nil)).Elem()
}

func (o ControlTowerAwsAccountArrayOutput) ToControlTowerAwsAccountArrayOutput() ControlTowerAwsAccountArrayOutput {
	return o
}

func (o ControlTowerAwsAccountArrayOutput) ToControlTowerAwsAccountArrayOutputWithContext(ctx context.Context) ControlTowerAwsAccountArrayOutput {
	return o
}

func (o ControlTowerAwsAccountArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ControlTowerAwsAccount] {
	return pulumix.Output[[]*ControlTowerAwsAccount]{
		OutputState: o.OutputState,
	}
}

func (o ControlTowerAwsAccountArrayOutput) Index(i pulumi.IntInput) ControlTowerAwsAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ControlTowerAwsAccount {
		return vs[0].([]*ControlTowerAwsAccount)[vs[1].(int)]
	}).(ControlTowerAwsAccountOutput)
}

type ControlTowerAwsAccountMapOutput struct{ *pulumi.OutputState }

func (ControlTowerAwsAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ControlTowerAwsAccount)(nil)).Elem()
}

func (o ControlTowerAwsAccountMapOutput) ToControlTowerAwsAccountMapOutput() ControlTowerAwsAccountMapOutput {
	return o
}

func (o ControlTowerAwsAccountMapOutput) ToControlTowerAwsAccountMapOutputWithContext(ctx context.Context) ControlTowerAwsAccountMapOutput {
	return o
}

func (o ControlTowerAwsAccountMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ControlTowerAwsAccount] {
	return pulumix.Output[map[string]*ControlTowerAwsAccount]{
		OutputState: o.OutputState,
	}
}

func (o ControlTowerAwsAccountMapOutput) MapIndex(k pulumi.StringInput) ControlTowerAwsAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ControlTowerAwsAccount {
		return vs[0].(map[string]*ControlTowerAwsAccount)[vs[1].(string)]
	}).(ControlTowerAwsAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ControlTowerAwsAccountInput)(nil)).Elem(), &ControlTowerAwsAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*ControlTowerAwsAccountArrayInput)(nil)).Elem(), ControlTowerAwsAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ControlTowerAwsAccountMapInput)(nil)).Elem(), ControlTowerAwsAccountMap{})
	pulumi.RegisterOutputType(ControlTowerAwsAccountOutput{})
	pulumi.RegisterOutputType(ControlTowerAwsAccountArrayOutput{})
	pulumi.RegisterOutputType(ControlTowerAwsAccountMapOutput{})
}
