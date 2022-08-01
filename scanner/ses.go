package scanner

import (
	"context"
	"log"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

type SES struct {
	Name string
}

func getSESIdentities(cfg aws.Config) []SES {
	// Context 생성
	ctx := context.TODO()
	// 클라이언트 생성
	client := ses.NewFromConfig(cfg)

	// 데이터 조회
	resp, err := client.ListIdentities(ctx, nil)
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
	// 데이터 추출
	var list []SES
	for _, identitie := range resp.Identities {
		// 대기열 정보 생성
		info := SES{
			Name: identitie,
		}
		// 목록에 추가
		list = append(list, info)
	}
	// 결과 반환
	return list
}
