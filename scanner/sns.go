package scanner

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"

	// Local
	"main.com/util"
)

type SNS struct {
	Arn          string
	Fifo         bool
	Name         string
	Subscription map[string]uint64
}

func getSNSTopics(ctx context.Context, conf aws.Config) []SNS {
	// 클라이언트 생성
	client := sns.NewFromConfig(conf)

	// 목록 생성
	var list []SNS
	// Pagination
	paginator := sns.NewListTopicsPaginator(client, nil)

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
		for _, topic := range resp.Topics {
			// 속성 조회
			attributes, err := client.GetTopicAttributes(ctx, &sns.GetTopicAttributesInput{TopicArn: topic.TopicArn})
			if err != nil {
				panic(err)
			}
			// 대기열 정보 생성
			info := SNS{
				Arn:  *topic.TopicArn,
				Fifo: func(status string) bool { return status == "true" }(attributes.Attributes["FifoTopic"]),
				Name: attributes.Attributes["DisplayName"],
				Subscription: map[string]uint64{
					"Confirmed": uint64(util.StringToInteger(attributes.Attributes["SubscriptionsConfirmed"])),
					"Pending":   uint64(util.StringToInteger(attributes.Attributes["SubscriptionsPending"])),
				},
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}
