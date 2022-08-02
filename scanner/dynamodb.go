package scanner

import (
	"context"
	"log"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Table struct {
	Id    string
	Name  string
	State string
}

func getDynamodbTables(ctx context.Context, cfg aws.Config) []Table {
	// 클라이언트 생성
	client := dynamodb.NewFromConfig(cfg)

	// 데이터 조회
	resp, err := client.ListTables(ctx, nil)
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
	// 데이터 추출
	var list []Table
	for _, table := range resp.TableNames {
		result, err := client.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(table)})
		if err != nil {
			log.Fatalf("[ERROR] %s", err)
		}
		// 테이블 정보 생성
		info := Table{
			Id:    *result.Table.TableId,
			Name:  *result.Table.TableName,
			State: string(result.Table.TableStatus),
		}
		// 목록에 추가
		list = append(list, info)
	}
	// 결과 반환
	return list
}
