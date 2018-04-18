package dhcpdb

import (
	"fmt"
	"hash/crc32"
	"net"
	"strings"
)

type RedisDriver struct {
	redisServerIp       string
	redisServerPort     string
	redisUser           string
	redisPass           string
	redisConnectHandler net.Conn
}

func NewRedisDriver() *RedisDriver {
	var d RedisDriver
	d.redisServerIp = ""
	d.redisServerPort = ""
	d.redisPass = ""
	d.redisUser = ""
	d.redisConnectHandler = nil
	return &d
}

func (p *RedisDriver) ConnectDbDriver(connectUrl string) bool {
	//解析连接字符串，从中解析出连接URL，端口以及鉴权信息
	var splitRs []string = strings.split(connectUrl, ";")
	if len(splitRs) <= 0 {
		fmt.Println("Redis connect url error, connectUrl=", connectUrl)
		return false
	}

	for idx := 0; idx < len(splitRs); idx++ {
		if strings.Contains(splitRs[idx], "server=") {
			p.redisServerIp = splitRs[idx][7:]
			continue
		}
		if strings.Contains(splitRs[idx], "port=") {
			p.redisServerPort = splitRs[idx][5:]
			continue
		}
		if strings.Contains(splitRs[idx], "user=") {
			p.redisUser = splitRs[idx][5:]
			continue
		}
		if strings.Contains(splitRs[idx], "pass=") {
			p.redisPass = splitRs[idx][5:]
			continue
		}
	}

	//对redis而言，connectUrl是一个url:port格式的连接字符串
	var err error
	p.redisConnectHandler, err = net.dial("tcp", p.redisServerIp+":"+p.redisServerPort)
	if err != nil {
		fmt.Println("Connect to redis server error, error info=", err)
		return false
	}

	if len(p.redisPass) > 0 {
		rediscmd := fmt.Sprint("auth %s", p.redisPass)
		err = p.runRedisCommand(rediscmd)
		if err != nil {
			fmt.Println("Run <Auth> command error")
			p.redisConnectHandler.Close()
			p.redisConnectHandler = nil
			return false
		}
	}

	fmt.Println("Connect redis server success")
	return true
}

func (p *RedisDriver) DisconnectDbDriver() bool {
	return true
}

/*
	objectName描述表名
	objectValue描述该表中的记录，objectValue中携带了列名称

	redis存储表按照如下方式进行
	1:用objectValue进行hash得到一个取值X，X的范围为一个32位描述的数值，相同的行具有相同的hash值
	2:构造一个set，以objectName为key，将X保存到该set集合中
	3:以objectName:X为key，将objectValue保存进行hset中
*/
func (p *RedisDriver) AddRecord(objectName string, objectValue []TobjectType) error {

	crc32Sum := crc32.ChecksumIEEE([]byte(objectValue))
	var rediscmd string
	var err error
	rediscmd = fmt.Sprint("sadd %s %d", objectName, crc32Sum)
	err = p.runRedisCommand(rediscmd)
	if err != nil {
		fmt.Print("Run redis command <%s> error\r\n", rediscmd)
		return err
	}
	for idx := 0; idx < len(objectValue); idx++ {
		rediscmd = fmt.Sprint("hset %s:%d %s %s", objectName, crc32Sum, objectValue[idx].objectKey, objectValue[idx].objectVal)
		err = p.runRedisCommand(rediscmd)
		if err != nil {
			fmt.Print("Run redis command <%s> error\r\n", rediscmd)
			return err
		}
	}
	return nil
}

/*
	objectName描述表名
	objectValue描述该表中的记录
	updateCondition描述更新条件

	objectValue中包括完整的一条记录
*/
func (p *RedisDriver) UpdateRecord(objectName string, objectValue []TobjectType, opCondition []TobjectType) error {

	return nil
}

func (p *RedisDriver) RemoveRecord(objectName string, opCondition []TobjectType) error {
	return nil
}

func (p *RedisDriver) QueryRecord(objectName string, opCondition []TobjectType) (string, error) {
	return "", nil
}

func (p *RedisDriver) runRedisCommand(cmd string) (string, error) {

}
