/* Get nodes in consul cluster*/

package main

import "github.com/hashicorp/consul/api"
import "fmt"

func main() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	node := client.Agent()

	leader, err := node.NodeName()
	fmt.Println("Leader = ", leader)
	members, err := node.Members(false)
	if err != nil {
		panic(err)
	}

	// https://godoc.org/github.com/hashicorp/consul/api#AgentMember
	fmt.Println("Consul Members:")
	for _, member := range members {
		fmt.Println(member.Name, member.Addr, member.Status)
	}
}
