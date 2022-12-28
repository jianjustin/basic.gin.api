package main

import (
	"basic.gin.api/middleware"
	"basic.gin.api/middleware/model"
	"basic.gin.api/users/entity"
	"basic.gin.api/users/service"
	"basic.gin.api/users/tool"
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
	"log"
	"time"
)

func main() {
	model.DB = middleware.BuildGormMysql()
	tool.RegisterTables()

	addr := flag.String("addr", "localhost:8972", "server address")

	//启动本地TCP Server
	go func() {
		s := server.NewServer()
		s.RegisterName("Arith", new(service.Arith), "")
		s.Serve("tcp", ":8972")
	}()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.JSON
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, opt)
	defer xclient.Close()

	args := service.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &service.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(time.Second)
	}

	user := &entity.SysUser{}
	model.DB.First(&user)

	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user": user,
		})
	})
	r.Run()

}
