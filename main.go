package main

import (
	"GinBlog/model"
	"GinBlog/routes"
)

func main() {
	//引用数据库
	model.InitDb()
	routes.InitRouter()
}
