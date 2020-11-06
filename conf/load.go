/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/28 14:00
 */
package conf

import (
	"github.com/BurntSushi/toml"
)

//TODO
//func loadFromNacos(id string, addr string) error {
//	sc := []constant.ServerConfig{
//		{
//			IpAddr: "console.nacos.io",
//			Port:   80,
//		},
//	}
//
//	cc := constant.ClientConfig{
//		NamespaceId:         id, //"e525eafa-f7d7-4029-83d9-008937f9d468", //namespace id
//		TimeoutMs:           5000,
//		NotLoadCacheAtStart: true,
//		LogDir:              "/tmp/nacos/log",
//		CacheDir:            "/tmp/nacos/cache",
//		RotateTime:          "1h",
//		MaxAge:              3,
//		LogLevel:            "debug",
//	}
//	client, err := clients.CreateConfigClient(map[string]interface{}{
//		"serverConfigs": sc,
//		"clientConfig":  cc,
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	content, err := client.GetConfig(vo.ConfigParam{
//		DataId: "test-data",
//		Group:  "test-group",
//	})
//
//	err = client.ListenConfig(vo.ConfigParam{
//		DataId: "test-data",
//		Group:  "test-group",
//		OnChange: func(namespace, group, dataId, data string) {
//			//fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
//		},
//	})
//}

func loadFromFile(filename string) error {
	_, err := toml.DecodeFile(filename, C)
	return err
}
