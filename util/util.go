package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type ResourceByRegion struct {
	Region    string         `json:"region,omitempty"`
	Resources map[string]any `json:"resources"`
	Usage     bool           `json:"usage"`
}

// 결과 출력
func Print(result map[string]ResourceByRegion) {
	// 작업 처리 시간 (Start)
	start := time.Now()
	// 파일 생성
	file := createOut()
	defer file.Close()
	// 데이터 변환
	transformed, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	// 파일 쓰기
	_, err = file.Write(transformed)
	if err != nil {
		panic(err)
	}
	// 작업 처리 시간 출력
	fmt.Printf("[NOTICE] Print process duration: %v\n", time.Since(start))
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

// 결과 파일 생성
func createOut() *os.File {
	// 현재 디렉터리 위치 파악 (프로세스 기준)
	workDir := extractWorkingDirPath()
	// 결과 디렉터리 경로
	outPath := filepath.Join(workDir, "../out")
	// 디렉터러 존재 여부 확인 및 생성
	if _, err := os.Stat(outPath); os.IsExist(err) {
		os.Mkdir(outPath, 0744)
	}
	// 결과 파일 생성
	file, err := os.Create(filepath.Join(outPath, "output.json"))
	if err != nil {
		panic(err)
	}
	// 파일 반환
	return file
}

// 프로세스 기준 디렉터리 경로 조회
func extractWorkingDirPath() string {
	// 프로세스 기준 디렉터리 경로 조회
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	// 운영체제에 따른 처리
	if runtime.GOOS == "windows" {
		// 조회된 경로에서 binary 파일 제외
		slice := strings.Split(path, "\\")
		// 제외 처리 후, 문자열 병합 및 반환
		return strings.Join(slice[:(len(slice)-1)], "\\")
	} else {
		// 조회된 경로에서 binary 파일 제외
		slice := strings.Split(path, "/")
		// 제외 처리 후, 문자열 병합 및 반환
		return strings.Join(slice[:(len(slice)-1)], "/")
	}
}
