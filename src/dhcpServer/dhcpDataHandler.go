package main

/*
func handleConnect(con net.Conn) int {
	var readBuf []byte
	readBuf = make([]byte, 4096, 4096)
	for true {
		readNum, err := con.Read(readBuf)
		if err == nil {
			go HandlerDhcpMessage(readBuf, readNum)
			continue
		}
		if err == io.EOF {
			fmt.Println("Connect closed by peer")
			return 0
		} else {
			fmt.Println("Read data from connect error, err=", err)
			return -1
		}
	}
	return 0
}

func listenAndAccept() {
	serverAddr, err := net.ResolveUDPAddr("udp", ":67")
	if err != nil {
		fmt.Println("Resolve address error, err=", err)
		return
	}
	listenCon, err := net.ListenUDP("udp", serverAddr) //DHCP监听67端口
	if err != nil {
		fmt.Println("Listen dhcp port(67) error, err=", err)
		return
	}
	defer listenCon.Close()
	for true {
		//读取数据
		handleConnect(listenCon)
	}
}

*/
