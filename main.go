package main

import (
	"goblog/model"
	"goblog/routes"
)

func main(){
	model.Init()
	routes.InitRouter()
}