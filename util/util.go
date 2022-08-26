package util

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	_ "os"
	_ "path/filepath"
	_ "runtime"
	"strconv"
	_ "strings"
	_ "time"

	// AWS
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	// Local
	awsConf "main.com/aws"
)

// 응답 헤더
var headers map[string]string = map[string]string{
	"Access-Control-Allow-Headers": "Content-Type",
	"Access-Control-Allow-Methods": "OPTIONS, POST",
	"Access-Control-Allow-Origin":  "*",
}

type ResourceByRegion struct {
	Region    string         `json:"region,omitempty"`
	Resources map[string]any `json:"resources"`
	Usage     bool           `json:"usage"`
}

// 결과 출력
func Print(key string, result map[string]ResourceByRegion) (events.APIGatewayProxyResponse, error) {
	// AWS Configuration
	cfg := awsConf.ConfigurationInternal("ap-northeast-2")
	// 클라이언트 생성
	client := s3.NewFromConfig(cfg)

	// 데이터 변환
	transformed, err := json.Marshal(result)
	if err != nil {
		return Response(500, "[ERROR] The result data format is not valid"), errors.New("the result data format is not valid")
	}

	// Input 생성
	input := &s3.PutObjectInput{
		Body:   bytes.NewReader(transformed),
		Bucket: aws.String("aws-resource-scanner"),
		Key:    aws.String("results/" + key),
	}
	// 업로드를 위한 객체 생성
	uploader := manager.NewUploader(client)
	// 파일 업로드
	_, err = uploader.Upload(context.TODO(), input)
	if err != nil {
		return Response(500, "[ERROR] Failed to upload data to AWS S3"), errors.New("failed to upload data to AWS S3")
	}
	// 반환
	return Response(200, "Success"), nil
}

// 문자열을 정수로 변환
func StringToInteger(text string) int {
	transformed, err := strconv.Atoi(text)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return 0
	}
	// 반환
	return transformed
}

// Response Header
func Response(statusCode int, message string) events.APIGatewayProxyResponse {
	// Map 생성
	data := map[string]string{
		"message": message,
	}
	// 문자열로 변환
	transformed, _ := json.Marshal(data)
	// 반환
	return events.APIGatewayProxyResponse{
		Body:       string(transformed),
		Headers:    headers,
		StatusCode: statusCode,
	}
}

// // 결과 파일 생성
// func createOutput() *os.File {
// 	// 현재 디렉터리 위치 파악 (프로세스 기준)
// 	workDir := extractWorkingDirPath()
// 	// 결과 디렉터리 경로
// 	outPath := filepath.Join(workDir, "../out")
// 	// 디렉터러 존재 여부 확인 및 생성
// 	if _, err := os.Stat(outPath); os.IsExist(err) {
// 		os.Mkdir(outPath, 0744)
// 	}
// 	// 결과 파일 생성
// 	file, err := os.Create(filepath.Join(outPath, "output.json"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	// 파일 반환
// 	return file
// }

// // 프로세스 기준 디렉터리 경로 조회
// func extractWorkingDirPath() string {
// 	// 프로세스 기준 디렉터리 경로 조회
// 	path, err := os.Executable()
// 	if err != nil {
// 		panic(err)
// 	}
// 	// 운영체제에 따른 처리
// 	if runtime.GOOS == "windows" {
// 		// 조회된 경로에서 binary 파일 제외
// 		slice := strings.Split(path, "\\")
// 		// 제외 처리 후, 문자열 병합 및 반환
// 		return strings.Join(slice[:(len(slice)-1)], "\\")
// 	} else {
// 		// 조회된 경로에서 binary 파일 제외
// 		slice := strings.Split(path, "/")
// 		// 제외 처리 후, 문자열 병합 및 반환
// 		return strings.Join(slice[:(len(slice)-1)], "/")
// 	}
// }
