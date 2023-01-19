package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// 集群列表
const ETCD_URL = "127.0.0.1:2379"

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		kv     clientv3.KV
		putOp  clientv3.Op
		getOp  clientv3.Op
		opResp clientv3.OpResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{ETCD_URL}, // 集群列表
		DialTimeout: 5 * time.Second,
	}

	// 建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 用于读etcd的键值对
	kv = clientv3.NewKV(client)

	// 创建Op: operation
	putOp = clientv3.OpPut("/cron/jobs/job8", "123123123")

	// 执行OP
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}

	// kv.Do(op)

	// kv.Put
	// kv.Get
	// kv.Delete

	fmt.Println("写入Revision:", opResp.Put().Header.Revision)

	// 创建Op
	getOp = clientv3.OpGet("/cron/jobs/job8")

	// 执行OP
	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}

	// 打印
	fmt.Println("数据Revision:", opResp.Get().Kvs[0].ModRevision) // create rev == mod rev
	fmt.Println("数据value:", string(opResp.Get().Kvs[0].Value))

}
