module github.com/aws/aws-sdk-go-v2/service/neptune

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.7.1
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.2.1
	github.com/aws/smithy-go v1.6.1-0.20210719175327-4970553d9934
	github.com/jmespath/go-jmespath v0.4.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../service/internal/presigned-url/
