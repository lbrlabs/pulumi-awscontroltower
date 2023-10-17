# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AssumeRole',
]

@pulumi.output_type
class AssumeRole(dict):
    def __init__(__self__, *,
                 duration_seconds: Optional[int] = None,
                 external_id: Optional[str] = None,
                 policy: Optional[str] = None,
                 policy_arns: Optional[Sequence[str]] = None,
                 role_arn: Optional[str] = None,
                 session_name: Optional[str] = None,
                 tags: Optional[Mapping[str, str]] = None,
                 transitive_tag_keys: Optional[Sequence[str]] = None):
        AssumeRole._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            duration_seconds=duration_seconds,
            external_id=external_id,
            policy=policy,
            policy_arns=policy_arns,
            role_arn=role_arn,
            session_name=session_name,
            tags=tags,
            transitive_tag_keys=transitive_tag_keys,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             duration_seconds: Optional[int] = None,
             external_id: Optional[str] = None,
             policy: Optional[str] = None,
             policy_arns: Optional[Sequence[str]] = None,
             role_arn: Optional[str] = None,
             session_name: Optional[str] = None,
             tags: Optional[Mapping[str, str]] = None,
             transitive_tag_keys: Optional[Sequence[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'durationSeconds' in kwargs:
            duration_seconds = kwargs['durationSeconds']
        if 'externalId' in kwargs:
            external_id = kwargs['externalId']
        if 'policyArns' in kwargs:
            policy_arns = kwargs['policyArns']
        if 'roleArn' in kwargs:
            role_arn = kwargs['roleArn']
        if 'sessionName' in kwargs:
            session_name = kwargs['sessionName']
        if 'transitiveTagKeys' in kwargs:
            transitive_tag_keys = kwargs['transitiveTagKeys']

        if duration_seconds is not None:
            _setter("duration_seconds", duration_seconds)
        if external_id is not None:
            _setter("external_id", external_id)
        if policy is not None:
            _setter("policy", policy)
        if policy_arns is not None:
            _setter("policy_arns", policy_arns)
        if role_arn is not None:
            _setter("role_arn", role_arn)
        if session_name is not None:
            _setter("session_name", session_name)
        if tags is not None:
            _setter("tags", tags)
        if transitive_tag_keys is not None:
            _setter("transitive_tag_keys", transitive_tag_keys)

    @property
    @pulumi.getter(name="durationSeconds")
    def duration_seconds(self) -> Optional[int]:
        return pulumi.get(self, "duration_seconds")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> Optional[str]:
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter
    def policy(self) -> Optional[str]:
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="policyArns")
    def policy_arns(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "policy_arns")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="sessionName")
    def session_name(self) -> Optional[str]:
        return pulumi.get(self, "session_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transitiveTagKeys")
    def transitive_tag_keys(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "transitive_tag_keys")


