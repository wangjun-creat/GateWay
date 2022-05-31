package main

import (
	"GateWay/GatewayDemo/proxy/middleware"
	"log"
	"net/http"

	proxy2 "GateWay/GatewayDemo/proxy/proxy"

	"GateWay/GatewayDemo/proxy/load_balance"
)

var (
	addr = "127.0.0.1:2002"
)

func main() {
	rb := load_balance.LoadBanlanceFactory(load_balance.LbWeightRoundRobin)
	rb.Add("http://127.0.0.1:2003", "50")
	proxy := proxy2.NewLoadBalanceReverseProxy(&middleware.SliceRouterContext{}, rb)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
