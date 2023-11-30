package main

import (
	"context"
	"github.com/earayu/go-playground/grpc_server"
	"google.golang.org/grpc"
	"math/rand"
	"time"
)

// protoc --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=.. --go-grpc_out=.. *.proto
func main() {
	addrList := []string{
		"localhost:16101",
		"localhost:16100",
		"localhost:16102",
	}

	setAsPrimary(addrList[2])
	//time.Sleep(1 * time.Second)
	setAsPrimary(addrList[1])
	//time.Sleep(1 * time.Second)
	setAsPrimary(addrList[0])

	//for i := 0; i < 1000; i++ {
	//	addrList = shuffleList(addrList)
	//	for _, addr := range addrList {
	//		go setAsPrimary(addr)
	//	}
	//	time.Sleep(10 * time.Millisecond)
	//}
	//
	//select {}
}

func setAsPrimary(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	tmc := grpc_server.NewTabletManagerClient(conn)
	resp, err := tmc.ChangeType(context.Background(), &grpc_server.ChangeTypeRequest{
		TabletType: grpc_server.TabletType_PRIMARY,
		SemiSync:   true,
	})
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func shuffleList(list []string) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 创建一个新的切片，避免修改原始列表
	shuffledList := make([]string, len(list))
	copy(shuffledList, list)

	// Fisher-Yates 洗牌算法
	for i := len(shuffledList) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i]
	}

	return shuffledList
}
