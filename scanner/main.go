package scanner

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// 리소스
type Resource struct {
	Count uint64 `json:"count"`
	Data  any    `json:"data,omitempty"`
	Type  string `json:"type"`
}

// Dynamodb 조회
func GetDynamodbs(config aws.Config, resources chan<- Resource) {
	resources <- getResources("dynamodb", getDynamodbTables(config))
}

// EBS 조회
func GetEBSs(config aws.Config, resources chan<- Resource) {
	resources <- getResources("ebs", getEBSVolumes(config))
}

// EC2 조회
func GetEC2s(config aws.Config, resources chan<- Resource) {
	resources <- getResources("ec2", getEC2Instances(config))
}

// ECR 조회
func GetECRs(config aws.Config, resources chan<- Resource) {
	resources <- getResources("ecr", getECRRepositories(config))
}

// ECS 조회
func GetECSs(config aws.Config, resources chan<- Resource) {
	resources <- getResources("ecs", getECSClusters(config))
}

// EFS 조회
func GetEFSs(config aws.Config, resources chan<- Resource) {
	resources <- getResources("efs", getEFSStorages(config))
}

// ELB 조회
func GetELBs(config aws.Config, resources chan<- Resource) {
	resources <- getResources("elb", getLoadBalancers(config))
}

// RDS 조회
func GetRDSs(config aws.Config, resources chan<- Resource) {
	resources <- getResources("rds", getRDSInstances(config))
}

// S3 조회
func GetS3s(config aws.Config, resources chan<- Resource) {
	resources <- getResources("s3", getS3Buckets(config))
}

// SES 조회
func GetSESs(config aws.Config, resources chan<- Resource) {
	resources <- getResources("ses", getSESIdentities(config))
}

// SNS 조회
func GetSNSs(config aws.Config, resources chan<- Resource) {
	resources <- getResources("sns", getSNSTopics(config))
}

// SQS 조회
func GetSQSs(config aws.Config, resources chan<- Resource) {
	resources <- getResources("sqs", getSQSQueues(config))
}

// 응답 데이터
func getResources(resourceType string, data any) Resource {
	return Resource{
		Count: uint64(reflect.ValueOf(data).Len()),
		Data:  data,
		Type:  resourceType,
	}
}
