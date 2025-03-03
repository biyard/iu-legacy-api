package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path"
	"time"

	"biyard.co/common/base"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Bucket struct {
	ctx      context.Context
	svc      *s3.S3
	uploader *s3manager.Uploader
	region   *string
}

func New(ctx context.Context, dbconf base.DatabaseConfig, conf base.AWSConfig) *S3Bucket {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(conf.Region),
		Credentials: credentials.NewStaticCredentials(conf.AccessKeyID, conf.SecretKey, ""),
	})

	if err != nil {
		panic(err)
	}

	svc := s3.New(session)
	uploader := s3manager.NewUploader(session)

	return &S3Bucket{
		ctx:      ctx,
		svc:      svc,
		uploader: uploader,
		region:   aws.String(conf.Region),
	}
}

func (r *S3Bucket) ListBuckets(ctx context.Context) ([]*s3.Bucket, error) {
	res, err := r.svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("bucket not found: %+v", err))
	}
	return res.Buckets, nil
}

// func (r *S3Bucket) CreateBucket(ctx context.Context, bucketName string) (*s3.CreateBucketOutput, error) {
// 	res, err := r.svc.CreateBucket(&s3.CreateBucketInput{
// 		Bucket: &bucketName,
// 		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
// 			LocationConstraint: r.region,
// 		},
// 	})

// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("bucket create failed: %+v", err))
// 	}

// 	return res, nil
// }

func (r *S3Bucket) GetPreSignedUrl(bucketname, filename string) (string, error) {
	req, _ := r.svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucketname),
		Key:    aws.String(path.Join("/images/", filename)),
	})

	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", errors.New(fmt.Sprintf("get PresignedUrl failed: %+v", err))
	}

	return urlStr, nil
}

func (r *S3Bucket) UploadFile(file io.ReadSeeker, bucketname, filePath, filename string) (*s3manager.UploadOutput, error) {
	res, err := r.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketname),
		Key:    aws.String(path.Join("/images/", filename)),
		Body:   file,
	})

	if err != nil {
		return nil, errors.New(fmt.Sprintf("upload file failed: %+v", err))
	}

	return res, err
}
