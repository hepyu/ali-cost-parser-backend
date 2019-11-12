//refer: https://cloud.tencent.com/developer/article/1066126
package settings

import (
	"github.com/go-ini/ini"
	"strconv"
	"time"
)

var AccessKeyId string
var SecretAccessKey string

var CdnRegionId string

//logrus log config
var LogPath string
var LogFileName string

//文件最大保存时间，单位是小时
var LogMaxAge time.Duration

//日志切割时间间隔，单位是小时
var LogRotationTime time.Duration
var CdnLogDownloadDir string

func init() {
	cfg, _ := ini.Load("/etc/aliyun/config.ini")

	//获取aliyun的accessKey等相关信息
	section_aliyunAccount := cfg.Section("aliyun-account")
	AccessKeyId = section_aliyunAccount.Key("accessKeyId").String()
	SecretAccessKey = section_aliyunAccount.Key("secretAccessKey").String()

	//获取aliyun cdn相关配置
	section_aliyunCDN := cfg.Section("aliyun-cdn")
	CdnRegionId = section_aliyunCDN.Key("cdnRegionId").String()
	CdnLogDownloadDir = section_aliyunCDN.Key("cdnLogDownloadDir").String()

	//获取log相关配置
	section_log := cfg.Section("logrus-config")
	LogPath = section_log.Key("logPath").String()
	LogFileName = section_log.Key("logFileName").String()
	//文件最大保存时间,int64,单位是小时
	maxAgeInt64, _ := strconv.ParseInt((section_log.Key("maxAge").String()), 10, 64)
	LogMaxAge = time.Duration(maxAgeInt64) * time.Hour
	//日志切割时间间隔,单位是小时
	rotationTimeInt64, _ := strconv.ParseInt(section_log.Key("rotationTime").String(), 10, 64)
	LogRotationTime = time.Duration(rotationTimeInt64) * time.Hour

}
