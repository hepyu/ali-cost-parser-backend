package wrapper

import (
	"AliyunStat/utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
)

/*
(1).中文注釋:
调用DescribeCdnDomainLogs接口获取指定域名的原始访问日志的下载地址。

(2).English comment:
Call the DescribeCdnDomainLogs interface to get the download address of the original access log of the specified domain name.

(3).aliyun official url:
https://help.aliyun.com/document_detail/27224.html?spm=5176.10695662.1996646101.searchclickresult.1a7e7f7aXrrWkT
*/
func AliyunCdnDescribeCdnDomainLogsRequest(domainName string, pageNumber int, pageSize int, startTimeISO8601Str string, endTimeISO8601Str string) string {

	//cdnClient, err := cdn.NewClientWithAccessKey(settings.CDN_RegionId, settings.AccessKeyId, settings.SecretAccessKey)

	cdnClient := utils.NewAliyunCdnClient()

	request := cdn.CreateDescribeCdnDomainLogsRequest()

	request.AcceptFormat = "json"
	request.DomainName = domainName
	request.PageSize = requests.NewInteger(pageSize)
	request.PageNumber = requests.NewInteger(pageNumber)
	request.StartTime = startTimeISO8601Str
	request.EndTime = endTimeISO8601Str

	response, err := cdnClient.DescribeCdnDomainLogs(request)

	if err != nil {
		utils.LogError(err)
	}

	return response.GetHttpContentString()
}

/*
(1).中文注釋:
1.获取加速域名的网络流量监控数据，单位：byte。
2.不指定StartTime和EndTime时，默认读取过去24小时的数据，同时支持按指定的起止时>间查询，两者需要同时指定。

(2).English comment:
1.Get network traffic monitoring data of the accelerated domain name, in bytes.
2.When the start time and end time are not specified, the data of the past 24 hours is read by default, and the query is supported by the specified start and end time, both of which need to be specified at the same time.

(3).aliyun official url:
https://help.aliyun.com/document_detail/91045.html?spm=5176.10695662.1996646101.searchclickresult.55b120f8ZVylPD
*/
func AliyunCdnDescribeDomainTrafficDataRequest(domainName string, interval string, startTimeISO8601Str string, endTimeISO8601Str string) string {
	cdnClient := utils.NewAliyunCdnClient()

	request := cdn.CreateDescribeDomainTrafficDataRequest()

	request.AcceptFormat = "json"
	request.DomainName = domainName
	request.StartTime = startTimeISO8601Str
	request.EndTime = endTimeISO8601Str
	//单位是秒，建议值为5分钟,对应的赋值是300，cdn带宽打点
	request.Interval = interval

	response, err := cdnClient.DescribeDomainTrafficData(request)

	if err != nil {
		utils.LogError(err)
	}
	return response.GetHttpContentString()
}

/*
(1).中文注釋:
调用DescribeUserDomains接口查询用户名下所有的域名与状态。

(2).English comment:
Call the DescribeUserDomains interface to query all domain names and statuses under the username.

(3).aliyun official url:
https://help.aliyun.com/document_detail/104911.html?spm=5176.10695662.1996646101.searchclickresult.5f035295y6THV6
*/
func AliyunCdnDescribeUserDomainsRequest(domainStatus string, pageNumber int, pageSize int) string {

	cdnClient := utils.NewAliyunCdnClient()

	request := cdn.CreateDescribeUserDomainsRequest()

	request.AcceptFormat = "json"
	request.PageNumber = requests.NewInteger(pageNumber)
	request.PageSize = requests.NewInteger(pageSize)
	request.DomainStatus = domainStatus

	response, err := cdnClient.DescribeUserDomains(request)
	if err != nil {
		utils.LogError(err)
	}

	return response.GetHttpContentString()
}
