package scanner

import (
	"context"
	"log"
	"strings"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Queue struct {
	Fifo bool
	Name string
	Url  string
}

func getSQSQueues(cfg aws.Config) []Queue {
	// Context 생성
	ctx := context.TODO()
	// 클라이언트 생성
	client := sqs.NewFromConfig(cfg)

	// 데이터 조회
	resp, err := client.ListQueues(ctx, nil)
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
	// 데이터 추출
	var list []Queue
	for _, queueUrl := range resp.QueueUrls {
		// URL 분리
		splice := strings.Split(queueUrl, "/")
		// 이름 추출
		queueName := splice[len(splice)-1]
		// FIFO 여부 확인
		isFifo := strings.Contains(queueName, ".fifo")

		// 대기열 정보 생성
		info := Queue{
			Fifo: isFifo,
			Name: queueName,
			Url:  queueUrl,
		}
		// 목록에 추가
		list = append(list, info)
	}
	// 결과 반환
	return list
}
