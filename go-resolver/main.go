package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {

	// hostname to resolve
	hostnames := []string{"danrl.com", "www.egym.de"}

	// create resolver
	resolver := net.Resolver{
		PreferGo: true,
	}

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	for _, hostname := range hostnames {
		// lookup host
		fmt.Printf("looking up `%v`\n", hostname)
		literals, err := resolver.LookupHost(ctx, hostname)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// print results
		for _, literal := range literals {
			fmt.Printf("* %v\n", literal)
		}
	}

}
