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

func getLoadBalancers(cfg aws.Config) []LoadBalancer {
	// 클라이언트 생성
	client := elasticloadbalancingv2.NewFromConfig(cfg)

	// 데이터 조회
	resp, err := client.DescribeLoadBalancers(context.TODO(), nil)
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
	// 데이터 추출
	var list []LoadBalancer
	for _, loadBalancer := range resp.LoadBalancers {
		// 로드 밸런서 정보 생성
		info := LoadBalancer{
			Name: *loadBalancer.LoadBalancerName,
			Type: string(loadBalancer.Type),
		}
		// 목록에 추가
		list = append(list, info)
	}
	// 결과 반환
	return list
}
