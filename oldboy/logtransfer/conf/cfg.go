package conf

type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"`
	ESCfg    `ini:"es"`
}

type KafkaCfg struct {
	Addr  string `ini:"addr"`
	Topic string `ini:"topic"`
}

type ESCfg struct {
	Addr     string `ini:"addr"`
	ChanSize int    `ini:"chan_size"`
	Nums     int    `ini:"nums"`
}
