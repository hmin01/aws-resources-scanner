package modules

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

type LoadBalancer struct {
	Arn    string `json:"arn"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

func GetLoadBalancers(ctx context.Context, conf aws.Config) []LoadBalancer {
	// 클라이언트 생성
	client := elasticloadbalancingv2.NewFromConfig(conf)

	// 목록 생성
	var list []LoadBalancer
	// Paginator 생성
	paginator := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(client, &elasticloadbalancingv2.DescribeLoadBalancersInput{PageSize: aws.Int32(100)})

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
		for _, loadBalancer := range resp.LoadBalancers {
			// 로드 밸런서 정보 생성
			info := LoadBalancer{
				Arn:    *loadBalancer.LoadBalancerArn,
				Name:   *loadBalancer.LoadBalancerName,
				Status: string(loadBalancer.State.Code),
				Type:   string(loadBalancer.Type),
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}
