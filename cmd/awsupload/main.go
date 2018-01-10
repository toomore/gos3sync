package main

import (
	"bytes"
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type fileinfo struct {
	info os.FileInfo
	path string
}

var (
	concurrency = flag.Int("c", 20, "Concurrency")
	dryRun      = flag.Bool("d", false, "Dry run")
	path        = flag.String("p", "", "Path")

	uploadnum chan struct{}

	sess *session.Session
	wg   sync.WaitGroup
)

func getSession() *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(os.Getenv("AWSUPLOADREGION")),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("AWSUPLOADTOKEN"), os.Getenv("AWSUPLOADSECRET"), ""),
		})
	if err != nil {
		log.Panicln(err)
	}
	return sess
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err == nil {
		if info.IsDir() {
			return nil
		}
		wg.Add(1)
		go uploadfile(sess, fileinfo{path: path, info: info})
	} else {
		log.Println(err)
	}
	return nil
}

func uploadfile(sess *session.Session, f fileinfo) {
	uploadnum <- struct{}{}
	defer wg.Done()

	data, _ := ioutil.ReadFile(f.path)
	fhex := fmt.Sprintf("%x", md5.Sum(data))

	log.Println("Upload:", f.path)
	if *dryRun {
		log.Println("[DryRun]", f.path, fhex)
	} else {
		uploader := s3manager.NewUploader(sess)
		params := &s3manager.UploadInput{
			Bucket: aws.String(os.Getenv("AWSUPLOADBUCKET")),
			Key:    aws.String(f.path),
			Body:   bytes.NewReader(data),
		}
		if resp, err := uploader.Upload(params); err != nil {
			log.Println("!Err:", f.path, err)
		} else {
			log.Println("[U.OK]", f.path, fhex, resp)
		}
	}
	<-uploadnum
}

func main() {
	flag.Parse()
	if *path == "" {
		log.Fatalln("Need Path")
	}
	uploadnum = make(chan struct{}, *concurrency)

	sess = getSession()
	filepath.Walk(*path, walkFunc)
	wg.Wait()
}
