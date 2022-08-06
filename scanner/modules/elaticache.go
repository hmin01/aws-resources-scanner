package modules

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
)

type ReplicationGroup struct {
	Arn              string  `json:"arn"`
	AuthTokenEnabled bool    `json:"authTokenEnabled"`
	ClusterEnabled   bool    `json:"clusterEnabled"`
	Endpoint         string  `json:"endpoint"`
	Engine           string  `json:"engine"`
	Id               string  `json:"id"`
	MultiAZ          string  `json:"multiAZ"`
	NodeType         string  `json:"nodeType"`
	Shards           []Shard `json:"shards"`
	Status           string  `json:"status"`
}

type Shard struct {
	Id     string   `json:"id"`
	Status string   `json:"status"`
	Nodes  []string `json:"nodes"`
}

// Elasticache 리플리카 그룹 조회
func GetElasticacheClusters(ctx context.Context, conf aws.Config) []ReplicationGroup {
	// 클라이언트 생성
	client := elasticache.NewFromConfig(conf)

	// 목록 생성
	var list []ReplicationGroup
	// Paginator 생성
	paginator := elasticache.NewDescribeReplicationGroupsPaginator(client, &elasticache.DescribeReplicationGroupsInput{MaxRecords: aws.Int32(100)})

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
		for _, replicationGroup := range resp.ReplicationGroups {
			// 샤드 정보 생성
			var shards []Shard
			for _, elem := range replicationGroup.NodeGroups {
				// 샤드 정보 생성
				shards = append(shards, Shard{
					Id:     *elem.NodeGroupId,
					Status: *elem.Status,
					Nodes: func(members []types.NodeGroupMember) []string {
						var nodes []string
						for _, member := range members {
							nodes = append(nodes, *member.CacheClusterId)
						}
						return nodes
					}(elem.NodeGroupMembers),
				})
			}
			// Engine 유형 조회를 위해 클러스터 조회
			output, err := client.DescribeCacheClusters(ctx, &elasticache.DescribeCacheClustersInput{CacheClusterId: &replicationGroup.MemberClusters[0]})
			if err != nil {
				panic(err)
			}
			// 리플리카 정보 생성
			info := ReplicationGroup{
				Arn:              *replicationGroup.ARN,
				AuthTokenEnabled: *replicationGroup.AuthTokenEnabled,
				Engine:           *output.CacheClusters[0].Engine,
				Id:               *replicationGroup.ReplicationGroupId,
				MultiAZ:          string(replicationGroup.MultiAZ),
				NodeType:         *replicationGroup.CacheNodeType,
				Status:           *replicationGroup.Status,
				Shards:           shards,
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}
