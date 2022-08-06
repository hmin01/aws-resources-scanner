package scanner

import (
	"context"
	"fmt"
	"strings"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

type Distribution struct {
	Domain      string   `json:"domain"`
	Enabled     bool     `json:"enabled"`
	HttpVersion string   `json:"httpVersion"`
	Id          string   `json:"id"`
	Origins     []Origin `json:"origins"`
	Status      string   `json:"status"`
}

type Origin struct {
	Domain string `json:"domain"`
	Name   string `json:"name"`
	Id     string `json:"id"`
	Path   string `json:"path"`
}

// CloudFront 배포 조회
func getCloudFrontDistributions(ctx context.Context, cfg aws.Config) []Distribution {
	// 클라이언트 생성
	client := cloudfront.NewFromConfig(cfg)

	// 목록 생성
	var list []Distribution
	// Paginator 생성
	paginator := cloudfront.NewListDistributionsPaginator(client, &cloudfront.ListDistributionsInput{MaxItems: aws.Int32(100)})

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[ERROR] %v\n", r)
		}
	}()

	// 데이터 추출
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			panic(err)
		}
		// 아이템 조회
		for _, distribution := range output.DistributionList.Items {
			// 배포 정보 생성
			info := Distribution{
				Domain:      *distribution.DomainName,
				Enabled:     *distribution.Enabled,
				HttpVersion: string(distribution.HttpVersion),
				Id:          *distribution.Id,
				Origins: func(origins []types.Origin) []Origin {
					var result []Origin
					for _, origin := range origins {
						// Domain에서 S3 오리진 이름 추출
						var name string = *origin.DomainName
						if strings.Contains(*origin.DomainName, ".s3.") {
							name = strings.Split(*origin.DomainName, ".s3.")[0]
						}
						// Origin 정보 생성
						info := Origin{
							Domain: *origin.DomainName,
							Name:   name,
							Id:     *origin.Id,
							Path:   *origin.OriginPath,
						}
						// 목록에 추가
						result = append(result, info)
					}
					return result
				}(distribution.Origins.Items),
				Status: *distribution.Status,
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}
