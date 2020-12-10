/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-09 14:57:29
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-09 14:57:29
 * @Description: 描述
 */

package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu Qiniu `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
}
