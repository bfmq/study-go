package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}

type KafkaConf struct {
	Addr        string `ini:"addr"`
	ChanMaxSize int    `ini:"chan_max_size"`
}

type EtcdConf struct {
	Addr    string `ini:"addr"`
	TimeOut int    `ini:"timeout"`
	Key     string `ini:"collect_log_key"`
}

// type TaillogConf struct {
// 	FileName string	`ini:"filename"`
// }
