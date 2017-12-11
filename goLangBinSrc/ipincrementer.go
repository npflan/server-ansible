package main

import (
	"os"
	"github.com/asaskevich/govalidator"
	"log"
	"net"
	"encoding/binary"
	"strconv"
	"fmt"
)

func main(){
	if len(os.Args) == 3 {
		var ip string = os.Args[1]
		var incrementAmount string = os.Args[2]

		if !govalidator.IsIPv4(ip){
			log.Fatal("ip is invalid")
		}
		if !govalidator.IsNumeric(incrementAmount){
			log.Fatal("incrementAmount is not numeric")
		}
		uint64IncrementAmount, err := strconv.ParseUint(incrementAmount, 10, 32)


		if err != nil{
			log.Fatal(err)
		}
		uint32IncrementAmount := uint32(uint64IncrementAmount)

		realIP := net.ParseIP(ip)

		ipUint32 := ip2int(realIP)

		ipUint32 += uint32IncrementAmount

		fmt.Print(int2ip(ipUint32))
	}else{
		printUsage()
	}
}

func printUsage(){
	fmt.Println(os.Args[0]+" startip incrementAmount")
	fmt.Println("Example: "+os.Args[0]+" 192.168.0.0 10")
}


func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
