package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"
)

//  from https://apple.stackexchange.com/questions/310061/how-to-find-all-devices-ip-address-hostname-mac-address-on-local-network

func main() {

	// get primary interface
	pIF, err := execWithBash("echo \"show State:/Network/Global/IPv4\" | scutil | awk -F: '/PrimaryInterface/{sub(/ /,\"\",$2); print $2}'")

	// get subnet of ip for primary interface (assuming /24)
	// sn=$(ipconfig getifaddr $pIF | sed -En 's/^([0-9]+\.[0-9]+\.[0-9]+).*/\1/p')

	sn, err := execWithBash(fmt.Sprintf("ipconfig getifaddr %s | sed -En 's/^([0-9]+\\.[0-9]+\\.[0-9]+).*/\\1/p'", pIF))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Primary Interface: |%s|\n", pIF)
	fmt.Printf("           Subnet: %s\n", sn)

	//  ping in parallell
	var wg sync.WaitGroup
	for i := 0; i <= 254; i++ {
		ip := fmt.Sprintf("%s.%d", sn, i)
		time.Sleep(50 * time.Millisecond)
		wg.Add(1)
		go func(msg string) {
			defer wg.Done()
			ping(ip)
		}(ip)
	}
	wg.Wait()

	arp(pIF)
}

func arp(pIF string) error {
	// arp -a | grep $pIF | sed -e 's/^\?/unnamed/' -e "s/\ at\ /${tab}/g" -e "s/\ on\ /${tab}/g" -e 's/\ ifscope.*$//g' | awk 'BEGIN { FS="\t"; OFS="\t"; printf "%-17s\t%-15s\t%s\n", "MAC","INTERFACE","HOSTNAME (IP)" } { if($2!="(incomplete)") {printf "%-17s\t%-15s\t%s\n",$2,$3,$1}}'

	arp, err := execWithBash(fmt.Sprintf("arp -a | grep %s |grep -v incomplete", pIF))
	if err == nil {
		fmt.Printf("           ARP:\n%s\n", arp)
	}
	scanner := bufio.NewScanner(strings.NewReader(arp))
	for scanner.Scan() {
		fmt.Println(len(scanner.Bytes()), scanner.Text())
		extractIPAndMac(scanner.Text())
	}
	return err
}

func extractIPAndMac(line string) {
	var name string
	var ip string
	var mac string
	var ifc string

	// drobo5n.local (169.254.5.47) at 0:1a:62:4:4a:6c on en0 [ethernet]
	// ? (192.168.86.34) at 86:fa:56:29:19:4 on en0 ifscope [ethernet]
	n, err := fmt.Sscanf(line,
		"%s (%s) at %s on %s ifscope [ethernet]\n", &name, &ip, &mac, &ifc)

	if err != nil {
		fmt.Println("err!=mil", n, len(line), line)
		log.Fatal(err)
	} else if n != 3 {
		fmt.Println("n!=3", n, len(line), line)
		log.Fatal(err)
	}
	fmt.Printf("- %s:\n%s\n", mac, ip)
}

func ping(ip string) error {
	intvl_wait := "1" // time between pings (-c1 implies we are waiting for a response, but will resend after this interval)
	reply_wait := 100

	ping, err := execWithBash(fmt.Sprintf("ping -i%s -W%d -c1 %s | grep from", intvl_wait, reply_wait, ip))
	if (len(ping) > 0) && (err == nil) {
		fmt.Printf("           Ping: %s |%s|=%d\n", ip, ping, len(ping))
	}
	return err
}
func execWithBash(cmd string) (string, error) {
	bya, err := exec.Command("bash", "-c", cmd).Output()
	str := strings.TrimSuffix(string(bya), "\n")
	return str, err

}
