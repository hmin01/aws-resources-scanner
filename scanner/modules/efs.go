package modules

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
)

type FileSystem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

func GetEFSStorages(ctx context.Context, conf aws.Config) []FileSystem {
	// 클라이언트 생성
	client := efs.NewFromConfig(conf)

	// 목록 생성
	var list []FileSystem
	// Paginator 생성
	paginator := efs.NewDescribeFileSystemsPaginator(client, &efs.DescribeFileSystemsInput{MaxItems: aws.Int32(100)})

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
	}
	// 결과 반환
	return list
}
