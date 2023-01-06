package main

import (
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)

	config = clientv3.Config{
		Endpoints:   []string{"42.192.41.249:2379"},
		DialTimeout: 20 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("开始连接:", err)
	client = client
}
