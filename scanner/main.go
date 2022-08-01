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

// EC2 조회
func GetEC2s(config aws.Config, resources chan<- Resource) {
	resources <- getResources("ec2", getEC2Instances(config))
}

// EBS 조회
func GetEBSs(config aws.Config, resources chan<- Resource) {
	resources <- getResources("ebs", getEBSVolumes(config))
}

// 응답 데이터
func getResources(resourceType string, data any) Resource {
	return Resource{
		Count: uint64(reflect.ValueOf(data).Len()),
		Data:  data,
		Type:  resourceType,
	}
}
