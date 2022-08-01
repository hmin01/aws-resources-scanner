package scanner

import (
	"context"
	"log"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type ECSCluster struct {
	Arn                     string
	ContainerInstancesCount uint64
	Name                    string
	Services                []ECSService
}

type ECSService struct {
	Arn        string
	LaunchType string
	Name       string
	Status     string
	TasksCount map[string]uint64
}

// ECS 클러스트 목록 조회
func getECSClusters(cfg aws.Config) []ECSCluster {
	// Context 생성
	ctx := context.TODO()
	// 클라이언트 생성
	client := ecs.NewFromConfig(cfg)

	// 데이터 조회
	resp, err := client.ListClusters(ctx, nil)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
	// 클러스터 조회
	clusters, err := client.DescribeClusters(ctx, &ecs.DescribeClustersInput{Clusters: resp.ClusterArns})
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
	// 데이터 추출
	var list []ECSCluster
	for _, cluster := range clusters.Clusters {
		// 클러스터 내 서비스 목록 조회
		output, err := client.ListServices(ctx, &ecs.ListServicesInput{Cluster: cluster.ClusterArn})
		if err != nil {
			log.Fatalf("[ERROR] %v", err)
		}
		// 서비스 조회
		var services []ECSService
		if len(output.ServiceArns) > 0 {
			rawServices, err := client.DescribeServices(ctx, &ecs.DescribeServicesInput{Cluster: cluster.ClusterArn, Services: output.ServiceArns})
			if err != nil {
				log.Fatalf("[ERROR] %v", err)
			}
			// 서비스 추출
			for _, service := range rawServices.Services {
				// 서비스 정보 생성
				info := ECSService{
					Arn:        *service.ServiceArn,
					LaunchType: string(service.LaunchType),
					Name:       *service.ServiceName,
					Status:     *service.Status,
					TasksCount: map[string]uint64{
						"Pending": uint64(service.PendingCount),
						"Running": uint64(service.RunningCount),
					},
				}
				// 목록에 추가
				services = append(services, info)
			}
		}
		// 저장소 정보 생성
		info := ECSCluster{
			Arn:                     *cluster.ClusterArn,
			ContainerInstancesCount: uint64(cluster.RegisteredContainerInstancesCount),
			Name:                    *cluster.ClusterName,
			Services:                services,
		}
		// 목록에 추가
		list = append(list, info)
	}
	// 결과 반환
	return list
}
