// 调用配置
package conf

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	conf *ini.File
)

func init() {
	var err error
	conf, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
}

func GetSection(sectionName string) (*ini.Section, error) {
	section, err := conf.GetSection(sectionName)
	if err != nil {
		log.Fatalf("Fail to get section '%v': %v", sectionName, err)
	}
	return section, err
}

func GetSectionKey(sectionName string, key string) *ini.Key {
	section, _ := GetSection(sectionName)
	return section.Key(key)
}
