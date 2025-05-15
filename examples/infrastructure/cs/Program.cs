using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Duplocloud = DuploCloud.Pulumi;

return await Deployment.RunAsync(() =>
{

    // Create a new infrastructure in AWS
    var infra = new Duplocloud.Infrastructure("pulumi-infra", new Duplocloud.InfrastructureArgs
    {
        InfraName = "pulumi-infra",
        Cloud = 0,
        Region = "us-east-2",
        Azcount = 2,
        EnableK8Cluster = true,
        AddressPrefix = "10.22.0.0/16",
        SubnetCidr = 24,
    });
    infra.VpcId.Apply(id =>
    {
        Pulumi.Log.Info($"VPC ID : {id}");
        return id;
    });
    return new Dictionary<string, object>
    {
        { "VpcId",  infra.VpcId },
    };

});