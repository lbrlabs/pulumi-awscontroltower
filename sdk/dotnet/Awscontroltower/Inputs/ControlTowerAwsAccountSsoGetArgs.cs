// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Awscontroltower.Inputs
{

    public sealed class ControlTowerAwsAccountSsoGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        [Input("firstName", required: true)]
        public Input<string> FirstName { get; set; } = null!;

        [Input("lastName", required: true)]
        public Input<string> LastName { get; set; } = null!;

        public ControlTowerAwsAccountSsoGetArgs()
        {
        }
        public static new ControlTowerAwsAccountSsoGetArgs Empty => new ControlTowerAwsAccountSsoGetArgs();
    }
}
