package modules

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

type Application struct {
	Arn          string        `json:"arn"`
	Name         string        `json:"name"`
	Environments []Environment `json:"environments"`
}

type Environment struct {
	Arn    string `json:"arn"`
	Health string `json:"health"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// Elastic Beanstalk 애플리케이션 조회
func GetElasticBeanstalkApplications(ctx context.Context, conf aws.Config) []Application {
	// 클라이언트 생성
	client := elasticbeanstalk.NewFromConfig(conf)
	// Next token
	var nextToken *string

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[ERROR] %v\n", r)
		}
	}()

	// 데이터 조회
	var list []Application
	resp, err := client.DescribeApplications(ctx, nil)
	if err != nil {
		panic(err)
	}
	// 데이터 추출
	for _, application := range resp.Applications {
		// 목록 생성
		var environments []Environment
		// 데이터 추출
	LOOP:
		for {
			// Param
			input := &elasticbeanstalk.DescribeEnvironmentsInput{
				ApplicationName: application.ApplicationName,
				MaxRecords:      aws.Int32(100),
				NextToken:       nextToken,
			}
			// 환경 조회
			output, err := client.DescribeEnvironments(ctx, input)
			if err != nil {
				panic(err)
			}
			// 데이터 추출
			for _, environment := range output.Environments {
				// 환ㄴ경 정보 생성
				info := Environment{
					Arn:    *environment.EnvironmentArn,
					Health: string(environment.HealthStatus),
					Id:     *environment.EnvironmentId,
					Name:   *environment.EnvironmentName,
					Status: string(environment.Status),
				}
				// 목록에 추가
				environments = append(environments, info)
			}
			// Escape
			if output.NextToken == nil {
				break LOOP
			} else {
				nextToken = output.NextToken
			}
		}
		// 애플리케이션 정보 생성
		info := Application{
			Arn:          *application.ApplicationArn,
			Name:         *application.ApplicationName,
			Environments: environments,
		}
		// 목록에 추가
		list = append(list, info)
	}
	// 결과 반환
	return list
}
