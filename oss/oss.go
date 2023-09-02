package oss

import (
	"net/http"
	"net/url"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/22722679/douyin-project/config"
)

var cli *cos.Client

func init() {
	u, _ := url.Parse(config.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	cli = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.TencentSecretId,
			SecretKey: config.TencentSecretKey,
		},
	})

}
