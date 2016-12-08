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
	"github.com/aws/aws-sdk-go/service/s3"
)

type fileinfo struct {
	info os.FileInfo
	path string
}

var (
	filelist []*fileinfo
	path     = flag.String("p", "", "Path")
	dryRun   = flag.Bool("d", false, "Dry run")
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
		filelist = append(filelist, &fileinfo{path: path, info: info})
	}
	return nil
}

func uploadfile(sess *session.Session, f fileinfo, wg *sync.WaitGroup) {
	defer wg.Done()

	data, _ := ioutil.ReadFile(f.path)
	fhex := fmt.Sprintf("%x", md5.Sum(data))

	params := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWSUPLOADBUCKET")),
		Key:    aws.String(f.path),
		Body:   bytes.NewReader(data),
	}
	up := s3.New(sess)
	log.Println("Upload:", f.path)
	if *dryRun {
		log.Println(f.path, fhex)
	} else {
		if resp, err := up.PutObject(params); err != nil {
			log.Println("Err:", f.path, err)
		} else {
			log.Println(f.path, fhex, *resp.ETag)
		}
	}
}

func main() {
	flag.Parse()
	if *path == "" {
		log.Fatalln("Need Path")
	}
	filepath.Walk(*path, walkFunc)
	var wg sync.WaitGroup
	wg.Add(len(filelist))
	sess := getSession()
	for _, f := range filelist {
		go uploadfile(sess, *f, &wg)
	}
	wg.Wait()
}
