package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

var sb strings.Builder

func append(data string) {
	f, _ := os.OpenFile("out.txt", os.O_APPEND|os.O_WRONLY, 0644)
	f.WriteString(data)
	f.Close()

}

func IsOpened(host string, port int) bool {

	timeout := 10 * time.Second
	target := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return false
	}

	if conn != nil {
		conn.Close()
		return true
	}

	return false
}

func check(ip string, port int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	openport := strconv.Itoa(port)
	if IsOpened(ip, port) {
		sb.WriteString(ip + ":" + openport + "\n")
		append(ip + ":" + openport + "\n")
	}
}

func main() {

	filename := filepath.Base(os.Args[0])
	if len(os.Args) != 4 {
		fmt.Println("PortScanner\nUsage: " + filename + " <ip> <start port> <end port>")
		os.Exit(0)

	}
	host := os.Args[1]
	min, _ := strconv.Atoi(os.Args[2])
	max, _ := strconv.Atoi(os.Args[3])
	max = max + 1
	var wg sync.WaitGroup
	var a int

	for i := min; i < max; i++ {
		a++
		fmt.Printf("Scanning: %s:%d\n", host, i)
		wg.Add(1)

		go check(host, i, &wg, a)
		time.Sleep(1000 * time.Nanosecond)
	}
	wg.Wait()

	for i := 1; i < 1500; i++ {
		fmt.Println("")
	}
	fmt.Print(sb.String())

}
