package networking

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	conf "tic-tac-toe/core/config"
)

func getLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatalf("IP cannot be retrieved")
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	log.Printf("Local address found: %s", localAddress.String())
	return localAddress.IP
}

func getSubnet(ip string) string {
	address := strings.Split(ip, ".")
	address[len(address)-1] = ""
	return strings.Join(address, ".")
}

func scanPort(ip string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func ScanNetwork() []string {
	var hostList []string
	localAddress := strings.Split(getLocalIP().String(), ":")[0]
	subnet := getSubnet(localAddress)
	timeout := 2 * time.Millisecond

	for i := 1; i < 254; i++ {
		ip := fmt.Sprintf("%s%d", subnet, i)
		if scanPort(ip, 50051, timeout) && (ip != localAddress || (conf.CONFIG.DEBUG && ip == localAddress)) {
			log.Printf("Active device found at %s", ip)
			hostList = append(hostList, ip)
		}
	}
	log.Printf("Sweep done, no. of hosts found: %d", len(hostList))
	return hostList
}
