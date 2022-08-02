package scanner

import (
	"context"
	"log"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

type LoadBalancer struct {
	Name string
	Type string
}

func getLoadBalancers(ctx context.Context, conf aws.Config) []LoadBalancer {
	// 클라이언트 생성
	client := elasticloadbalancingv2.NewFromConfig(conf)

	// 목록 생성
	var list []LoadBalancer
	// Paginator 생성
	paginator := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(client, &elasticloadbalancingv2.DescribeLoadBalancersInput{PageSize: aws.Int32(100)})

	// 데이터 조회
	for paginator.HasMorePages() {
		resp, err := paginator.NextPage(ctx)
		if err != nil {
			log.Fatalf("[ERROR] %s", err)
		}
		// 데이터 추출
		for _, loadBalancer := range resp.LoadBalancers {
			// 로드 밸런서 정보 생성
			info := LoadBalancer{
				Name: *loadBalancer.LoadBalancerName,
				Type: string(loadBalancer.Type),
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}
