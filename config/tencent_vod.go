//go:build tencent && vod

package config

const (
	TencentVodRegionSeoul         = "ap-seoul"         // 亚太东北（首尔）
	TencentVodRegionTokyo         = "ap-tokyo"         // 亚太东北（东京）
	TencentVodRegionMumbai        = "ap-mumbai"        // 亚太南部（孟买）
	TencentVodRegionBangkok       = "ap-bangkok"       // 亚太东南（曼谷）
	TencentVodRegionBeijing       = "ap-beijing"       // 华北地区（北京）
	TencentVodRegionJakarta       = "ap-jakarta"       // 亚太东南（雅加达）
	TencentVodRegionAshburn       = "na-ashburn"       // 美国东部（弗吉尼亚）
	TencentVodRegionToronto       = "na-toronto"       // 北美地区（多伦多）
	TencentVodRegionChengdu       = "ap-chengdu "      // 西南地区（成都）
	TencentVodRegionSaoPaulo      = "sa-saopaulo"      // 南美地区（圣保罗）
	TencentVodRegionHongKong      = "ap-hongkong"      // 港澳台地区（中国香港）
	TencentVodRegionShanghai      = "ap-shanghai"      // 华东地区（上海）
	TencentVodRegionFrankfurt     = "eu-frankfurt"     // 欧洲地区（法兰克福）
	TencentVodRegionSingapore     = "ap-singapore"     // 亚太东南（新加坡）
	TencentVodRegionChongqing     = "ap-chongqing"     // 西南地区（重庆）
	TencentVodRegionGuangzhou     = "ap-guangzhou"     // 华南地区（广州）
	TencentVodRegionShanghaiFsi   = "ap-shanghai-fsi"  // 华东地区（上海金融）
	TencentVodRegionShenzhenFsi   = "ap-shenzhen-fsi"  // 华南地区（深圳金融）
	TencentVodRegionSiliconValley = "na-siliconvalley" // 美国西部（硅谷）
)

type TencentVod struct {
	SecretId  string  `json:"SecretId" yaml:"SecretId" mapstructure:"SecretId"`    // 访问密钥 Id
	SecretKey string  `json:"SecretKey" yaml:"SecretKey" mapstructure:"SecretKey"` // 访问密钥 Secret
	Region    string  `json:"Region" yaml:"Region" mapstructure:"Region"`          // 区域
	Tencent   Tencent `json:"Tencent" yaml:"-" mapstructure:"-"`
}
