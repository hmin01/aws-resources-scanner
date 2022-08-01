package scanner

import (
	"context"
	"log"

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

func getSNSTopics(cfg aws.Config) []SNS {
	// Context 생성
	ctx := context.TODO()
	// 클라이언트 생성
	client := sns.NewFromConfig(cfg)

	// 데이터 조회
	resp, err := client.ListTopics(ctx, nil)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
	// 데이터 추출
	var list []SNS
	for _, topic := range resp.Topics {
		// 속성 조회
		attributes, err := client.GetTopicAttributes(ctx, &sns.GetTopicAttributesInput{TopicArn: topic.TopicArn})
		if err != nil {
			log.Fatalf("[ERROR] %v", err)
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
	// 결과 반환
	return list
}
