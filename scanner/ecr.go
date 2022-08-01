package scanner

import (
	"context"
	"log"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

type ECR struct {
	Arn  string
	Name string
	Uri  string
}

// ECR 저장소 목록 조회
func getECRRepositories(cfg aws.Config) []ECR {
	// Context 생성
	ctx := context.TODO()
	// 클라이언트 생성
	client := ecr.NewFromConfig(cfg)

	// 데이터 조회
	resp, err := client.DescribeRepositories(ctx, nil)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
	// 데이터 추출
	var list []ECR
	for _, repository := range resp.Repositories {
		// 저장소 정보 생성
		info := ECR{
			Arn:  *repository.RepositoryArn,
			Name: *repository.RepositoryName,
			Uri:  *repository.RepositoryUri,
		}
		// 목록에 추가
		list = append(list, info)
	}
	// 결과 반환
	return list
}
