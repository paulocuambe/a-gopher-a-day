package main

import (
	"fmt"
	"net"
	"os"
)

func MustErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	fmt.Printf("LookupIP\n")
	ips, err := net.LookupIP("google.com")
	MustErr(err)

	for _, ip := range ips {
		fmt.Printf("%#v\n", ip.String())
	}

	fmt.Printf("\n\nLookupTXT\n")
	rs, err := net.LookupTXT("google.com")
	MustErr(err)

	for _, r := range rs {
		fmt.Printf("%#v\n", r)
	}

	fmt.Printf("\n\nLookupNS\n")
	nss, err := net.LookupNS("google.com")
	MustErr(err)

	for _, ns := range nss {
		fmt.Printf("%#v\n", ns)
	}

	fmt.Printf("\n\nLookupCNAME\n")
	cname, err := net.LookupCNAME("google.com")
	MustErr(err)

	fmt.Printf("%#v\n", cname)
}
