package main

import (
	"fmt"
	"net"

	"github.com/miekg/dns"
)

func main() {

	// create client
	client := new(dns.Client)
	message := new(dns.Msg)

	message.SetQuestion("www.egym.de.", dns.TypeAAAA)
	message.RecursionDesired = true

	server := net.JoinHostPort("2001:4860:4860::8888", "53")
	r, _, err := client.Exchange(message, server)
	if err != nil {
		fmt.Println(err)
	}
	if r.Rcode != dns.RcodeSuccess {
		fmt.Println(err)
	}
	for _, records := range r.Answer {
		if record, ok := records.(*dns.AAAA); ok {
			fmt.Printf("%s\n", record.String())
		}
	}

}
