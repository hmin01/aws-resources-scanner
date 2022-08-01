package scanner

import (
	"context"
	"log"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// EC2 인스턴스에 대한 정보
type Instance struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	State string `json:"state"`
}

// EBS 볼륨에 대한 정보
type Volume struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Size  int32  `json:"size"`
	State string `json:"state"`
}

// EC2 인스턴스 목록 조회
func getEC2Instances(cfg aws.Config) []Instance {
	// EC2 클라이언트 생성
	client := ec2.NewFromConfig(cfg)

	// 데이터 조회
	resp, err := client.DescribeInstances(context.TODO(), nil)
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
	// 데이터 추출
	var list []Instance
	for _, elem := range resp.Reservations {
		for _, instance := range elem.Instances {
			// 인스턴스 정보 생성
			info := Instance{
				Id:    *instance.InstanceId,
				Type:  string(instance.InstanceType),
				State: string(instance.State.Name),
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}

// EBS 볼륨 목록 조회
func getEBSVolumes(cfg aws.Config) []Volume {
	// EC2 클라이언트 생성
	client := ec2.NewFromConfig(cfg)

	// 데이터 조회
	resp, err := client.DescribeVolumes(context.TODO(), nil)
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
	// 데이터 추출
	var list []Volume
	for _, volume := range resp.Volumes {
		info := Volume{
			// 볼륨 정보 생성
			Id:    *volume.VolumeId,
			Type:  string(volume.VolumeType),
			Size:  *volume.Size,
			State: string(volume.State),
		}
		// 목록에 추가
		list = append(list, info)
	}
	// 결과 반환
	return list
}
