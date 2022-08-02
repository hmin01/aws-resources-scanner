package scanner

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type LambdaFunction struct {
	Arn     string
	Memory  uint64
	Name    string
	Runtime string
	Timeout uint64
}

func getLambdaFunctions(ctx context.Context, conf aws.Config) []LambdaFunction {
	// 클라이언트 생성
	client := lambda.NewFromConfig(conf)

	// 목록 생성
	var list []LambdaFunction
	// Paginator 생성
	paginator := lambda.NewListFunctionsPaginator(client, &lambda.ListFunctionsInput{MaxItems: aws.Int32(100)})

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
		for _, function := range resp.Functions {
			// 함수 정보 생성
			info := LambdaFunction{
				Arn:     *function.FunctionArn,
				Memory:  uint64(*function.MemorySize),
				Name:    *function.FunctionName,
				Runtime: string(function.Runtime),
				Timeout: uint64(*function.Timeout),
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}
