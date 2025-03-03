package dynamodb

import (
	"context"

	"biyard.co/common/base"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	MaximumBatchSize = 25
)

type DynamoDB struct {
	ctx    context.Context
	svc    *dynamodb.DynamoDB
	prefix string
}

func New(ctx context.Context, dbconf base.DatabaseConfig, conf base.AWSConfig) *DynamoDB {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(conf.Region),
		Credentials: credentials.NewStaticCredentials(conf.AccessKeyID, conf.SecretKey, ""),
	})

	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(session)

	return &DynamoDB{
		svc:    svc,
		prefix: dbconf.Name,
	}
}
