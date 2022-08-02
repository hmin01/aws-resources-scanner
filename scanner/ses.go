package scanner

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

type SES struct {
	Name string
	Type string
}

func getSESIdentities(ctx context.Context, conf aws.Config) []SES {
	// 클라이언트 생성
	client := sesv2.NewFromConfig(conf)

	// 목록 생성
	var list []SES
	// Pagination
	paginator := sesv2.NewListEmailIdentitiesPaginator(client, &sesv2.ListEmailIdentitiesInput{PageSize: aws.Int32(1000)})

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
		for _, identity := range resp.EmailIdentities {
			// SES 정보 생성
			info := SES{
				Name: *identity.IdentityName,
				Type: string(identity.IdentityType),
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}
