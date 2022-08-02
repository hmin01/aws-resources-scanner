package scanner

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoUserPool struct {
	Id     string
	Name   string
	Status string
}

func getCognitoUserPools(ctx context.Context, conf aws.Config) []CognitoUserPool {
	// 클라이언트 생성
	client := cognitoidentityprovider.NewFromConfig(conf)

	// 목록 생성
	var list []CognitoUserPool
	// Paginator 생성
	paginator := cognitoidentityprovider.NewListUserPoolsPaginator(client, &cognitoidentityprovider.ListUserPoolsInput{MaxResults: int32(10)})

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
		for _, userPool := range resp.UserPools {
			// 사용자 풀 정보 생성
			info := CognitoUserPool{
				Id:     *userPool.Id,
				Name:   *userPool.Name,
				Status: string(userPool.Status),
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}
