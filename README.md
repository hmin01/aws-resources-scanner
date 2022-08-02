AWS Resource Scanner
====================
A project to query AWS resource in use within one AWS account

Look up available regions within your account, get brief information about AWS resources in use within that region, and create a list.

## Services currently available for inquiry
- APIGateway
- Cognito (UserPool)
- EBS (Volume)
- EC2 (Instance)
- ECR (Regository)
- ECS (Cluster, Service)
- EFS (FileSystem)
- ELB (LoadBalance)
- RDS (DBInstance)
- SES (Identity)
- SNS (Topic)
- SQS (Queue)

## Development scheduled
Global services such as S3, CloudFont have not yet been implemented due to result display issues...
- S3 (Bucket)
- CloudFront (Distribution)

## How to use?
### 0. Installation Golang
"Go" is basically used, so "Go" must be installed in advance.
[Installation Golang](https://go.dev/doc/install)
### 1. Build
```
go build -o ./bin/<excutable_file>
```
### 2. Output
```
./bin/<excutable_file>
```