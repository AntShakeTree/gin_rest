package common

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	App        map[string]string
	Datasource map[string](map[string]string)
	Logger     map[string]string
	Broker     map[string]string
	All        map[string]string
	Err        map[string]int
}

func (cfg *Config) Parse(fpath string) {
	//初始化
	cfg.App = make(map[string]string)
	cfg.All = make(map[string]string)
	//cfg.Session = make(map[string]string)
	cfg.Datasource = make(map[string](map[string]string))
	//cfg.Static = make(map[string]string)
	//cfg.StaticFile = make(map[string]string)
	//cfg.View = make(map[string]string)
	cfg.Logger = make(map[string]string)
	//cfg.Logger = make(map[string]string)
	//cfg.TempFileMap = make(map[string]int)
	cfg.Broker = make(map[string]string)
	cfg.Err = make(map[string]int)

	fi, err := os.Open(fpath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		panic(err)
		return
	}
	defer fi.Close()
	br := bufio.NewReader(fi)

	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		tmp := strings.TrimLeft(string(a), " ")
		tmp = strings.TrimRight(tmp, " ")
		if len(tmp) == 0 || strings.Index(tmp, "#") == 0 {
			continue
		}
		o := strings.Split(tmp, "=")
		//配置文件没值就忽略

		if len(o) < 2 {
			continue
		}

		if len(o) == 2 {
			cfg.All[o[0]] = o[1]
		} else {
			cfg.All[o[0]] = strings.TrimPrefix(tmp, o[0]+"=")
		}

	}

	////轮训
	for k, v := range cfg.All {
		if strings.Index(k, "app.") == 0 {
			tmp := strings.TrimPrefix(k, "app.")
			cfg.App[tmp] = v
		} else if strings.Index(k, "logger.") == 0 {
			tmp := strings.TrimPrefix(k, "logger.")
			cfg.Logger[tmp] = v
		} else if strings.Index(k, "broker.") == 0 {
			tmp := strings.TrimPrefix(k, "broker.")
			cfg.Broker[tmp] = v
		} else if strings.Index(k, "err.") == 0 {
			tmp := strings.TrimPrefix(k, "err.")
			i, err := strconv.Atoi(v)
			if err != nil {
				cfg.Err[tmp] = -1
			} else {
				cfg.Err[tmp] = i
			}
		} else if strings.Index(k, "database.") == 0 {
			tmp := strings.TrimPrefix(k, "database.")
			var sd = strings.Split(tmp, ".")
			if cfg.Datasource[sd[0]] == nil {
				cfg.Datasource[sd[0]] = make(map[string]string)
			}
			cfg.Datasource[sd[0]][sd[1]] = v
		}
	}
}
func  GetBrokerId() string {
	return GetBroker()["role"] + "_" +  GetBroker()["id"]
}

//获取整数,
func (cfg *Config) LoadCfg(key string) string {
	return cfg.All[key]
}

//获取字符串配置
func (cfg *Config) LoadString(key string) string {
	return cfg.All[key]
}
func GetBroker()  map[string]string{
	return GetCfg().Broker
}
//获取整数,
func (cfg *Config) LoadInt(key string) (int, error) {
	return strconv.Atoi(cfg.All[key])
}

//获取32位整数
func (cfg *Config) LoadInt64(key string) (int64, error) {
	return strconv.ParseInt(cfg.All[key], 10, 64)
}

//获取64位整数
func (cfg *Config) LoadInt32(key string) (int64, error) {
	return strconv.ParseInt(cfg.All[key], 10, 32)
}

//获取布尔配置
func (cfg *Config) LoadBool(key string) bool {
	return cfg.All[key] == "true" || "TRUE" == cfg.All[key]
}

//cfg.loadCfg("restgo.weixin.appid")

var _cfg *Config = nil

func SetCfg(c *Config) {
	_cfg = c
}
func GetCfg() *Config {
	return _cfg
}
