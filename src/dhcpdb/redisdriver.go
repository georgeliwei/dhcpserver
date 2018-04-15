package dhcpdb

import (
	"fmt"
	"hash/crc32"
	"net"
	//"strings"
)

type RedisDriver struct {
	redisServerIp       string
	redisServerPort     string
	redisConnectHandler net.Conn
}

func NewRedisDriver() *RedisDriver {
	var d RedisDriver
	d.redisServerIp = ""
	d.redisServerPort = ""
	d.redisConnectHandler = nil
	return &d
}

func (p *RedisDriver) ConnectDbDriver(connectUrl string) bool {

	//对redis而言，connectUrl是一个url:port格式的连接字符串
	p.redisConnectHandler = net.dial("tcp", connectUrl)

	return true
}

func (p *RedisDriver) DisconnectDbDriver() bool {
	return true
}

func (p *RedisDriver) AddRecord(objectName string, objectValue []TobjectType) error {
	//redis存储表按照如下方式进行
	/*
		1:用objectValue进行hash得到一个取值X，X的范围为一个32位描述的数值
		2:构造一个set，以objectName为key，将X保存到该set集合中
		3:以objectName:X为key，将objectValue保存进行hset中
	*/
	crc32Sum := crc32.ChecksumIEEE([]byte(objectValue))
	var rediscmd string
	var err error
	rediscmd = fmt.Sprint("zadd %s %d %d", objectName, crc32Sum, crc32Sum)
	err = p.runRedisCommand(rediscmd)
	if err != nil {
		fmt.Print("Run redis command <%s> error\r\n", rediscmd)
		return err
	}
	for idx := 0; idx < len(objectValue); idx++ {
		rediscmd = fmt.Sprint("hset %s:%d %s %s", objectValue[idx].objectKey, objectValue[idx].objectVal)
		err = p.runRedisCommand(rediscmd)
		if err != nil {
			fmt.Print("Run redis command <%s> error\r\n", rediscmd)
			return err
		}
	}
	return nil
}

func (p *RedisDriver) UpdateRecord(objectName string, objectValue []TobjectType) error {
	return nil
}

func (p *RedisDriver) RemoveRecord(objectName string, objectValue []TobjectType) error {
	return nil
}

func (p *RedisDriver) QueryRecord(objectName string, objectValue []TobjectType) (string, error) {
	return "", nil
}

func (p *RedisDriver) runRedisCommand(cmd string) error {

}
