package modules

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Bucket struct {
	Loction string `json:"location"`
	Name    string `json:"name"`
}

// S3 버킷 조회
func GetS3Buckets(ctx context.Context, cfg aws.Config) []Bucket {
	// 클라이언트 생성
	client := s3.NewFromConfig(cfg)

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[ERROR] %v\n", r)
		}
	}()

	// 데이터 조회
	resp, err := client.ListBuckets(ctx, nil)
	if err != nil {
		panic(err)
	}
	// 데이터 추출
	var list []Bucket
	for _, bucket := range resp.Buckets {
		// 버켓에 대한 리전 확인
		output, err := client.GetBucketLocation(ctx, &s3.GetBucketLocationInput{Bucket: bucket.Name})
		if err != nil {
			panic(err)
		}
		// 버킷 정보 생성
		info := Bucket{
			Loction: string(output.LocationConstraint),
			Name:    *bucket.Name,
		}
		// 목록에 추가
		list = append(list, info)
	}
	// 결과 반환
	return list
}
