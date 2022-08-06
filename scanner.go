package main

import (
	"context"
	"fmt"
	"log"
	"time"

	// Custom
	scan "main.com/scanner"
	"main.com/util"
)

const DEFAULT_REGION string = "ap-northeast-2"

func main() {
	// Init (커맨드라인에서 Argument 가져오기)
	scan.Init()

	// 작업 결과 채널 생성
	integrations := make(chan util.ResourceByRegion, 10)
	// 조회 최종 결과
	result := make(map[string]util.ResourceByRegion)
	// 작업 카운트
	var ops int = 0

	config := scan.Configuration(DEFAULT_REGION)
	// 사용 가능한 리전 조회
	regions := scan.GetRegions(config)
	// regions := []string{"ap-northeast-2"}
	// 사용 가능 리전이 없을 경우, 종료
	if len(regions) == 0 {
		log.Fatal("[NOTICE] 사용 가능 리전이 없습니다.")
	} else {
		fmt.Println("=-=-=-=- 작업을 진행할 리전 목록 -=-=-=-=")
	}

	// 조회 처리 시간 (Start)
	start := time.Now()
	// 리전별 리소스 조회 (병렬)
	var stmt string = ""
	for index, region := range regions {
		stmt += region
		if index != len(regions)-1 {
			stmt += ", "
		}
		// 작업 수행
		go ScanResources(region, integrations)
	}
	// 글로벌 리소스 조회 (병렬)
	go ScanGlobalResources(integrations)
	// 로그 출력
	fmt.Printf("%s\n\n", stmt)

	// 리소스 통합 및 결과 출력
	for integration := range integrations {
		// 리전별 리소스 통합
		result[integration.Region] = util.ResourceByRegion{
			Resources: integration.Resources,
			Usage:     integration.Usage,
		}
		// 작업 완료 카운트
		ops += 1
		// 모든 작업 완료 여부 확인
		if ops == len(regions) {
			// 조회 처리 시간 출력
			fmt.Printf("\n[NOTICE] Query process duration: %v\n", time.Since(start))
			// 채널 종료
			close(integrations)
			// 결과 출력
			util.Print(result)
		}
	}
}

// 리전에 따른 사용 중인 리소스 조회
func ScanResources(region string, result chan<- util.ResourceByRegion) {
	// Context 생성
	ctx := context.TODO()
	// 데이터 처리를 위한 채널 생성
	resources := make(chan scan.Resource, 20)
	// 스캔을 위한 AWS 설정
	config := scan.Configuration(region)
	// 통합된 리소스 데이터
	integration := make(map[string]any)
	// 작업 카운트
	var ops uint64 = 0
	// 해당 리전 사용 여부 (리소스 존재 여부)
	var usage bool = false

	// 스캐너 생성
	scanner := scan.CreateScanner(ctx, config, resources)
	// 각 리소스 조회
	go scanner.GetApiGateways()
	go scanner.GetCognitos()
	go scanner.GetDynamodbs()
	go scanner.GetEBSs()
	go scanner.GetEC2s()
	go scanner.GetECRs()
	go scanner.GetECSs()
	go scanner.GetEFSs()
	go scanner.GetElasticaches()
	go scanner.GetElasticBeanstalks()
	go scanner.GetELBs()
	go scanner.GetEventBridges()
	go scanner.GetLambdas()
	go scanner.GetQLDBs()
	go scanner.GetRDSs()
	go scanner.GetSESs()
	go scanner.GetSNSs()
	go scanner.GetSQSs()

	for resource := range resources {
		// 리소스 통합
		integration[resource.Type] = resource.Data
		// 리소스 존재 여부
		if !usage && resource.Count > 0 {
			usage = true
		}
		// 작업 완료 카운트
		ops += 1
		// 모든 작업 완료 여부 확인
		if ops == scanner.R_OPS() {
			// 채널 종료
			close(resources)
			// 완료된 데이터를 채널로 전송
			result <- util.ResourceByRegion{
				Region:    region,
				Resources: integration,
				Usage:     usage,
			}
			// Log
			fmt.Printf("[NOTICE] %s 에 대한 리소스 조회 완료\n", region)
		}
	}
}

// 글로벌 서비스 조회
func ScanGlobalResources(result chan<- util.ResourceByRegion) {
	// Context 생성
	ctx := context.TODO()
	// 데이터 처리를 위한 채널 생성
	resources := make(chan scan.Resource, 2)
	// 스캔을 위한 AWS 설정
	config := scan.Configuration(DEFAULT_REGION)
	// 통합된 리소스 데이터
	integration := make(map[string]any)
	// 작업 카운트
	var ops uint64 = 0
	// 글로벌 서비스 사용 여부 (리소스 존재 여부)
	var usage bool = false

	// 스캐너 생성
	scanner := scan.CreateScanner(ctx, config, resources)
	// 각 리소스 조회
	go scanner.GetCloudFronts()
	go scanner.GetRoute53s()
	go scanner.GetS3s()

	for resource := range resources {
		// 리소스 통합
		integration[resource.Type] = resource.Data
		// 작업 완료 카운트
		ops += 1
		// 리소스 존재 여부
		if !usage && resource.Count > 0 {
			usage = true
		}
		// 모든 작업 완료 여부 확인
		if ops == scanner.G_OPS() {
			// 채널 종료
			close(resources)
			// 완료된 데이터를 채널로 전송
			result <- util.ResourceByRegion{
				Region:    "global",
				Resources: integration,
				Usage:     usage,
			}
			// Log
			fmt.Println("[NOTICE] global 에 대한 리소스 조회 완료")
		}
	}
}
