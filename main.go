
package main

import (
	"fmt"
	"net"
)

func main() {

	iprecords, _ := net.LookupIP("google.com")
	for _, ip := range iprecords {
	fmt.Print(ip)
	}
}