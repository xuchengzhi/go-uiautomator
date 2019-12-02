package Base

import (
    "context"
    "encoding/json"
    "net/http"
    "strings"
    "bytes"
    "github.com/xuchengzhi/Library/Time"
    "io/ioutil"
    "log"
    "time"
    "net/url"
    "fmt"
    // "unsafe"
    // "reflect"
)


type TimeoutRequestsSession struct {
    Host string
    Port string
}


type Params struct {

}

type JsonStr struct {
    Jsonrpc string `json:"jsonrpc"`
    ID      string `json:"id"`
    Method  string `json:"method"`
    Params  Params `json:"params"`
}

type Par struct {
    Url    string
    Params interface{}
}


type Server struct {
    Addr string
    server_url string
}

type ApiJson struct {
    Status int         `json:"code"`
    Msg    interface{} `json:"msg"`
    Data   interface{} `json:"info"`
    Time   interface{} `json:"responsetime"`
    Dates  interface{} `json:"rundates"`
}


type infoJson struct {
    Jsonrpc string `json:"jsonrpc"`
    ID      string `json:"id"`
    Result  struct {
        CurrentPackageName string `json:"currentPackageName"`
        DisplayHeight      int    `json:"displayHeight"`
        DisplayRotation    int    `json:"displayRotation"`
        DisplaySizeDpX     int    `json:"displaySizeDpX"`
        DisplaySizeDpY     int    `json:"displaySizeDpY"`
        DisplayWidth       int    `json:"displayWidth"`
        ProductName        string `json:"productName"`
        ScreenOn           bool   `json:"screenOn"`
        SdkInt             int    `json:"sdkInt"`
        NaturalOrientation bool   `json:"naturalOrientation"`
    } `json:"result"`
}

var timeout time.Duration

var is_res, is_proxy bool

func init() {

    timeout = time.Duration(1000 * time.Millisecond)
    is_res = true
    is_proxy = true
}


func Post(urls string,pars map[string]string, is_proxy bool) ApiJson {
    var ch ApiJson
    log.Println(pars)
    var clusterinfo = url.Values{}
    for key, val := range pars {
        clusterinfo.Add(key, string(val))
    }

    data := clusterinfo.Encode()


    // bytesData, err := json.Marshal(par)
    req, _ := http.NewRequest("POST", urls, strings.NewReader(data))
    ctx, _ := context.WithTimeout(context.Background(), timeout) //设置超时时间
    req = req.WithContext(ctx)
    req.Header.Set("User-Agent", "goTest")
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    // req.Header.Set("Content-Type", "application/json;charset=UTF-8")

    proxy := func(*http.Request) (*url.URL, error) {
        return url.Parse("http://127.0.0.1:8888")
    }

    transport := &http.Transport{}
    if is_proxy {
        transport = &http.Transport{Proxy: proxy}
    }

    client := &http.Client{
        Timeout:   timeout,
        Transport: transport,
    }
    t1 := time.Now()
    resp, err := client.Do(req)
    // resp, err := transport.RoundTrip(req)
    t2 := time.Now()
    // runtime := (t2.Sub(t1))
    var duration time.Duration = t2.Sub(t1)

    runtime := fmt.Sprintf("%.03f S", duration.Seconds())
    
    // fmt.Println(duration)
    if err != nil {
        log.Println(err)
        ch = ApiJson{2001,  "error","req err", runtime, GetTime.TS()}
        return ch
    }

    defer resp.Body.Close()
    body, errs := ioutil.ReadAll(resp.Body)
    log.Println(resp.StatusCode)
    code := resp.StatusCode
    if errs != nil {
        log.Println(errs)
        ch = ApiJson{2002, "error",  "res Error", runtime, GetTime.TS()}
        return ch
    }


    if code == 200{
        var info infoJson
        json.Unmarshal([]byte(body), &info)
        ch = ApiJson{200, "success",info, runtime, GetTime.TS()}
    }
    return ch
}

func PostJson(p Par, is_proxy bool) ApiJson {
    var ch ApiJson
    urls := p.Url
    par := p.Params

    bytesData, err := json.Marshal(par)
    req, _ := http.NewRequest("POST", urls, bytes.NewReader(bytesData))
    ctx, _ := context.WithTimeout(context.Background(), timeout) //设置超时时间
    req = req.WithContext(ctx)
    req.Header.Set("User-Agent", "goTest")
    // req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Content-Type", "application/json;charset=UTF-8")

    proxy := func(*http.Request) (*url.URL, error) {
        return url.Parse("http://127.0.0.1:8888")
    }

    transport := &http.Transport{}
    if is_proxy {
        transport = &http.Transport{Proxy: proxy}
    }

    client := &http.Client{
        Timeout:   timeout,
        Transport: transport,
    }
    t1 := time.Now()
    resp, err := client.Do(req)
    // resp, err := transport.RoundTrip(req)
    t2 := time.Now()
    // runtime := (t2.Sub(t1))
    var duration time.Duration = t2.Sub(t1)

    runtime := fmt.Sprintf("%.03f S", duration.Seconds())
    
    // fmt.Println(duration)
    if err != nil {
        log.Println(err)
        ch = ApiJson{2001,  "error","req err", runtime, GetTime.TS()}
        return ch
    }

    defer resp.Body.Close()
    body, errs := ioutil.ReadAll(resp.Body)
    log.Println(resp.StatusCode)
    code := resp.StatusCode
    if errs != nil {
        log.Println(errs)
        ch = ApiJson{2002, "error",  "res Error", runtime, GetTime.TS()}
        return ch
    }


    if code == 200{
        var info infoJson
        json.Unmarshal([]byte(body), &info)
        ch = ApiJson{200, "success",info, runtime, GetTime.TS()}
    }
    return ch
}