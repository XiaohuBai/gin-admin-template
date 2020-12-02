/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-02 13:54:46
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-02 15:32:05
 * @Description: 描述
 */

package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	Name          string `mapstructure:"name" json:"name" yaml:"name"`
	Logo          string `mapstructure:"logo" json:"logo" yaml:"logo"`
}
