/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/28 13:57
 */
package conf

import (
	"errors"
	"flag"
	"strings"
)

var C = new(Config)

var config string
var nacosID string

func init() {
	flag.StringVar(&nacosID, "nacos_id", "", "")
	flag.StringVar(&config, "conf", "config.toml", "")
	flag.Parse()

	if config == "" {
		panic("invalid config")
	}
	if err := Prepare(); err != nil {
		panic("load config failed")
	}
}

func Prepare() error {
	if strings.HasPrefix(config, "http://") || strings.HasPrefix(config, "https://") {
		//return loadFromNacos(nacosID, config) //TODO
		return errors.New("not supported this moment")
	}
	return loadFromFile(config)
}
