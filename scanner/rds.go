package scanner

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

type DBInstance struct {
	AllocatedStorage uint64 `json:"allocatedStorage"`
	Class            string `json:"class"`
	Engine           string `json:"engine"`
	Id               string `json:"id"`
	MultiAZ          bool   `json:"multiAZ"`
	Name             string `json:"name"`
	Status           string `json:"status"`
	StorageType      string `json:"storageType"`
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
			// 데이터베이스 인스턴스 정보 생성
			info := DBInstance{
				AllocatedStorage: uint64(dbInstance.AllocatedStorage),
				Class:            *dbInstance.DBInstanceClass,
				Engine:           *dbInstance.Engine,
				Id:               *dbInstance.DBInstanceIdentifier,
				MultiAZ:          dbInstance.MultiAZ,
				Name: func(name *string) string {
					if name != nil {
						return *name
					}
					return ""
				}(dbInstance.DBName),
				Status:      *dbInstance.DBInstanceStatus,
				StorageType: *dbInstance.StorageType,
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}
