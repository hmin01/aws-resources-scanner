package scanner

import (
	"context"
	"log"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// AWS SDK를 위한 설정 함수
func Configuration(region string) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// 설정 반환
	return cfg
}

// 사용 가능한 리전 조회
func GetRegions(config aws.Config) []string {
	// 클라이언트 생성
	client := ec2.NewFromConfig(config)
	// SDK 호출
	resp, err := client.DescribeRegions(context.TODO(), nil)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
	// 데이터 처리
	var list []string
	for _, output := range resp.Regions {
		list = append(list, *output.RegionName)
	}
	// 결과 반환
	return list
}
