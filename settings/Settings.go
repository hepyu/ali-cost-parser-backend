//refer: https://cloud.tencent.com/developer/article/1066126
package settings

import (
	"github.com/go-ini/ini"
)

var AccessKeyId string
var SecretAccessKey string

var CdnRegionId string

//logrus log config
var LogPath string
var LogFileName string
var MaxAge string
var RotationTime string

func init() {
	cfg, _ := ini.Load("/etc/aliyun/config.ini")

	//获取aliyun的accessKey等相关信息
	section := cfg.Section("aliyun-account")
	AccessKeyId = section.Key("accessKeyId").String()
	SecretAccessKey = section.Key("secretAccessKey").String()

	//获取aliyun cdn相关配置
	section := cfg.Section("aliyun-cdn")
	CdnRegionId = section.Key("cdnRegionId").String()

	//获取log相关配置
	section := cfg.Section("logrus-config")
	LogPath = section.Key("logPath").String()
	LogFileName = section.Key("logFileName").String()
	//文件最大保存时间
	MaxAge = section.Key("maxAge").String()
	//日志切割时间间隔
	RotationTime = section.Key("rotationTime").String()

}
