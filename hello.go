// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 1.

// Helloworld is our first Go program.
//!+
package main

import "fmt"

func main() {
        fmt.Println("Hello, 世界")
        
	var snmpConfigured bool = true
	var vlanID int = 1024
	if !snmpConfigured {
		fmt.Println("SNMP is configured!")
	} else {
		fmt.Println("SNMP is not configured!")
	}
	if vlanID < 100 {
		fmt.Println("VLAN ID is less than 100")
	} else if vlanID > 100 && vlanID < 1000 {
		fmt.Println("VLAN ID is between 100 and 1000")
	} else {
		fmt.Println("VLAN ID is greater than 1000")
	}

	i := 1
	for i <= 5 {
		fmt.Printf("VLAN %d\n", i)
		i++
	}

	for i := 1; i <= 5; i++ {
		fmt.Printf("VLAN %d\n", i)
	}

	var vlan [3]int
	vlan[0] = 1

	vlans2 := [3]int{1, 2, 3}

	fmt.Println(vlans2)

	//var intSlice []int
	//var stringSlice []string
	var vlanSlice = []int{11, 12, 13, 14, 15}
	//vlanSlice2 := []int{11, 22, 33, 44, 55}

	vlanSlice = append(vlanSlice, 66)
	fmt.Println(vlanSlice)
	fmt.Printf("vlanSlice cap is %d, len is %d \n", cap(vlanSlice), len(vlanSlice))
	vlanSlice = append(vlanSlice, 6)
	vlanSlice = append(vlanSlice, 76)
	vlanSlice = append(vlanSlice, 86)
	vlanSlice = append(vlanSlice, 96)
	vlanSlice = append(vlanSlice, 9)
	fmt.Printf("vlanSlice cap is %d, len is %d \n", cap(vlanSlice), len(vlanSlice))

	preaccocatedVlanSlice := make([]int, 2, 50)
	fmt.Printf("preallocatedVlanSlice cap is %d, len is %d\n", cap(preaccocatedVlanSlice), len(preaccocatedVlanSlice))
	preaccocatedVlanSlice[0] = 1
	preaccocatedVlanSlice[1] = 2
	for i := 3; i <= 50; i++ {
		preaccocatedVlanSlice = append(preaccocatedVlanSlice, i)
	}
	fmt.Printf("preallocatedVlanSlice cap is %d, len is %d\n", cap(preaccocatedVlanSlice), len(preaccocatedVlanSlice))
}

}
