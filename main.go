package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	_ "goserver/libs/certificate"
	"goserver/libs/conf"
	"goserver/router"
)

func main() {
	os.Setenv("TZ", "Asia/Shanghai")

	serverCfg, _ := conf.GetSection("server")
	addr := fmt.Sprintf(":%d", serverCfg.Key("HTTP_PORT").MustInt())
	readTimeout := time.Duration(serverCfg.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	writeTimeout := time.Duration(serverCfg.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	server := &http.Server{
		Addr:           addr,
		Handler:        router.Router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServeTLS("conf/cert.pem", "conf/key.pem")
}
