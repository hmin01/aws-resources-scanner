package scanner

import (
	"context"
	"reflect"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	// Local
	"main.com/scanner/modules"
)

// 리소스
type Resource struct {
	Count uint64 `json:"count"`
	Data  any    `json:"data,omitempty"`
	Type  string `json:"type"`
}

// 스캐너
type Scanner struct {
	ctx       context.Context
	config    aws.Config
	resources chan<- Resource
	r_ops     uint64
	g_ops     uint64
}

// 스캐너 생성 함수
func CreateScanner(ctx context.Context, config aws.Config, resources chan<- Resource) *Scanner {
	return &Scanner{ctx, config, resources, 18, 3}
}

// 리전 종속 서비스 수
func (s *Scanner) R_OPS() uint64 {
	return s.r_ops
}

// 글로벌 서비스 수
func (s *Scanner) G_OPS() uint64 {
	return s.g_ops
}

// API Gateway 조회
func (s *Scanner) GetApiGateways() {
	s.resources <- s.getResources("apigateway", modules.GetApiGateways(s.ctx, s.config))
}

// CloudFront 조회
func (s *Scanner) GetCloudFronts() {
	s.resources <- s.getResources("cloudfront", modules.GetCloudFrontDistributions(s.ctx, s.config))
}

// Cognito 조회
func (s *Scanner) GetCognitos() {
	s.resources <- s.getResources("cognito", modules.GetCognitoUserPools(s.ctx, s.config))
}

// Dynamodb 조회
func (s *Scanner) GetDynamodbs() {
	s.resources <- s.getResources("dynamodb", modules.GetDynamodbTables(s.ctx, s.config))
}

// EBS 조회
func (s *Scanner) GetEBSs() {
	s.resources <- s.getResources("ebs", modules.GetEBSVolumes(s.ctx, s.config))
}

// EC2 조회
func (s *Scanner) GetEC2s() {
	s.resources <- s.getResources("ec2", modules.GetEC2Instances(s.ctx, s.config))
}

// ECR 조회
func (s *Scanner) GetECRs() {
	s.resources <- s.getResources("ecr", modules.GetECRRepositories(s.ctx, s.config))
}

// ECS 조회
func (s *Scanner) GetECSs() {
	s.resources <- s.getResources("ecs", modules.GetECSClusters(s.ctx, s.config))
}

// EFS 조회
func (s *Scanner) GetEFSs() {
	s.resources <- s.getResources("efs", modules.GetEFSStorages(s.ctx, s.config))
}

// Elastic Beanstalk 조회
func (s *Scanner) GetElasticaches() {
	s.resources <- s.getResources("elasticache", modules.GetElasticacheClusters(s.ctx, s.config))
}

// Elastic Beanstalk 조회
func (s *Scanner) GetElasticBeanstalks() {
	s.resources <- s.getResources("elasticbeanstalk", modules.GetElasticBeanstalkApplications(s.ctx, s.config))
}

// ELB 조회
func (s *Scanner) GetELBs() {
	s.resources <- s.getResources("elb", modules.GetLoadBalancers(s.ctx, s.config))
}

// Event Bridge 조회
func (s *Scanner) GetEventBridges() {
	s.resources <- s.getResources("eventbridge", modules.GetEventBridgeRules(s.ctx, s.config))
}

// Lambda 조회
func (s *Scanner) GetLambdas() {
	s.resources <- s.getResources("lambda", modules.GetLambdaFunctions(s.ctx, s.config))
}

// QLDB 조회
func (s *Scanner) GetQLDBs() {
	s.resources <- s.getResources("qldb", modules.GetQLDBLedgers(s.ctx, s.config))
}

// RDS 조회
func (s *Scanner) GetRDSs() {
	s.resources <- s.getResources("rds", modules.GetRDSInstances(s.ctx, s.config))
}

// Route53 조회
func (s *Scanner) GetRoute53s() {
	s.resources <- s.getResources("route53", modules.GetRoute53HostedZones(s.ctx, s.config))
}

// S3 조회
func (s *Scanner) GetS3s() {
	s.resources <- s.getResources("s3", modules.GetS3Buckets(s.ctx, s.config))
}

// SES 조회
func (s *Scanner) GetSESs() {
	s.resources <- s.getResources("ses", modules.GetSESIdentities(s.ctx, s.config))
}

// SNS 조회
func (s *Scanner) GetSNSs() {
	s.resources <- s.getResources("sns", modules.GetSNSTopics(s.ctx, s.config))
}

// SQS 조회
func (s *Scanner) GetSQSs() {
	s.resources <- s.getResources("sqs", modules.GetSQSQueues(s.ctx, s.config))
}

// 응답 데이터
func (s *Scanner) getResources(resourceType string, data any) Resource {
	return Resource{
		Count: uint64(reflect.ValueOf(data).Len()),
		Data:  data,
		Type:  resourceType,
	}
}
