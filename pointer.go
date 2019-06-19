package main

import (
	"bufio"
	"bytes"
	"io"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

//	A MAC address is 6 bytes
type MacAddress [6]byte

// A MagicPacket is constituted of 6 bytes of 0xFF followed by
// 16 groups of the destination MAC address.
type MagicPacket struct{
	header	[6]byte
	payload	[16]MacAddress
}

func main() {
	//	get network interfaces currently available
	interfaces, err := net.Interfaces()
	//	error handling
	if err != nil {
		fmt.Println("Error in detecting network interfaces: " + err.Error())
		os.Exit(0)
	}
	if len(interfaces) > 0 {
		fmt.Println("\nList of available network interfaces:")
		for index, i := range interfaces {
			fmt.Printf("%d.%v\n", index, i.Name)
		}
		fmt.Printf("\nSelect the interface that you want to use: (0-%d): ", len(interfaces))
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(char)
	}
	wakeOnLan("127.0.0.1")

}

func wakeOnLan(ip string) {
	//	 port to connect to
	var port_num int
	//	split the ip string to check for any port number
	ip_addr := strings.Split(ip, ":")
	//	if port is given, store it in port_num or
	//	else assign port_num a default port value (UDP 9)
	if len(ip_addr) > 1 {
		//	convert from string to integer
		port_num, _ = strconv.Atoi(ip_addr[1])
	} else {
		//	default UDP port number for WakeOnLan service
		port_num = 9
	}
	//	resolve the IP address
	//	here im just concatenating the ip address with the port number ip:port
	addr, err := net.ResolveUDPAddr("ip", ip_addr[0]+":"+strconv.Itoa(port_num))
	if err != nil {
		fmt.Println("Unable to get a UDP address for %s\n", addr,err.Error())
		os.Exit(1)
	}
	//	connect to the IP address
	//	 keep the local address nil
	conn,err := net.DialUDP("udp",nil,addr)
	if err != nil{
		fmt.Println("Could not connect to the destination: %s\n",addr,err.Error())
		os.Exit(1)
	}
	//	when all is done, close the connection
	defer conn.close()


	fmt.Println(addr.IP)
	fmt.Println(port_num)
}
