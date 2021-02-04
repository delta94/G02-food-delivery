package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"fooddlv/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName string, region string, apiKey string, secret string, domain string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey, // Access key ID
			provider.secret, // Secret access key
			""),             // Token can be ignore
	})

	if err != nil {
		log.Fatalln(err)
	}

	provider.session = s3Session

	return provider
}

func (provider *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)

	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(provider.bucketName),
		Key:    aws.String(dst),
		ACL:    aws.String("private"),
		Body:   fileBytes,
	})

	//t := time.Now().Add(time.Second * 20)
	//f, _ := s3.New(provider.session).PutObjectRequest(&s3.PutObjectInput{
	//	Bucket:  aws.String(provider.bucketName),
	//	Key:     aws.String(dst),
	//	ACL:     aws.String("private"),
	//	Expires: &t,
	//})
	//
	//f.Sign()

	if err != nil {
		return nil, err
	}

	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "s3",
	}

	return img, nil
}
