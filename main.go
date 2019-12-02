package main

import (
    // "reflect"
    // "github.com/xuchengzhi/UIautomator2/Base"
    "github.com/xuchengzhi/UIautomator2/UIAutomatorServer"
    "log"
)


func main(){
    var s UIAutomatorServer.Server
    s.Addr = "192.168.248.54"
    // info := s.Info()
    res := (UIAutomatorServer.App_start(s,"com.nearme.themespace"))
    log.Println(res)
    // if info.Status == 200{
    //     log.Println(info.Data)
    // }else{
    //     log.Println(info.Data)
    // }
}