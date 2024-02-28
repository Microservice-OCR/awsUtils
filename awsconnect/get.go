package awsconnect

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Get(filename string) ([]byte,error){

	bucket, ok := os.LookupEnv("AWS_BUCKET_NAME")
	if !ok {
		log.Fatal("AWS_BUCKET_NAME not found")
	}

    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("eu-west-3"),
    }))

    s3Svc := s3.New(sess)
    result, err := s3Svc.GetObject(&s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(filename),
    })

    if err != nil {
        fmt.Println("Erreur lors de la récupération de l'image", err)
        return nil,err
    }

    body, err := ioutil.ReadAll(result.Body)
    if err != nil {
        fmt.Println("Erreur lors de la lecture du contenu", err)
        return nil,err
    }

    fmt.Println("Image téléchargée avec succès!")
	return body,nil
}
