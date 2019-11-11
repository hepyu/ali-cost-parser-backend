package wrapper

import (
	"AliyunStat/settings"
	"AliyunStat/utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
	"github.com/sirupsen/logrus"
)

func AliyunCdnLogUrlsByDomain(domainName string, pageNumber int, pageSize int, startTimeISO8601Str string, endTimeISO8601Str string) {

	cdnClient, err := cdn.NewClientWithAccessKey(settings.CDN_RegionId, settings.AccessKeyId, settings.SecretAccessKey)

	if err != nil {
		utils.LogError(err)
	}

	request := cdn.CreateDescribeCdnDomainLogsRequest()
	request.DomainName = domainName
	request.PageSize = pageSize
	request.PageNumber = pageNumber
	request.StartTime = startTimeISO8601Str
	request.EndTime = endTimeISO8601Str

	response, err := cdnClient.DescribeCdnDomainLogs(request)

	if err != nil {
		utils.LogError(err)
	}

	return response
}
