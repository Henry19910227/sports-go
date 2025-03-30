package main

import (
	"api_gateway/engine"
	clientFactory "api_gateway/factory/client"
	controllerFactory "api_gateway/factory/controller"
	userRouter "api_gateway/router/user"
	"fmt"
	"sports-go/shared/pkg/tool/crypto"
	"strconv"
)

func main() {
	// 創建GRPC客戶端工廠
	clientFactory, err := clientFactory.New()
	if err != nil {
		fmt.Println(err)
	}
	// 創建控制器工廠
	factory := controllerFactory.New(clientFactory)

	// 創建 game engine
	engine := engine.New()
	// 添加路由解析器邏輯
	engine.PathResolver(func(b []byte) string {
		tool := crypto.New()
		mid, sid, _, _ := tool.UnMarshal(b)
		return "/" + strconv.Itoa(int(mid)) + "/" + strconv.Itoa(int(sid))
	})
	// 設定Base路由組
	baseGroup := engine.Group("/")
	baseGroup.Use(factory.MiddController().UnMarshalData)
	// 設定User路由組
	userGroup := baseGroup.Group("500/")
	userRouter.SetRoute(userGroup, factory)
	// 運行 game engine
	_ = engine.Run(":7070")
}
