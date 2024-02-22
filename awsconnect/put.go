package awsconnect

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

func Put(file *multipart.FileHeader) (string, error) {

	bucket, ok := os.LookupEnv("AWS_BUCKET_NAME")
	if !ok {
		log.Fatal("AWS_BUCKET_NAME not found")
	}
	
	src, err := file.Open()
	if err != nil {
		return "",err
	}
	defer src.Close()

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-west-3"),
	}))

	s3Svc := s3.New(sess)

	id := uuid.New()
	filename := fmt.Sprintf("%s.png", id.String())
	fmt.Println(filename)

	_, err = s3Svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(filename),
		Body:        src,
		ContentType: aws.String("image/png"),
	})

	if err != nil {
		fmt.Println("Erreur lors de l'upload de l'image", err)
		return "",nil
	}

	fmt.Println("Image uploadée avec succès!")
	return filename,nil
}
