package modules

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
)

type Ledger struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func GetQLDBLedgers(ctx context.Context, conf aws.Config) []Ledger {
	// 클라이언트 생성
	client := qldb.NewFromConfig(conf)

	// 목록 생성
	var list []Ledger
	// Paginator 생성
	paginator := qldb.NewListLedgersPaginator(client, &qldb.ListLedgersInput{MaxResults: aws.Int32(100)})

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
		for _, ledger := range resp.Ledgers {
			// 사용자 풀 정보 생성
			info := Ledger{
				Name:   *ledger.Name,
				Status: string(ledger.State),
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}
