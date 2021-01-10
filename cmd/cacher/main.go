package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/memberlist"
)

const retries uint8 = 15

func becomeMember(list *memberlist.Memberlist, knownHosts []string) error {
	fmt.Printf("joining list via known member(s) %s\n", knownHosts)

	var err error
	var nrMembers int

	for retry := uint8(0); retry < retries; retry++ {
		fmt.Printf("- becoming member, try %d\n", retry)
		nrMembers, err = list.Join(knownHosts)
		if err == nil {
			fmt.Printf("connected with %d members\n", nrMembers)
			return nil
		}
		fmt.Printf(" -- no member yet with error %s\n", err.Error())
		time.Sleep(time.Second * 1)
	}
	return err
}

func main() {
	list, err := memberlist.Create(memberlist.DefaultLocalConfig())
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	if len(os.Args[1:]) > 0 {
		err = becomeMember(list, os.Args[1:])
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
