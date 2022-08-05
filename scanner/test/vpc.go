package test

import (
	"context"
	"fmt"

	// AWS
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// EndPoint
type EndPoint struct {
	Id      string `json:"id"`
	Service string `json:"service"`
	Status  string `json:"status"`
	Type    string `json:"type"`
}

// NAT Gateway
type NatGateway struct {
	ConnectivityType string `json:"connectivityType"`
	Id               string `json:"id"`
	Status           string `json:"status"`
	Subnet           string `json:"subnet"`
	Vpc              string `json:"vpc"`
}

// VPC
type VPC struct {
	CidrBlock string `json:"cidrBlock"`
	Default   bool   `json:"default"`
	Id        string `json:"id"`
	Status    string `json:"status"`
	Tenancy   string `json:"tenancy"`
}

// VPC endpoint 조회
func GetNatGateways(ctx context.Context, conf aws.Config) []NatGateway {
	// 클라이언트 생성
	client := ec2.NewFromConfig(conf)

	// 목록 생성
	var list []NatGateway
	// Paginator 생성
	paginator := ec2.NewDescribeNatGatewaysPaginator(client, &ec2.DescribeNatGatewaysInput{MaxResults: aws.Int32(100)})

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
		for _, natGateway := range resp.NatGateways {
			// Endpoint 정보 생성
			info := NatGateway{
				ConnectivityType: string(natGateway.ConnectivityType),
				Id:               *natGateway.NatGatewayId,
				Status:           string(natGateway.State),
				Subnet:           *natGateway.SubnetId,
				Vpc:              *natGateway.VpcId,
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}

// VPC 조회
func GetVpcs(ctx context.Context, conf aws.Config) []VPC {
	// 클라이언트 생성
	client := ec2.NewFromConfig(conf)

	// 목록 생성
	var list []VPC
	// Paginator 생성
	paginator := ec2.NewDescribeVpcsPaginator(client, &ec2.DescribeVpcsInput{MaxResults: aws.Int32(100)})

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
		for _, vpc := range resp.Vpcs {
			// Endpoint 정보 생성
			info := VPC{
				CidrBlock: *vpc.CidrBlock,
				Default:   *vpc.IsDefault,
				Id:        *vpc.VpcId,
				Status:    string(vpc.State),
				Tenancy:   string(vpc.InstanceTenancy),
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}

// VPC endpoint 조회
func GetVpcEndpoints(ctx context.Context, conf aws.Config) []EndPoint {
	// 클라이언트 생성
	client := ec2.NewFromConfig(conf)

	// 목록 생성
	var list []EndPoint
	// Paginator 생성
	paginator := ec2.NewDescribeVpcEndpointsPaginator(client, &ec2.DescribeVpcEndpointsInput{MaxResults: aws.Int32(100)})

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
		for _, endpoint := range resp.VpcEndpoints {
			// Endpoint 정보 생성
			info := EndPoint{
				Id:      *endpoint.VpcEndpointId,
				Service: *endpoint.ServiceName,
				Status:  string(endpoint.State),
				Type:    string(endpoint.VpcEndpointType),
			}
			// 목록에 추가
			list = append(list, info)
		}
	}
	// 결과 반환
	return list
}
