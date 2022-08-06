module modules

go 1.18

require (
	github.com/aws/aws-sdk-go-v2 v1.16.8
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.15.11
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.12.9
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.18.5
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.18.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.15.10
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.51.1
	github.com/aws/aws-sdk-go-v2/service/ecr v1.17.9
	github.com/aws/aws-sdk-go-v2/service/ecs v1.18.12
	github.com/aws/aws-sdk-go-v2/service/efs v1.17.7
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.22.1
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.14.10
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.18.9
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.16.6
	github.com/aws/aws-sdk-go-v2/service/lambda v1.23.5
	github.com/aws/aws-sdk-go-v2/service/qldb v1.14.10
	github.com/aws/aws-sdk-go-v2/service/rds v1.23.2
	github.com/aws/aws-sdk-go-v2/service/route53 v1.21.4
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.2
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.13.10
	github.com/aws/aws-sdk-go-v2/service/sns v1.17.10
	github.com/aws/aws-sdk-go-v2/service/sqs v1.19.1
	main.com/util v0.0.0
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.9 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.9 // indirect
	github.com/aws/smithy-go v1.12.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

replace main.com/util => ../../util
