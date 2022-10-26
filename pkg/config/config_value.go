package config

import (
	"fmt"
	"github.com/paas-mate/gutil"
)

// kafka config
var (
	ClusterEnable                 bool
	RaftEnable                    bool
	ZkAddress                     string
	KafkaSocketSendBufferBytes    int64
	KafkaSocketReceiveBufferBytes int64
	KafkaMessageMaxBytes          int64
	KafkaFetchMessageMaxBytes     int64
	ReplicaFetchMaxBytes          int64
	KafkaAdvertiseAddress         string
	KafkaAdvertiseInf             string
	RemoteMode                    bool
	KafkaHost                     string
	KafkaPort                     int
	KafkaAddr                     string
)

func init() {
	ClusterEnable = gutil.GetEnvBool("CLUSTER_ENABLE", false)
	RaftEnable = gutil.GetEnvBool("RAFT_ENABLE", true)
	ZkAddress = gutil.GetEnvStr("ZK_ADDR", "localhost:2181")
	KafkaSocketSendBufferBytes = gutil.GetEnvInt64("KAFKA_SOCKET_SEND_BUFFER_BYTES", 102400)
	KafkaSocketReceiveBufferBytes = gutil.GetEnvInt64("KAFKA_SOCKET_RECEIVE_BUFFER_BYTES", 102400)
	KafkaMessageMaxBytes = gutil.GetEnvInt64("KAFKA_MESSAGE_MAX_BYTES", -1)
	KafkaFetchMessageMaxBytes = gutil.GetEnvInt64("KAFKA_FETCH_MESSAGE_MAX_BYTES", -1)
	ReplicaFetchMaxBytes = gutil.GetEnvInt64("REPLICA_FETCH_MAX_BYTES", -1)
	KafkaAdvertiseAddress = gutil.GetEnvStr("KAFKA_ADVERTISE_ADDRESS", "")
	KafkaAdvertiseInf = gutil.GetEnvStr("KAFKA_ADVERTISE_INF", "")
	RemoteMode = gutil.GetEnvBool("REMOTE_MODE", true)
	KafkaHost = "127.0.0.1"
	KafkaPort = 9092
	KafkaAddr = fmt.Sprintf("%s:%d", KafkaHost, KafkaPort)
}
