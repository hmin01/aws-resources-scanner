package scanner

import (
	"context"
	"log"

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

func getRDSInstances(cfg aws.Config) []DBInstance {
	// 클라이언트 생성
	client := rds.NewFromConfig(cfg)

	// 데이터 조회
	resp, err := client.DescribeDBInstances(context.TODO(), nil)
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
	// 데이터 추출
	var list []DBInstance
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
	// 결과 반환
	return list
}
