package main

import (
	"dhcpdb"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const (
	IPPOOL  string = "/v1/ippool/"
	IP      string = "/v1/ips/"
	DYNINFO string = "/v1/dyninfo"
)

const (
	PUT    string = "PUT"    //更新
	GET    string = "GET"    //查询
	POST   string = "POST"   //创建
	DELETE string = "DELETE" //删除
)

const (
	OPERATION_SUCCESS = "operation success"
	OPERATION_FAIL    = "operation fail"
)

//定义一个函数用于处理不同方法的请求
type restfulHandler func(w http.ResponseWriter, req *http.Request)

type cProvisionInstance struct {
	ippoolHandler  map[string]restfulHandler
	ipHandler      map[string]restfulHandler
	dyninfoHandler map[string]restfulHandler
}

type tIpPool struct {
	Ip      string
	NetMask string
	GateWay string
}
type tIps struct {
	Mac     string
	Ip      string
	NetMask string
	GateWay string
}
type tDynInfo struct {
	Mac string
	Ip  string
}
type tResult struct {
	Result string
}

var gProvisionInstance *cProvisionInstance
var gLock *sync.Mutex = &sync.Mutex{}

func ippoolPost(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	body_str := string(body)
	var ipPool tIpPool

	err := json.Unmarshal([]byte(body_str), &ipPool)
	if err != nil {
		fmt.Println("Decode ip pool post message error, errinfo=", err)
		w.WriteHeader(500)
		w.Write([]byte(OPERATION_FAIL))
		return
	}

	w.WriteHeader(http.StatusOK)
	var rs tResult = tResult{OPERATION_SUCCESS}
	jsrs, err := json.Marshal(rs)
	w.Write([]byte(jsrs))
	return
}

func ippoolPut(w http.ResponseWriter, req *http.Request) {

}
func ippoolGet(w http.ResponseWriter, req *http.Request) {

}
func ippoolDelete(w http.ResponseWriter, req *http.Request) {

}
func ipsPost(w http.ResponseWriter, req *http.Request) {

}
func ipsPut(w http.ResponseWriter, req *http.Request) {

}
func ipsGet(w http.ResponseWriter, req *http.Request) {

}
func ipsDelete(w http.ResponseWriter, req *http.Request) {

}
func dyninfoPost(w http.ResponseWriter, req *http.Request) {

}
func dyninfoPut(w http.ResponseWriter, req *http.Request) {

}
func dyninfoGet(w http.ResponseWriter, req *http.Request) {

}
func dyninfoDelete(w http.ResponseWriter, req *http.Request) {

}

func GetProvisionInstance() *cProvisionInstance {
	if nil == gProvisionInstance {
		gLock.Lock()
		defer gLock.Unlock()
		if nil == gProvisionInstance {
			gProvisionInstance = &cProvisionInstance{}
			gProvisionInstance.init()
		}
	}
	return gProvisionInstance
}

func (this *cProvisionInstance) init() {
	//为不同的对象登记不同的处理函数
	//为IPPOOL登记处理方法
	this.ippoolHandler = make(map[string]restfulHandler)
	this.ipHandler = make(map[string]restfulHandler)
	this.dyninfoHandler = make(map[string]restfulHandler)

	this.ippoolHandler[POST] = ippoolPost
	this.ippoolHandler[GET] = ippoolGet
	this.ippoolHandler[PUT] = ippoolPut
	this.ippoolHandler[DELETE] = ippoolDelete

	this.ipHandler[POST] = ipsPost
	this.ipHandler[GET] = ipsGet
	this.ipHandler[PUT] = ipsPut
	this.ipHandler[DELETE] = ipsDelete

	this.dyninfoHandler[POST] = dyninfoPost
	this.dyninfoHandler[GET] = dyninfoGet
	this.dyninfoHandler[PUT] = dyninfoPut
	this.dyninfoHandler[DELETE] = dyninfoDelete

}

func (this *cProvisionInstance) GetRestfulHandler(restObj string) *(map[string]restfulHandler) {
	switch restObj {
	case IPPOOL:
		return &this.ippoolHandler
	case IP:
		return &this.ipHandler
	case DYNINFO:
		return &this.dyninfoHandler
	default:
		return nil
	}
}

func IpPoolHandler(w http.ResponseWriter, req *http.Request) {
	var handler *map[string]restfulHandler = nil
	handler = GetProvisionInstance().GetRestfulHandler(IPPOOL)
	if nil == handler {
		return
	}
	if handlerfunc, ok := (*handler)[req.Method]; ok {
		handlerfunc(w, req)
	}
}

func IpHandler(w http.ResponseWriter, req *http.Request) {
	var handler *map[string]restfulHandler = nil
	handler = GetProvisionInstance().GetRestfulHandler(IP)
	if nil == handler {
		return
	}
	if handlerfunc, ok := (*handler)[req.Method]; ok {
		handlerfunc(w, req)
	}
}

func DynInfoHandler(w http.ResponseWriter, req *http.Request) {
	var handler *map[string]restfulHandler = nil
	handler = GetProvisionInstance().GetRestfulHandler(DYNINFO)
	if nil == handler {
		return
	}
	if handlerfunc, ok := (*handler)[req.Method]; ok {
		handlerfunc(w, req)
	}
}

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Method=", req.Method)
	fmt.Println("Boday=", req.Body)
	body, _ := ioutil.ReadAll(req.Body)
	body_str := string(body)
	fmt.Println("Boday=", body_str)
}

func main() {
	//http.HandleFunc(IPPOOL, IpPoolHandler)
	//http.HandleFunc(IP, IpHandler)
	//http.HandleFunc(DYNINFO, DynInfoHandler)
	//http.HandleFunc("/test/", test)
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	var redisdb dhcpdb.RedisDriver
	err := redisdb.ConnectDb("127.0.0.1:6379")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("Connect redis db success")

	err = redisdb.AddRecord("home", "wjdh")
	if err != nil {
		fmt.Println("Add key to redis error")
	}

	redisdb.DisconnectDb()
}
