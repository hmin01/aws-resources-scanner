package scanner

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

type ECR struct {
	Arn  string `json:"arn"`
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

// ECR 저장소 목록 조회
func getECRRepositories(ctx context.Context, conf aws.Config) []ECR {
	// 클라이언트 생성
	client := ecr.NewFromConfig(conf)

	// 목록 생성
	var list []ECR
	// Paginator 생성
	paginator := ecr.NewDescribeRepositoriesPaginator(client, &ecr.DescribeRepositoriesInput{MaxResults: aws.Int32(100)})

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[ERROR] %v\n", r)
		}
	}()

	// 데이터 조회
	for paginator.HasMorePages() {
		resp, err := paginator.NextPage(ctx)
		if err != nil {
			panic(err)
		}
		// 데이터 추출
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
	}
	// 결과 반환
	return list
}
