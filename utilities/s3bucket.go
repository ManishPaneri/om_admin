package utilities

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
	"om_admin/models"
	// "os"
)

var S3_REGION, S3_BUCKET string

func LogS3BucketInitialization() {
	S3REGION, errRegion := models.GetConfigByTypeAndKey("admin/request/s3/region", "S3_REGION")
	S3BUCKET, errBucket := models.GetConfigByTypeAndKey("admin/request/s3/bucket", "S3_BUCKET")

	if errRegion != nil || errBucket != nil {
		log.Fatal(errRegion, errBucket)
		return
	}

	S3_REGION = S3REGION.Value
	S3_BUCKET = S3BUCKET.Value
}

func ImportS3BucketInCSVLogs(fileName string, fileData string) {
	// LogS3BucketInitialization()

	if S3_REGION == "" || S3_BUCKET == "" {
		fmt.Println("*****S3 Variable is empty!********")
		return
	}

	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		log.Fatal(err)
	}

	// Upload
	err = AddFileToS3(s, fileName, fileData)
	if err != nil {
		log.Fatal(err)
	}

	s.Handlers.Clear()
	return
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, fileDir string, fileData string) error {

	// Get file size and read the file content into a buffer
	buffer := []byte(fileData)
	size := int64(len(fileData))

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(S3_BUCKET),
		Key:                  aws.String(fileDir),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	fmt.Println("s3bucket error:", err)
	return err
}
