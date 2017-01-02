# Exporting Pachyderm Data with SQL

## Intro

In this tutorial, we're going to set up a simple *egress pipeline* that exports
data using SQL. We'll move data from a dummy pachyderm repo containing some
JSON records into an Amazon Redshift database, but this pipeline could easily
be modified to export data to PostgreSQL or anything similar.

## Setup

In this demo, we'll need a pachyderm cluster to run our jobs, and we'll also
need a Redshift cluster. We'll assume you have a Pachyderm cluster up and
running, but if not, check out our [Getting Started
guide](http://docs.pachyderm.io/en/v1.3.2/getting_started/getting_started.html).

To set up Redshift, you'll need to come up with a cluster password and put it
in the file `redshift_passwd.txt`.  Then run the following commands (feel free
to replace e.g. `CLUSTER_IDENTIFIER` with any values you prefer):

```
$ CLUSTER_IDENTIFIER=test-redshift-cluster
$ DB_NAME=test-db
$ USERNAME=testuser

$ aws redshift create-cluster \
  --cluster-identifier=${CLUSTER_IDENTIFIER} \
  --cluster-type=single-node \
  --node-type=dc1.large \
  --master-username=${USERNAME} \
  --master-user-password="$( cat redshift_passwd.txt )"
{
    "Cluster": {
        "PubliclyAccessible": true,
        "ClusterSubnetGroupName": "default",
        "NodeType": "dc1.large",
        "ClusterSecurityGroups": [],
        "ClusterVersion": "1.0",
        "AutomatedSnapshotRetentionPeriod": 1,
        "Tags": [],
        "ClusterIdentifier": "test-redshift-cluster",
        "VpcId": "vpc-123abcde",
        "ClusterStatus": "creating",
        "Encrypted": false,
        "PendingModifiedValues": {
            "MasterUserPassword": "****"
        },
        "PreferredMaintenanceWindow": "mon:12:00-mon:12:30",
        "VpcSecurityGroups": [],
        "AllowVersionUpgrade": true,
        "MasterUsername": "testuser",
        "ClusterParameterGroups": [
            {
                "ParameterGroupName": "default.redshift-1.0",
                "ParameterApplyStatus": "in-sync"
            }
        ],
        "NumberOfNodes": 1
    }
}
```

To access the redshift cluster, you'll need to modify its security group to
allow ingress traffic from Pachyderm.

