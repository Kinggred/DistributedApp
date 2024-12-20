package networking

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	conf "tic-tac-toe/core/config"
	log "tic-tac-toe/core/config/logging"
)

func getLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.ClientLogger.Fatalf("IP cannot be retrieved")
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	log.ClientLogger.Printf("Local address found: %s", localAddress.String())
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
	var mu sync.Mutex
	var wg sync.WaitGroup
	localAddress := strings.Split(getLocalIP().String(), ":")[0]
	subnet := getSubnet(localAddress)
	// Just spent 3 days debugging not being able to see the servers.
	// Forgot about this shit, setting this below 150ms works basically
	timeout := 200 * time.Millisecond

	for i := 1; i < 254; i += 5 {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			for j := start; j < start+5 && j < 254; j++ {
				ip := fmt.Sprintf("%s%d", subnet, j)
				if scanPort(ip, 50051, timeout) && (ip != localAddress || (conf.CONFIG.DEBUG && ip == localAddress)) {
					log.ClientLogger.Printf("Active device found at %s", ip)
					mu.Lock()
					hostList = append(hostList, ip)
					mu.Unlock()
				}
			}
		}(i)
	}
	wg.Wait()
	log.ClientLogger.Printf("Sweep done, no. of hosts found: %d", len(hostList))
	return hostList
}
