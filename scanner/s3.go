package scanner

import (
	"context"
	"log"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Bucket struct {
	Name string
}

func getS3Buckets(ctx context.Context, cfg aws.Config) []Bucket {
	// 클라이언트 생성
	client := s3.NewFromConfig(cfg)

	// 데이터 조회
	resp, err := client.ListBuckets(ctx, nil)
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
	// 데이터 추출
	var list []Bucket
	for _, bucket := range resp.Buckets {
		// 로드 밸런서 정보 생성
		info := Bucket{
			Name: *bucket.Name,
		}
		// 목록에 추가
		list = append(list, info)
	}
	// 결과 반환
	return list
}
