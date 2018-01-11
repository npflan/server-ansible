package main

import (
	"github.com/asaskevich/govalidator"
	"log"
	"net"
	"fmt"
	"flag"
	"os"
)

var ExcludeBroadcast = flag.Bool("--exclude-broadcast", false, "Exclude Broadcast Address")
var ExcludeNetID = flag.Bool("--exclude-netid", false, "Exclude NetID Address")

func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	
	if *ExcludeBroadcast && *ExcludeNetID {
		return ips[1 : len(ips)-1], nil
	}
	if *ExcludeBroadcast {
		return ips[: len(ips)-1], nil
	}
	if *ExcludeNetID {
		return ips[1:], nil
	}
	
	return ips, nil
}

//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func main(){
	var (
		loggerError = log.New(os.Stderr,"",0)
	)

	flag.Parse()
	if len(flag.Args()) >= 1{
		for _,cidr := range(flag.Args()) {

			if !govalidator.IsCIDR(cidr){
				loggerError.Println(fmt.Sprintf("ip '%s' is not a valid cidr notation", cidr))
			}else{
				hosts, err := Hosts(cidr)
				
				if err != nil {
					loggerError.Println(err)
				}else{
					for i,ip := range(hosts) {
						if i > 0 {
							fmt.Printf("\n")
						}
						fmt.Printf(ip)
					}
				}
			}
		}
	}else{
		flag.PrintDefaults()
		printUsage()
	}
}

func printUsage(){

}
