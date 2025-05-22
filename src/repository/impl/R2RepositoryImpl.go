package impl

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/ruiborda/ecommerce-product-service/src/repository"
	"log/slog"
	"mime"
	"net/http"
)

type R2RepositoryImpl struct {
	bucketName      string
	accountId       string
	accessKeyId     string
	accessKeySecret string
}

func NewR2RepositoryImpl(
	bucketName string,
	accountId string,
	accessKeyId string,
	accessKeySecret string,
) *R2RepositoryImpl {
	return &R2RepositoryImpl{
		bucketName:      bucketName,
		accountId:       accountId,
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
	}
}

func (this *R2RepositoryImpl) UploadFile(file *[]byte) (fileName string, err error) {
	fileName = ""
	detectedContentType := http.DetectContentType(*file)
	extensions, err := mime.ExtensionsByType(detectedContentType)

	if err != nil || len(extensions) == 0 {
		err = fmt.Errorf("could not detect file extension")
		slog.Error("Error detecting file extension", "error", err)
		return
	}

	fileName = fmt.Sprintf("%s%s", uuid.New().String(), extensions[0])

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(this.accessKeyId, this.accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		slog.Error("Error loading default config", "error", err)
		return
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", this.accountId))
	})

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &this.bucketName,
		Key:         &fileName,
		Body:        bytes.NewReader(*file),
		ContentType: &detectedContentType,
	})
	if err != nil {
		slog.Error("Error uploading file", "error", err)
		return "", err
	}
	return
}

func (this *R2RepositoryImpl) UploadBase64File(base64File *string) (fileName string, err error) {
	decodedData, err := base64.StdEncoding.DecodeString(*base64File)
	if err != nil {
		slog.Error("Error decoding base64 file", "error", err)
		return "", err
	}
	return this.UploadFile(&decodedData)
}

func (this *R2RepositoryImpl) DeleteFile(fileName string) (err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(this.accessKeyId, this.accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", this.accountId))
	})

	_, err = client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &this.bucketName,
		Key:    &fileName,
	})

	return
}

func (this *R2RepositoryImpl) HeadObject(fileName string) *repository.HeadObject {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(this.accessKeyId, this.accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", this.accountId))
	})

	response, err := client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: &this.bucketName,
		Key:    &fileName,
	})
	if err != nil {
		return nil
	}

	return &repository.HeadObject{
		FileName:      fileName,
		ContentLength: *response.ContentLength,
		ContentType:   *response.ContentType,
		LastModified:  response.LastModified.Unix(),
	}
}
