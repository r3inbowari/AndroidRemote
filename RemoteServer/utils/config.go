package bilicoin

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"regexp"
	"time"
)

// LocalConfig /**
type LocalConfig struct {
	Finger         string    `json:"finger"`           // canvas指纹
	APIAddr        string    `json:"api_addr"`         // API服务ADDR
	CacheTime      time.Time `json:"-"`                // 缓存时间
	LoggerLevel    *string   `json:"log_level"`        // 日志等级
	WsAddr         string    `json:"ws_addr"`          // Ws服务ADDR
	RdbURL         string    `json:"rdb_url"`          // redis url
	RdbPassword    string    `json:"rdb_password"`     // redis password
	RdbIndex       int       `json:"rdb_index"`        // Rdb 数据库索引
	DeviceRenewTTL int       `json:"device_renew_ttl"` // 设备续约时间
	RmqURL         string    `json:"rmq_url"`          // RMQ url
	RmqPassword    string    `json:"rmq_password"`     // RMQ password
	RmqIndex       int       `json:"rmq_index"`        // RMQ 数据库索引
}

var config = new(LocalConfig)
var configPath = "config.json"

func GetConfig() *LocalConfig {
	if config.CacheTime.Before(time.Now()) {
		if err := LoadConfig(configPath, config); err != nil {
			Fatal("loading file failed")
			return nil
		}
		config.CacheTime = time.Now().Add(time.Second * 60)
	}
	return config
}

/**
 * save cnf/conf.json
 */
func (lc *LocalConfig) SetConfig() error {
	fp, err := os.Create(configPath)
	if err != nil {
		Fatal("loading file failed", logrus.Fields{"err": err})
	}
	defer fp.Close()
	data, err := json.Marshal(lc)
	if err != nil {
		Fatal("marshal file failed", logrus.Fields{"err": err})
	}
	n, err := fp.Write(data)
	if err != nil {
		Fatal("write file failed", logrus.Fields{"err": err})
	}
	Info("[FILE] Update user configuration", logrus.Fields{"size": n})
	return nil
}

const configFileSizeLimit = 10 << 20

/**
 * Load File
 * @param path 文件路径
 * @param dist 存放目标
 */
func LoadConfig(path string, dist interface{}) error {
	configFile, err := os.Open(path)
	if err != nil {
		Fatal("Failed to open config file.", logrus.Fields{"path": path, "err": err})
		return err
	}

	fi, _ := configFile.Stat()
	if size := fi.Size(); size > (configFileSizeLimit) {
		Fatal("Config file size exceeds reasonable limited", logrus.Fields{"path": path, "size": size})
		return errors.New("limited")
	}

	if fi.Size() == 0 {
		Fatal("Config file is empty, skipping", logrus.Fields{"path": path, "size": 0})
		return errors.New("empty")
	}

	buffer := make([]byte, fi.Size())
	_, err = configFile.Read(buffer)
	buffer, err = StripComments(buffer)
	if err != nil {
		Fatal("Failed to strip comments from json", logrus.Fields{"err": err})
		return err
	}

	buffer = []byte(os.ExpandEnv(string(buffer)))

	err = json.Unmarshal(buffer, &dist)
	if err != nil {
		Fatal("Failed unmarshalling json", logrus.Fields{"err": err})
		return err
	}
	return nil
}

/**
 * 注释清除
 */
func StripComments(data []byte) ([]byte, error) {
	data = bytes.Replace(data, []byte("\r"), []byte(""), 0)
	lines := bytes.Split(data, []byte("\n"))
	filtered := make([][]byte, 0)

	for _, line := range lines {
		match, err := regexp.Match(`^\s*#`, line)
		if err != nil {
			return nil, err
		}
		if !match {
			filtered = append(filtered, line)
		}
	}
	return bytes.Join(filtered, []byte("\n")), nil
}
