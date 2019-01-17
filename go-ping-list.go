package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"sort"
	"strconv"
	"sync"
	"github.com/sparrc/go-ping"
)

var usage = `
Usage:
    go-ping-list [-s start IP] [-e end IP]
Example:
    # go-ping-list -s 172.20.13.1 -e 172.20.13.254
`

type pingResult struct {
	msg string
	ip  net.IP
}

//worker for ping ip-addresses
func worker(workerNumber int, ipAddress <-chan string, workerResult chan<- pingResult) {
	for job := range ipAddress {

		pinger, err := ping.NewPinger(job)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}

		//display output options
		pinger.OnFinish = func(stats *ping.Statistics) {
			packetsloss := fmt.Sprintf("%v", stats.PacketLoss)
			packetslossfloat, err := strconv.ParseFloat(packetsloss, 64)
			if err == nil {
			}

			zero := float64(0)
			hundred := float64(100)

			var result string

			//if packets loss in ping function = 0 cycle print "HOST ACTIVE" else "HOST DISABLED"
			if packetslossfloat == zero {
				//	fmt.Printf("%s, %v%% packet loss, %s\n", pinger.Addr(), stats.PacketLoss, "HOST ACTIVE")
				result = fmt.Sprintf("%s %s", pinger.Addr(), "- HOST ACTIVE")
			} else if packetslossfloat < hundred {
				result = fmt.Sprintf("%s %s, %v%% packet loss", pinger.Addr(), "- HOST WARNING", stats.PacketLoss)
			} else if packetslossfloat == hundred {
				//	fmt.Printf("%s, %v%% packet loss, %s\n", pinger.Addr(), stats.PacketLoss, "HOST DISABLED")
				result = fmt.Sprintf("%s %s", pinger.Addr(), "- HOST DISABLED")
			}

			workerResult <- pingResult{
				msg: result,
				ip:  net.ParseIP(job),
			}
		}

		//ping options
		pinger.Count = 3
		pinger.Timeout = 10000000000

		//this for work on windows 10 or linux with sudo
		pinger.SetPrivileged(true)

		//run ping
		pinger.Run()
	}
}

//function for parse input flags and counting number goroutines from number IP-addresses in enter subnets
func f1() {
	startIPflag := flag.String("s", "172.20.13.0", "start IP")
	endIPflag := flag.String("e", "172.20.13.254", "end IP")

	helpFlag := flag.String("h", "", "")
	helpFlag1 := flag.String("help", "", "")

	flag.Usage = func() {
		fmt.Printf(usage)
	}

	//parse input flags
	flag.Parse()

	if *helpFlag == "h" {
		flag.Usage()
		return
	}

	if *helpFlag1 == "help" {
		flag.Usage()
		return
	}

	startIP := net.ParseIP(*startIPflag)
	endIP := net.ParseIP(*endIPflag)

	startIP = startIP.To4()
	endIP = endIP.To4()
	maxIP := int(endIP[3])

	countIPaddresses := int(endIP[3] - startIP[3])

	if startIP[0] != endIP[0] {
		fmt.Printf("Error! Start and end IP's have different subnets\n")
		return
	} else if startIP[1] != endIP[1] {
		fmt.Printf("Error! Start and end IP's have different subnets\n")
		return
	} else if startIP[2] != endIP[2] {
		fmt.Printf("Error! Start and end IP's have different subnets\n")
		return
	} else if maxIP > 254 {
		fmt.Printf("Error! Max end IP-address x.x.x.254\n")
		return
	}


	var pingResults []pingResult

	ipAddress := make(chan string)
	workerResult := make(chan pingResult)

	//start workers functions
	var workersWaitGroup sync.WaitGroup
	for workerNumber := 0; workerNumber <= countIPaddresses; workerNumber++ {
		workersWaitGroup.Add(1)

		go func(workerNumber int) {
			worker(workerNumber, ipAddress, workerResult)
			workersWaitGroup.Done()
		}(workerNumber)
	}

	//start reader results
	var readerWaitGroup sync.WaitGroup
	readerWaitGroup.Add(1)
	go func() {
		for r := range workerResult {
			pingResults = append(pingResults, r)
		}
		readerWaitGroup.Done()
	}()


	//parse input IP-addresses and transfer to worker function
	for i := startIP[3]; i <= endIP[3]; i++ {
		octet1 := startIP[0]
		octet2 := startIP[1]
		octet3 := startIP[2]

		output := fmt.Sprintf("%d.%d.%d.%d", octet1, octet2, octet3, i)
		ipAddress <- output
	}



	close(ipAddress)
	workersWaitGroup.Wait()

	close(workerResult)
	readerWaitGroup.Wait()


	//sort results IP-addresses and print results
	sort.Slice(pingResults, func(i, j int) bool {
		return bytes.Compare(pingResults[i].ip, pingResults[j].ip) < 0
	})
	for _, r := range pingResults {
		fmt.Println(r.msg)
	}
}

func main() {
	f1()
}
