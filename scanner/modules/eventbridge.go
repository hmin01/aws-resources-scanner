package modules

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

type Rule struct {
	Arn      string `json:"arn"`
	EventBus string `json:"eventBus"`
	Name     string `json:"name"`
	Status   string `json:"status"`
}

// Elastic Beanstalk 애플리케이션 조회
func GetEventBridgeRules(ctx context.Context, conf aws.Config) []Rule {
	// 클라이언트 생성
	client := eventbridge.NewFromConfig(conf)
	// Next token
	var nextToken *string
	// 목록 생성
	var list []Rule

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[ERROR] (for event bridge) %v\n", r)
		}
	}()

LOOP:
	for {
		// Param
		input := &eventbridge.ListRulesInput{
			Limit:     aws.Int32(100),
			NextToken: nextToken,
		}
		// 데이터 조회
		output, err := client.ListRules(ctx, input)
		if err != nil {
			panic(err)
		}
		// 데이터 추출
		for _, rule := range output.Rules {
			// 환경 정보 생성
			info := Rule{
				Arn:      *rule.Arn,
				EventBus: *rule.EventBusName,
				Name:     *rule.Name,
				Status:   string(rule.State),
			}
			// 목록에 추가
			list = append(list, info)
		}
		// Escape
		if output.NextToken == nil {
			break LOOP
		} else {
			nextToken = output.NextToken
		}
	}
	// 결과 반환
	return list
}
