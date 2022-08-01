package scanner

import (
	"context"
	"log"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
)

type FileSystem struct {
	Id   string
	Name string
	Size int64
}

func getEFSStorages(cfg aws.Config) []FileSystem {
	// 클라이언트 생성
	client := efs.NewFromConfig(cfg)

	// 데이터 조회
	resp, err := client.DescribeFileSystems(context.TODO(), nil)
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
	// 데이터 추출
	var list []FileSystem
	for _, fileSystem := range resp.FileSystems {
		// 파일 시스템 정보 생성
		info := FileSystem{
			Id:   *fileSystem.FileSystemId,
			Name: *fileSystem.Name,
			Size: fileSystem.SizeInBytes.Value,
		}
		// 목록에 추가
		list = append(list, info)
	}
	// 결과 반환
	return list
}
