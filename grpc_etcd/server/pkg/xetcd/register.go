package xetcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

const (
	etcdEndpoints = "localhost:2388" // etcd 服务地址
	serviceName   = "search"
	serviceKey    = "services/" + serviceName + "/1"
	serviceValue  = "127.0.0.1:9099"
)

func RegisterEtcdService() {
	cfg := clientv3.Config{
		Endpoints:         []string{etcdEndpoints},
		DialKeepAliveTime: 5 * time.Second,
	}
	client, err := clientv3.New(cfg)
	if err != nil {
		panic(err)
	}

	leaseResp, err := client.Grant(context.TODO(), 5)
	if err != nil {
		panic(err)
	}
	/*manager, _ := endpoints.NewManager(client, serviceKey)
	manager.AddEndpoint()
	manager.DeleteEndpoint()*/
	putResp, err := client.Put(context.Background(), serviceKey, serviceValue, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		panic(err)
	}
	fmt.Println(putResp)

	// 保持租约
	lkaRespCh, err := client.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case ka, ok := <-lkaRespCh:
				if !ok {
					log.Println("keep alive channel closed")
					return
				}
				log.Printf("Received keep alive response: %v\n", ka)
			}
		}
	}()
}
