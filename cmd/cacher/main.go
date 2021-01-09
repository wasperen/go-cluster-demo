package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/memberlist"
)

func main() {
	list, err := memberlist.Create(memberlist.DefaultLocalConfig())
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	if len(os.Args[1:]) > 0 {
		fmt.Printf("joining list via known member(s) %s\n", os.Args[1:])
		_, err = list.Join(os.Args[1:])
		if err != nil {
			panic("Failed to join cluster: " + err.Error())
		}
	}

	for {
		fmt.Printf("============= members at %s\n", time.Now().Format(time.RFC3339))
		for _, member := range list.Members() {
			fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
		}
		fmt.Printf("=============\n")
		time.Sleep(time.Second * 5)
	}
}
