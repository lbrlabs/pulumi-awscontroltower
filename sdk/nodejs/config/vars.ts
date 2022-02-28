// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("awscontroltower");

/**
 * This is the AWS access key. It must be provided, but it can also be sourced from the `AWS_ACCESS_KEY_ID` environment
 * variable, or via a shared credentials file if `profile` is specified.
 */
export declare const accessKey: string | undefined;
Object.defineProperty(exports, "accessKey", {
    get() {
        return __config.get("accessKey");
    },
    enumerable: true,
});

/**
 * Settings for making use of the AWS Assume Role functionality.
 */
export declare const assumeRole: outputs.config.AssumeRole | undefined;
Object.defineProperty(exports, "assumeRole", {
    get() {
        return __config.getObject<outputs.config.AssumeRole>("assumeRole");
    },
    enumerable: true,
});

/**
 * This is the maximum number of times an API call is retried, in the case where requests are being throttled or
 * experiencing transient failures. The delay between the subsequent API calls increases exponentially. If omitted, the
 * default value is `25`.
 */
export declare const maxRetries: number | undefined;
Object.defineProperty(exports, "maxRetries", {
    get() {
        return __config.getObject<number>("maxRetries");
    },
    enumerable: true,
});

/**
 * This is the AWS profile name as set in the shared credentials file.
 */
export declare const profile: string | undefined;
Object.defineProperty(exports, "profile", {
    get() {
        return __config.get("profile");
    },
    enumerable: true,
});

/**
 * This is the AWS region. It must be provided, but it can also be sourced from the `AWS_DEFAULT_REGION` environment
 * variables, or via a shared credentials file if `profile` is specified.
 */
export declare const region: string | undefined;
Object.defineProperty(exports, "region", {
    get() {
        return __config.get("region");
    },
    enumerable: true,
});

/**
 * This is the AWS secret key. It must be provided, but it can also be sourced from the `AWS_SECRET_ACCESS_KEY` environment
 * variable, or via a shared credentials file if `profile` is specified.
 */
export declare const secretKey: string | undefined;
Object.defineProperty(exports, "secretKey", {
    get() {
        return __config.get("secretKey");
    },
    enumerable: true,
});

/**
 * This is the path to the shared credentials file. If this is not set and a profile is specified, `~/.aws/credentials`
 * will be used.
 */
export declare const sharedCredentialsFile: string | undefined;
Object.defineProperty(exports, "sharedCredentialsFile", {
    get() {
        return __config.get("sharedCredentialsFile");
    },
    enumerable: true,
});

/**
 * Skip the credentials validation via the STS API. Useful for AWS API implementations that do not have STS available or
 * implemented.
 */
export declare const skipCredentialsValidation: boolean | undefined;
Object.defineProperty(exports, "skipCredentialsValidation", {
    get() {
        return __config.getObject<boolean>("skipCredentialsValidation");
    },
    enumerable: true,
});

/**
 * Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to
 * `true` prevents Terraform from authenticating via the Metadata API. You may need to use other authentication methods
 * like static credentials, configuration variables, or environment variables.
 */
export declare const skipMetadataApiCheck: boolean | undefined;
Object.defineProperty(exports, "skipMetadataApiCheck", {
    get() {
        return __config.getObject<boolean>("skipMetadataApiCheck");
    },
    enumerable: true,
});

/**
 * Skip requesting the account ID. Useful for AWS API implementations that do not have the IAM, STS API, or metadata API.
 */
export declare const skipRequestingAccountId: boolean | undefined;
Object.defineProperty(exports, "skipRequestingAccountId", {
    get() {
        return __config.getObject<boolean>("skipRequestingAccountId");
    },
    enumerable: true,
});

/**
 * Session token for validating temporary credentials. Typically provided after successful identity federation or
 * Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit
 * MFA code used to get temporary credentials. It can also be sourced from the `AWS_SESSION_TOKEN` environment variable.
 */
export declare const token: string | undefined;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token");
    },
    enumerable: true,
});

