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
		config         clientv3.Config
		client         *clientv3.Client
		err            error
		kv             clientv3.KV
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
		keepResp       *clientv3.LeaseKeepAliveResponse
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
		ctx            context.Context
		cancelFunc     context.CancelFunc
		txn            clientv3.Txn
		txnResp        *clientv3.TxnResponse
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

	//申请lease，自动续租，抢占租约key
	lease = clientv3.NewLease(client)
	//申请5s租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 5); err != nil {
		fmt.Println(err)
		return
	}
	//获取租约id
	leaseId = leaseGrantResp.ID

	//取消自动续租的context
	ctx, cancelFunc = context.WithCancel(context.TODO())

	//确保函数退出后，自动续租会停止
	defer cancelFunc()
	defer lease.Revoke(context.TODO(), leaseId)

	//自动续租
	if keepRespChan, err = lease.KeepAlive(ctx, leaseId); err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已经失效")
					goto END
				} else {
					fmt.Println("收到自动续租应答", keepResp.ID)
				}

			}
		}
	END:
	}()

	//不存在 then 设置，else 失败
	kv = clientv3.NewKV(client)

	//创建，定义
	txn = kv.Txn(context.TODO())
	//如果key不存在
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job9"), "=", 0)).Then(clientv3.OpPut("/cron/lock/job9", "jobcher", clientv3.WithLease(leaseId))).Else(clientv3.OpGet("/cron/lock/job9"))

	if txnResp, err = txn.Commit(); err != nil {
		fmt.Println(err)
		return
	}

	//判断是否抢到锁
	if !txnResp.Succeeded {
		fmt.Println("锁被占用", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
	}

	//处理业务
	fmt.Println("处理业务")
	time.Sleep(5 * time.Second)

	//释放锁，取消自动续租，释放租约
}
