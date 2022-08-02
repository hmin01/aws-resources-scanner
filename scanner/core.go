package scanner

import (
	"context"
	"log"
	"os"
	"strings"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

var RoleArn string = ""

func Init() {
	// 커맨드라인 Argument 확인
	if len(os.Args) == 1 {
		log.Fatalln("[COMMAND ERROR] AWS IAM Role에 대한 ARN은 필수입니다.")
	} else if len(os.Args) > 2 {
		log.Fatalln("[COMMAND ERROR] Argument가 너무 많습니다.")
	}
	// 커맨드라인 Argument에서 Role arn 가져오기
	RoleArn = strings.Join(os.Args[1:2], "")
}

// AWS SDK를 위한 설정 함수
func Configuration(region string) aws.Config {
	// Context 생성
	ctx := context.TODO()
	// AWS Configuration
	conf, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		log.Fatalf("[CONFIG ERROR], %v", err)
	}

	// STS 클라이언트 생성
	client := sts.NewFromConfig(conf)
	// Credentials 생성
	credentials := stscreds.NewAssumeRoleProvider(client, RoleArn)
	// Configuration에 Credentials 추가
	conf.Credentials = aws.NewCredentialsCache(credentials)
	// 설정 반환
	return conf
}

// 사용 가능한 리전 조회
func GetRegions(config aws.Config) []string {
	// 클라이언트 생성
	client := ec2.NewFromConfig(config)
	// SDK 호출
	resp, err := client.DescribeRegions(context.TODO(), nil)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
	// 데이터 처리
	var list []string
	for _, output := range resp.Regions {
		list = append(list, *output.RegionName)
	}
	// 결과 반환
	return list
}
