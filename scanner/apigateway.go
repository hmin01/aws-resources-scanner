package scanner

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
)

type ApiGatway struct {
	Id       string `json:"id"`
	IsRest   bool   `json:"isRest"`
	Name     string `json:"name"`
	Protocol string `json:"protocol,omitempty"`
}

func getApiGateways(ctx context.Context, conf aws.Config) []ApiGatway {
	// 클라이언트 생성 (v1 is REST API)
	clientForV1 := apigateway.NewFromConfig(conf)
	// 클라이언트 생성 (v2 is HTTP, WebSocket)
	clientForV2 := apigatewayv2.NewFromConfig(conf)

	// 목록 생성
	var list []ApiGatway
	// Paginator 생성
	paginatorForV1 := apigateway.NewGetRestApisPaginator(clientForV1, &apigateway.GetRestApisInput{Limit: aws.Int32(100)})

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[ERROR] %v\n", r)
		}
	}()

	// 데이터 조회 (REST)
	for paginatorForV1.HasMorePages() {
		resp, err := paginatorForV1.NextPage(ctx)
		if err != nil {
			panic(err)
		}
		// 데이터 추출
		for _, api := range resp.Items {
			// API Gateway 정보 생성
			info := ApiGatway{
				Id:     *api.Id,
				IsRest: true,
				Name:   *api.Name,
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 데이터 조회 (HTTP, WebSocket)
	var nextToken *string
	for {
		resp, err := clientForV2.GetApis(ctx, &apigatewayv2.GetApisInput{NextToken: nextToken})
		if err != nil {
			panic(err)
		}
		// 데이터 추출
		for _, api := range resp.Items {
			// API Gateway 정보 생성
			info := ApiGatway{
				Id:       *api.ApiId,
				IsRest:   false,
				Name:     *api.Name,
				Protocol: string(api.ProtocolType),
			}
			// 목록에 추가
			list = append(list, info)
		}
		// Escape
		if resp.NextToken == nil {
			break
		} else {
			nextToken = resp.NextToken
		}
	}

	// 결과 반환
	return list
}
