package main

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	_ "rabbitMQ/routers"
)

func main() {
	//beego.Run()
	var caonima = make(map[string]interface{})
	caonima["123"] = 123
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	_ = jsonEncoder.Encode(caonima)
	logs.Info(string(bytes.NewBuffer(bf.Bytes()).Bytes()))
}
