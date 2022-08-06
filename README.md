AWS Resource Scanner
====================
A project to query AWS resource in use within one AWS account

Look up available regions within your account, get brief information about AWS resources in use within that region, and create a list.

## Services currently available for inquiry
- APIGateway
- CloudFront (Distribution) / Global
- Cognito (UserPool)
- EBS (Volume)
- EC2 (Instance)
- ECR (Regository)
- ECS (Cluster, Service)
- EFS (FileSystem)
- Elasticache (ReplicationGroup, Cluster)
- Elastic Beanstalk (Application, Environment)
- ELB (LoadBalance)
- ELB (LoadBalance)
- Lambda (Function)
- QLDB (Ledger)
- RDS (DBInstance)
- Route53 (Hosted Zone) / Global
- S3 (Bucket) / Global
- SES (Identity)
- SNS (Topic)
- SQS (Queue)

## Development scheduled
Scan and represent data for various resources such as Internat and NAT Gateway, Endpoint for vpc configuration.

be going to implement a search for other resources you need.

## How to use?
### 0. Prior requirements
"Go" is basically used, so "Go" must be installed in advance. [Installation Golang](https://go.dev/doc/install)

also, need an AWS IAM role arn with a "ReadOnlyAccess" policy for resource scanning.
### 1. Build
```
go build -o ./bin/<excutable_file>
```
### 2. Output
```
./bin/<excutable_file> <RoleArnForScan>
```