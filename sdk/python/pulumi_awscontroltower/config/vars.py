# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

import types

__config__ = pulumi.Config('awscontroltower')


class _ExportableConfig(types.ModuleType):
    @property
    def access_key(self) -> Optional[str]:
        """
        This is the AWS access key. It must be provided, but it can also be sourced from the `AWS_ACCESS_KEY_ID` environment
        variable, or via a shared credentials file if `profile` is specified.
        """
        return __config__.get('accessKey')

    @property
    def assume_role(self) -> Optional[str]:
        """
        Settings for making use of the AWS Assume Role functionality.
        """
        return __config__.get('assumeRole')

    @property
    def max_retries(self) -> Optional[int]:
        """
        This is the maximum number of times an API call is retried, in the case where requests are being throttled or
        experiencing transient failures. The delay between the subsequent API calls increases exponentially. If omitted, the
        default value is `25`.
        """
        return __config__.get_int('maxRetries')

    @property
    def profile(self) -> Optional[str]:
        """
        This is the AWS profile name as set in the shared credentials file.
        """
        return __config__.get('profile')

    @property
    def region(self) -> Optional[str]:
        """
        This is the AWS region. It must be provided, but it can also be sourced from the `AWS_DEFAULT_REGION` environment
        variables, or via a shared credentials file if `profile` is specified.
        """
        return __config__.get('region')

    @property
    def secret_key(self) -> Optional[str]:
        """
        This is the AWS secret key. It must be provided, but it can also be sourced from the `AWS_SECRET_ACCESS_KEY` environment
        variable, or via a shared credentials file if `profile` is specified.
        """
        return __config__.get('secretKey')

    @property
    def shared_credentials_file(self) -> Optional[str]:
        """
        This is the path to the shared credentials file. If this is not set and a profile is specified, `~/.aws/credentials`
        will be used.
        """
        return __config__.get('sharedCredentialsFile')

    @property
    def skip_credentials_validation(self) -> Optional[bool]:
        """
        Skip the credentials validation via the STS API. Useful for AWS API implementations that do not have STS available or
        implemented.
        """
        return __config__.get_bool('skipCredentialsValidation')

    @property
    def skip_metadata_api_check(self) -> Optional[bool]:
        """
        Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to
        `true` prevents Terraform from authenticating via the Metadata API. You may need to use other authentication methods
        like static credentials, configuration variables, or environment variables.
        """
        return __config__.get_bool('skipMetadataApiCheck')

    @property
    def skip_requesting_account_id(self) -> Optional[bool]:
        """
        Skip requesting the account ID. Useful for AWS API implementations that do not have the IAM, STS API, or metadata API.
        """
        return __config__.get_bool('skipRequestingAccountId')

    @property
    def token(self) -> Optional[str]:
        """
        Session token for validating temporary credentials. Typically provided after successful identity federation or
        Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit
        MFA code used to get temporary credentials. It can also be sourced from the `AWS_SESSION_TOKEN` environment variable.
        """
        return __config__.get('token')

