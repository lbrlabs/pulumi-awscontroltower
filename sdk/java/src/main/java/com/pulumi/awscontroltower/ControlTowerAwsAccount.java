// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.awscontroltower;

import com.pulumi.awscontroltower.ControlTowerAwsAccountArgs;
import com.pulumi.awscontroltower.Utilities;
import com.pulumi.awscontroltower.inputs.ControlTowerAwsAccountState;
import com.pulumi.awscontroltower.outputs.ControlTowerAwsAccountSso;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="awscontroltower:index/controlTowerAwsAccount:controlTowerAwsAccount")
public class ControlTowerAwsAccount extends com.pulumi.resources.CustomResource {
    /**
     * ID of the AWS account.
     * 
     */
    @Export(name="accountId", type=String.class, parameters={})
    private Output<String> accountId;

    /**
     * @return ID of the AWS account.
     * 
     */
    public Output<String> accountId() {
        return this.accountId;
    }
    /**
     * If enabled, this will close the AWS account on resource deletion, beginning the 90-day suspension period. Otherwise, the
     * account will just be unenrolled from Control Tower.
     * 
     */
    @Export(name="closeAccountOnDelete", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> closeAccountOnDelete;

    /**
     * @return If enabled, this will close the AWS account on resource deletion, beginning the 90-day suspension period. Otherwise, the
     * account will just be unenrolled from Control Tower.
     * 
     */
    public Output<Optional<Boolean>> closeAccountOnDelete() {
        return Codegen.optional(this.closeAccountOnDelete);
    }
    /**
     * Root email of the account.
     * 
     */
    @Export(name="email", type=String.class, parameters={})
    private Output<String> email;

    /**
     * @return Root email of the account.
     * 
     */
    public Output<String> email() {
        return this.email;
    }
    /**
     * Name of the account.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the account.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Name of the Organizational Unit under which the account resides.
     * 
     */
    @Export(name="organizationalUnit", type=String.class, parameters={})
    private Output<String> organizationalUnit;

    /**
     * @return Name of the Organizational Unit under which the account resides.
     * 
     */
    public Output<String> organizationalUnit() {
        return this.organizationalUnit;
    }
    /**
     * ID of the Organizational Unit to which the account should be moved when the resource is deleted. If no value is
     * provided, the account will not be moved.
     * 
     */
    @Export(name="organizationalUnitIdOnDelete", type=String.class, parameters={})
    private Output</* @Nullable */ String> organizationalUnitIdOnDelete;

    /**
     * @return ID of the Organizational Unit to which the account should be moved when the resource is deleted. If no value is
     * provided, the account will not be moved.
     * 
     */
    public Output<Optional<String>> organizationalUnitIdOnDelete() {
        return Codegen.optional(this.organizationalUnitIdOnDelete);
    }
    /**
     * Name of the path identifier of the product. This value is optional if the product has a default path, and required if
     * the product has more than one path. To list the paths for a product, use ListLaunchPaths.
     * 
     */
    @Export(name="pathId", type=String.class, parameters={})
    private Output<String> pathId;

    /**
     * @return Name of the path identifier of the product. This value is optional if the product has a default path, and required if
     * the product has more than one path. To list the paths for a product, use ListLaunchPaths.
     * 
     */
    public Output<String> pathId() {
        return this.pathId;
    }
    /**
     * Name of the service catalog product that is provisioned. Defaults to a slugified version of the account name.
     * 
     */
    @Export(name="provisionedProductName", type=String.class, parameters={})
    private Output<String> provisionedProductName;

    /**
     * @return Name of the service catalog product that is provisioned. Defaults to a slugified version of the account name.
     * 
     */
    public Output<String> provisionedProductName() {
        return this.provisionedProductName;
    }
    /**
     * Assigned SSO user settings.
     * 
     */
    @Export(name="sso", type=ControlTowerAwsAccountSso.class, parameters={})
    private Output<ControlTowerAwsAccountSso> sso;

    /**
     * @return Assigned SSO user settings.
     * 
     */
    public Output<ControlTowerAwsAccountSso> sso() {
        return this.sso;
    }
    /**
     * Key-value map of resource tags for the account.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags for the account.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ControlTowerAwsAccount(String name) {
        this(name, ControlTowerAwsAccountArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ControlTowerAwsAccount(String name, ControlTowerAwsAccountArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ControlTowerAwsAccount(String name, ControlTowerAwsAccountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("awscontroltower:index/controlTowerAwsAccount:controlTowerAwsAccount", name, args == null ? ControlTowerAwsAccountArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ControlTowerAwsAccount(String name, Output<String> id, @Nullable ControlTowerAwsAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("awscontroltower:index/controlTowerAwsAccount:controlTowerAwsAccount", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ControlTowerAwsAccount get(String name, Output<String> id, @Nullable ControlTowerAwsAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ControlTowerAwsAccount(name, id, state, options);
    }
}