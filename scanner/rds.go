package scanner

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

type DBInstance struct {
	Class       string
	Id          string
	Name        string
	State       string
	StorageType string
}

func getRDSInstances(ctx context.Context, conf aws.Config) []DBInstance {
	// 클라이언트 생성
	client := rds.NewFromConfig(conf)

	// 목록 생성
	var list []DBInstance
	// Paginator 생성
	paginator := rds.NewDescribeDBInstancesPaginator(client, &rds.DescribeDBInstancesInput{MaxRecords: aws.Int32(100)})

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
		for _, dbInstance := range resp.DBInstances {
			// 로드 밸런서 정보 생성
			info := DBInstance{
				Class: *dbInstance.DBInstanceClass,
				Id:    *dbInstance.DBInstanceIdentifier,
				Name: func(name *string) string {
					if name != nil {
						return *name
					}
					return ""
				}(dbInstance.DBName),
				State:       *dbInstance.DBInstanceStatus,
				StorageType: *dbInstance.StorageType,
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}
