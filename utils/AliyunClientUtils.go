package utils

import (
	"AliyunStat/settings"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
)

func NewAliyunCdnClient() *cdn.Client {
	cdnClient, err := cdn.NewClientWithAccessKey(settings.CdnRegionId, settings.AccessKeyId, settings.SecretAccessKey)

	if err != nil {
		LogError(err)
	}
	return cdnClient
}
