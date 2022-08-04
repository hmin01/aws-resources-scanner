package scanner

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type ECSCluster struct {
	Arn                     string       `json:"arn"`
	ContainerInstancesCount uint64       `json:"containerInstanceCount"`
	Name                    string       `json:"name"`
	Services                []ECSService `json:"services"`
}

type ECSService struct {
	Arn        string            `json:"arn"`
	LaunchType string            `json:"launchType"`
	Name       string            `json:"name"`
	Status     string            `json:"status"`
	TasksCount map[string]uint64 `json:"tasksCount"`
}

// ECS 클러스트 목록 조회
func getECSClusters(ctx context.Context, conf aws.Config) []ECSCluster {
	// 클라이언트 생성
	client := ecs.NewFromConfig(conf)

	// 목록 생성
	var list []ECSCluster
	// Pagination
	paginatorForCluster := ecs.NewListClustersPaginator(client, &ecs.ListClustersInput{MaxResults: aws.Int32(100)})

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[ERROR] %v\n", r)
		}
	}()

	// 데이터 조회
	for paginatorForCluster.HasMorePages() {
		resp, err := paginatorForCluster.NextPage(ctx)
		if err != nil {
			panic(err)
		}
		// 클러스터 조회
		clusters, err := client.DescribeClusters(ctx, &ecs.DescribeClustersInput{Clusters: resp.ClusterArns})
		if err != nil {
			panic(err)
		}
		// 데이터 추출
		for _, cluster := range clusters.Clusters {
			// Pagination for service
			paginatorForService := ecs.NewListServicesPaginator(client, &ecs.ListServicesInput{Cluster: cluster.ClusterArn, MaxResults: aws.Int32(10)})
			// 서비스 목록
			var services []ECSService
			// 서비스 조회
			for paginatorForService.HasMorePages() {
				output, err := paginatorForService.NextPage(ctx)
				if err != nil {
					panic(err)
				}
				// 서비스 정보 조회
				if len(output.ServiceArns) > 0 {
					rawServices, err := client.DescribeServices(ctx, &ecs.DescribeServicesInput{Cluster: cluster.ClusterArn, Services: output.ServiceArns})
					if err != nil {
						panic(err)
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
								"pending": uint64(service.PendingCount),
								"running": uint64(service.RunningCount),
							},
						}
						// 목록에 추가
						services = append(services, info)
					}
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
	}
	// 결과 반환
	return list
}
