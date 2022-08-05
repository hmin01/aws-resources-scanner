package modules

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

type HostedZone struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetRoute53HostedZones(ctx context.Context, conf aws.Config) []HostedZone {
	// 클라이언트 생성
	client := route53.NewFromConfig(conf)

	// 목록 생성
	var list []HostedZone
	// Paginator 생성
	paginator := route53.NewListHostedZonesPaginator(client, &route53.ListHostedZonesInput{MaxItems: aws.Int32(10)})

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
		for _, hostedZone := range resp.HostedZones {
			// 사용자 풀 정보 생성
			info := HostedZone{
				Id:   *hostedZone.Id,
				Name: *hostedZone.Name,
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}
