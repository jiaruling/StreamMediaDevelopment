package oss

import (
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jiaruling/StreamMediaDevelopment/scheduler/config"
)

var (
	EP string
	AK string
	SK string
)

func init() {
	AK = config.GetAccessKeyId()
	SK = config.GetAccessKeySecret()
	EP = config.GetOssAddr()
}

func UploadToOss(filename string, path string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Getting Service error: %s", err)
		return false
	}
	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("Getting bucket error: %s", err)
		return false
	}
	err = bucket.PutObjectFromFile(filename, path)
	if err != nil {
		log.Printf("upload object error: %s", err)
		return false
	}
	return true
}

func DeleteObject(filename string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Getting Service error: %s", err)
		return false
	}
	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("Getting bucket error: %s", err)
		return false
	}
	err = bucket.DeleteObject(filename)
	if err != nil {
		log.Printf("delete object error: %s", err)
		return false
	}
	return true
}
