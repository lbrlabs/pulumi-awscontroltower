// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.awscontroltower;

import com.pulumi.awscontroltower.config.inputs.AssumeRole;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("awscontroltower");
/**
 * This is the AWS access key. It must be provided, but it can also be sourced from the `AWS_ACCESS_KEY_ID` environment
 * variable, or via a shared credentials file if `profile` is specified.
 * 
 */
    public Optional<String> accessKey() {
        return Codegen.stringProp("accessKey").config(config).get();
    }
/**
 * Settings for making use of the AWS Assume Role functionality.
 * 
 */
    public Optional<AssumeRole> assumeRole() {
        return Codegen.objectProp("assumeRole", AssumeRole.class).config(config).get();
    }
/**
 * This is the maximum number of times an API call is retried, in the case where requests are being throttled or
 * experiencing transient failures. The delay between the subsequent API calls increases exponentially. If omitted, the
 * default value is `25`.
 * 
 */
    public Optional<Integer> maxRetries() {
        return Codegen.integerProp("maxRetries").config(config).get();
    }
/**
 * This is the AWS profile name as set in the shared credentials file.
 * 
 */
    public Optional<String> profile() {
        return Codegen.stringProp("profile").config(config).get();
    }
/**
 * This is the AWS region. It must be provided, but it can also be sourced from the `AWS_DEFAULT_REGION` environment
 * variables, or via a shared credentials file if `profile` is specified.
 * 
 */
    public String region() {
        return Codegen.stringProp("region").config(config).require();
    }
/**
 * This is the AWS secret key. It must be provided, but it can also be sourced from the `AWS_SECRET_ACCESS_KEY` environment
 * variable, or via a shared credentials file if `profile` is specified.
 * 
 */
    public Optional<String> secretKey() {
        return Codegen.stringProp("secretKey").config(config).get();
    }
/**
 * This is the path to the shared credentials file. If this is not set and a profile is specified, `~/.aws/credentials`
 * will be used.
 * 
 */
    public Optional<String> sharedCredentialsFile() {
        return Codegen.stringProp("sharedCredentialsFile").config(config).get();
    }
/**
 * Skip the credentials validation via the STS API. Useful for AWS API implementations that do not have STS available or
 * implemented.
 * 
 */
    public Optional<Boolean> skipCredentialsValidation() {
        return Codegen.booleanProp("skipCredentialsValidation").config(config).get();
    }
/**
 * Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to
 * `true` prevents Terraform from authenticating via the Metadata API. You may need to use other authentication methods
 * like static credentials, configuration variables, or environment variables.
 * 
 */
    public Optional<Boolean> skipMetadataApiCheck() {
        return Codegen.booleanProp("skipMetadataApiCheck").config(config).get();
    }
/**
 * Skip requesting the account ID. Useful for AWS API implementations that do not have the IAM, STS API, or metadata API.
 * 
 */
    public Optional<Boolean> skipRequestingAccountId() {
        return Codegen.booleanProp("skipRequestingAccountId").config(config).get();
    }
/**
 * Session token for validating temporary credentials. Typically provided after successful identity federation or
 * Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit
 * MFA code used to get temporary credentials. It can also be sourced from the `AWS_SESSION_TOKEN` environment variable.
 * 
 */
    public Optional<String> token() {
        return Codegen.stringProp("token").config(config).get();
    }
}
