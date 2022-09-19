// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.awscontroltower.inputs;

import com.pulumi.awscontroltower.inputs.ControlTowerAwsAccountSsoArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ControlTowerAwsAccountState extends com.pulumi.resources.ResourceArgs {

    public static final ControlTowerAwsAccountState Empty = new ControlTowerAwsAccountState();

    /**
     * ID of the AWS account.
     * 
     */
    @Import(name="accountId")
    private @Nullable Output<String> accountId;

    /**
     * @return ID of the AWS account.
     * 
     */
    public Optional<Output<String>> accountId() {
        return Optional.ofNullable(this.accountId);
    }

    /**
     * If enabled, this will close the AWS account on resource deletion, beginning the 90-day suspension period. Otherwise, the
     * account will just be unenrolled from Control Tower.
     * 
     */
    @Import(name="closeAccountOnDelete")
    private @Nullable Output<Boolean> closeAccountOnDelete;

    /**
     * @return If enabled, this will close the AWS account on resource deletion, beginning the 90-day suspension period. Otherwise, the
     * account will just be unenrolled from Control Tower.
     * 
     */
    public Optional<Output<Boolean>> closeAccountOnDelete() {
        return Optional.ofNullable(this.closeAccountOnDelete);
    }

    /**
     * Root email of the account.
     * 
     */
    @Import(name="email")
    private @Nullable Output<String> email;

    /**
     * @return Root email of the account.
     * 
     */
    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * Name of the account.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the account.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Name of the Organizational Unit under which the account resides.
     * 
     */
    @Import(name="organizationalUnit")
    private @Nullable Output<String> organizationalUnit;

    /**
     * @return Name of the Organizational Unit under which the account resides.
     * 
     */
    public Optional<Output<String>> organizationalUnit() {
        return Optional.ofNullable(this.organizationalUnit);
    }

    /**
     * ID of the Organizational Unit to which the account should be moved when the resource is deleted. If no value is
     * provided, the account will not be moved.
     * 
     */
    @Import(name="organizationalUnitIdOnDelete")
    private @Nullable Output<String> organizationalUnitIdOnDelete;

    /**
     * @return ID of the Organizational Unit to which the account should be moved when the resource is deleted. If no value is
     * provided, the account will not be moved.
     * 
     */
    public Optional<Output<String>> organizationalUnitIdOnDelete() {
        return Optional.ofNullable(this.organizationalUnitIdOnDelete);
    }

    /**
     * Name of the path identifier of the product. This value is optional if the product has a default path, and required if
     * the product has more than one path. To list the paths for a product, use ListLaunchPaths.
     * 
     */
    @Import(name="pathId")
    private @Nullable Output<String> pathId;

    /**
     * @return Name of the path identifier of the product. This value is optional if the product has a default path, and required if
     * the product has more than one path. To list the paths for a product, use ListLaunchPaths.
     * 
     */
    public Optional<Output<String>> pathId() {
        return Optional.ofNullable(this.pathId);
    }

    /**
     * Name of the service catalog product that is provisioned. Defaults to a slugified version of the account name.
     * 
     */
    @Import(name="provisionedProductName")
    private @Nullable Output<String> provisionedProductName;

    /**
     * @return Name of the service catalog product that is provisioned. Defaults to a slugified version of the account name.
     * 
     */
    public Optional<Output<String>> provisionedProductName() {
        return Optional.ofNullable(this.provisionedProductName);
    }

    /**
     * Assigned SSO user settings.
     * 
     */
    @Import(name="sso")
    private @Nullable Output<ControlTowerAwsAccountSsoArgs> sso;

    /**
     * @return Assigned SSO user settings.
     * 
     */
    public Optional<Output<ControlTowerAwsAccountSsoArgs>> sso() {
        return Optional.ofNullable(this.sso);
    }

    /**
     * Key-value map of resource tags for the account.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags for the account.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private ControlTowerAwsAccountState() {}

    private ControlTowerAwsAccountState(ControlTowerAwsAccountState $) {
        this.accountId = $.accountId;
        this.closeAccountOnDelete = $.closeAccountOnDelete;
        this.email = $.email;
        this.name = $.name;
        this.organizationalUnit = $.organizationalUnit;
        this.organizationalUnitIdOnDelete = $.organizationalUnitIdOnDelete;
        this.pathId = $.pathId;
        this.provisionedProductName = $.provisionedProductName;
        this.sso = $.sso;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ControlTowerAwsAccountState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ControlTowerAwsAccountState $;

        public Builder() {
            $ = new ControlTowerAwsAccountState();
        }

        public Builder(ControlTowerAwsAccountState defaults) {
            $ = new ControlTowerAwsAccountState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountId ID of the AWS account.
         * 
         * @return builder
         * 
         */
        public Builder accountId(@Nullable Output<String> accountId) {
            $.accountId = accountId;
            return this;
        }

        /**
         * @param accountId ID of the AWS account.
         * 
         * @return builder
         * 
         */
        public Builder accountId(String accountId) {
            return accountId(Output.of(accountId));
        }

        /**
         * @param closeAccountOnDelete If enabled, this will close the AWS account on resource deletion, beginning the 90-day suspension period. Otherwise, the
         * account will just be unenrolled from Control Tower.
         * 
         * @return builder
         * 
         */
        public Builder closeAccountOnDelete(@Nullable Output<Boolean> closeAccountOnDelete) {
            $.closeAccountOnDelete = closeAccountOnDelete;
            return this;
        }

        /**
         * @param closeAccountOnDelete If enabled, this will close the AWS account on resource deletion, beginning the 90-day suspension period. Otherwise, the
         * account will just be unenrolled from Control Tower.
         * 
         * @return builder
         * 
         */
        public Builder closeAccountOnDelete(Boolean closeAccountOnDelete) {
            return closeAccountOnDelete(Output.of(closeAccountOnDelete));
        }

        /**
         * @param email Root email of the account.
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email Root email of the account.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param name Name of the account.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the account.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param organizationalUnit Name of the Organizational Unit under which the account resides.
         * 
         * @return builder
         * 
         */
        public Builder organizationalUnit(@Nullable Output<String> organizationalUnit) {
            $.organizationalUnit = organizationalUnit;
            return this;
        }

        /**
         * @param organizationalUnit Name of the Organizational Unit under which the account resides.
         * 
         * @return builder
         * 
         */
        public Builder organizationalUnit(String organizationalUnit) {
            return organizationalUnit(Output.of(organizationalUnit));
        }

        /**
         * @param organizationalUnitIdOnDelete ID of the Organizational Unit to which the account should be moved when the resource is deleted. If no value is
         * provided, the account will not be moved.
         * 
         * @return builder
         * 
         */
        public Builder organizationalUnitIdOnDelete(@Nullable Output<String> organizationalUnitIdOnDelete) {
            $.organizationalUnitIdOnDelete = organizationalUnitIdOnDelete;
            return this;
        }

        /**
         * @param organizationalUnitIdOnDelete ID of the Organizational Unit to which the account should be moved when the resource is deleted. If no value is
         * provided, the account will not be moved.
         * 
         * @return builder
         * 
         */
        public Builder organizationalUnitIdOnDelete(String organizationalUnitIdOnDelete) {
            return organizationalUnitIdOnDelete(Output.of(organizationalUnitIdOnDelete));
        }

        /**
         * @param pathId Name of the path identifier of the product. This value is optional if the product has a default path, and required if
         * the product has more than one path. To list the paths for a product, use ListLaunchPaths.
         * 
         * @return builder
         * 
         */
        public Builder pathId(@Nullable Output<String> pathId) {
            $.pathId = pathId;
            return this;
        }

        /**
         * @param pathId Name of the path identifier of the product. This value is optional if the product has a default path, and required if
         * the product has more than one path. To list the paths for a product, use ListLaunchPaths.
         * 
         * @return builder
         * 
         */
        public Builder pathId(String pathId) {
            return pathId(Output.of(pathId));
        }

        /**
         * @param provisionedProductName Name of the service catalog product that is provisioned. Defaults to a slugified version of the account name.
         * 
         * @return builder
         * 
         */
        public Builder provisionedProductName(@Nullable Output<String> provisionedProductName) {
            $.provisionedProductName = provisionedProductName;
            return this;
        }

        /**
         * @param provisionedProductName Name of the service catalog product that is provisioned. Defaults to a slugified version of the account name.
         * 
         * @return builder
         * 
         */
        public Builder provisionedProductName(String provisionedProductName) {
            return provisionedProductName(Output.of(provisionedProductName));
        }

        /**
         * @param sso Assigned SSO user settings.
         * 
         * @return builder
         * 
         */
        public Builder sso(@Nullable Output<ControlTowerAwsAccountSsoArgs> sso) {
            $.sso = sso;
            return this;
        }

        /**
         * @param sso Assigned SSO user settings.
         * 
         * @return builder
         * 
         */
        public Builder sso(ControlTowerAwsAccountSsoArgs sso) {
            return sso(Output.of(sso));
        }

        /**
         * @param tags Key-value map of resource tags for the account.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags for the account.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public ControlTowerAwsAccountState build() {
            return $;
        }
    }

}
