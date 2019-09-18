package config

const (
	// AsyncTransferEnable 是否开启文件异步转移(默认:是)
	AsyncTransferEnable = true
	// RabbitURL rabbitmq服务入口
	RabbitURL = "amqp://guest:guest@127.0.0.1:5672/"
	// TransExchangeName 传输交换机名
	TransExchangeName = "uploadserver.trans"
	// TransOSSQueueName oss转移队列名
	TransOSSQueueName = "uploadserver.trans.oss"
	// TransOSSErrQueueName oss转移失败后写入另一个队列的队列名
	TransOSSErrQueueName = "uploadserver.trans.oss.err"
	// TransOSSRoutingKey routingkey
	TransOSSRoutingKey = "oss"
)
