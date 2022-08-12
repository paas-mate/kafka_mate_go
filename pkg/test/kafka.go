package test

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"kafka_mate_go/pkg/util"
	"net"
	"sync"
	"time"
)

const (
	kafkaContainerName = "test-kafka"
)

func SetupKafka() {
	var once sync.Once
	now := time.Now()
	for {
		if time.Since(now).Minutes() > 3 {
			panic("kafka still not started")
		}
		if simpleTcpCheck() {
			break
		} else {
			util.Logger().Error("connect kafka error ")
			once.Do(startKafka)
			time.Sleep(15 * time.Second)
			continue
		}
	}
}

func startKafka() {
	util.Logger().Info("start kafka container")
	err := startKafkaInternal()
	if err != nil {
		panic(err)
	}
}

func startKafkaInternal() error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	portSpecs, _, err := nat.ParsePortSpecs([]string{"9092"})
	if err != nil {
		return err
	}
	resp, err := cli.ContainerCreate(context.TODO(), &container.Config{
		Image:        "ttbb/kafka:mate",
		Env:          []string{"REMOTE_MODE=false"},
		ExposedPorts: portSpecs,
		Tty:          false,
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"9092/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "9092",
				},
			},
		},
	}, nil, nil, kafkaContainerName)
	if err != nil {
		panic(err)
	}

	return cli.ContainerStart(context.TODO(), resp.ID, types.ContainerStartOptions{})
}

func simpleTcpCheck() bool {
	servAddr := "localhost:9092"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		return false
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
