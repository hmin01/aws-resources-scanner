module aws-resource-scanner

go 1.18

require (
	main.com/scanner v0.0.0
	main.com/util v0.0.0
)

require (
	github.com/aws/aws-sdk-go-v2 v1.16.8 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.3 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.15.14 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.12.9 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.9 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.15.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.12.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.18.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.17.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.15.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.51.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecr v1.17.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecs v1.18.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/efs v1.17.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.18.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/lambda v1.23.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/rds v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.13.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/sns v1.17.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.19.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.9 // indirect
	github.com/aws/smithy-go v1.12.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

replace (
	main.com/scanner => ./scanner
	main.com/util => ./util
)
