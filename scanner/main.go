package scanner

import (
	"context"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// 리소스
type Resource struct {
	Count uint64 `json:"count"`
	Data  any    `json:"data,omitempty"`
	Type  string `json:"type"`
}

// API Gateway 조회
func GetApiGateways(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("apigateway", getApiGateways(ctx, config))
}

// Cognito 조회
func GetCognitos(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("apigateway", getCognitoUserPools(ctx, config))
}

// Dynamodb 조회
func GetDynamodbs(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("dynamodb", getDynamodbTables(ctx, config))
}

// EBS 조회
func GetEBSs(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("ebs", getEBSVolumes(ctx, config))
}

// EC2 조회
func GetEC2s(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("ec2", getEC2Instances(ctx, config))
}

// ECR 조회
func GetECRs(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("ecr", getECRRepositories(ctx, config))
}

// ECS 조회
func GetECSs(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("ecs", getECSClusters(ctx, config))
}

// EFS 조회
func GetEFSs(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("efs", getEFSStorages(ctx, config))
}

// ELB 조회
func GetELBs(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("elb", getLoadBalancers(ctx, config))
}

// Lambda 조회
func GetLambdas(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("lambda", getLambdaFunctions(ctx, config))
}

// RDS 조회
func GetRDSs(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("rds", getRDSInstances(ctx, config))
}

// S3 조회
func GetS3s(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("s3", getS3Buckets(ctx, config))
}

// SES 조회
func GetSESs(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("ses", getSESIdentities(ctx, config))
}

// SNS 조회
func GetSNSs(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("sns", getSNSTopics(ctx, config))
}

// SQS 조회
func GetSQSs(ctx context.Context, config aws.Config, resources chan<- Resource) {
	resources <- getResources("sqs", getSQSQueues(ctx, config))
}

// 응답 데이터
func getResources(resourceType string, data any) Resource {
	return Resource{
		Count: uint64(reflect.ValueOf(data).Len()),
		Data:  data,
		Type:  resourceType,
	}
}
