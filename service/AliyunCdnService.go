package service

import (
	"AliyunStat/settings"
	"AliyunStat/utils"
	"AliyunStat/wrapper"
	//"os"
	"time"
	//"regexp"
	"fmt"
)

//下载某个自然天的日志
//date格式：2019-01-01
func AliyunCdnGetCdnLogUrls(domainName string, datestr string) {

	startTime := utils.DateFormatByDateInLoca(datestr)
	endTime := startTime.AddDate(0, 0, 1)

	data = wrapper.AliyunCdnDescribeCdnDomainLogsRequest(domainName, 1, 10000, utils.DateFormatGetStrByISO8601Time(startTime), utils.DateFormatGetStrByISO8601Time(endTime))

	cdnLogDir = settings.CdnLogDownloadDir
	existDir, _ = utils.LocalPathExists(cdnLogDir)
	if !existDir {
		utils.CreateLocalDir(cdnLogDir)
	}

	/*
		reg = ("%s\\/(.+?)\\?" % (args.date.replace("-", "_")))
		for _, url := range data {
			fname = cdnLogDir +
		}
	*/
	fmt.Println("hehe means ok a lot")
}
