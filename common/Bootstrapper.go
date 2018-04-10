package common


func init()  {
	//
	cfg := new(Config)
	cfg.Parse("config/broker.properties")
	SetCfg(cfg)
	//logger
	initLogger()
	//mongoDB
	initMgo()
	}