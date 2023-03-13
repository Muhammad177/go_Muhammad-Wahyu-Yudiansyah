package main

import "fmt"

func getMinServers(expected_load int32, servers []int32) int32 {

    n := int32(len(servers))
    var minServers int32 = n + 1
    var totalCapacity int32 = 0

    for _, server := range servers {
        totalCapacity += server
    }

    if totalCapacity < expected_load {
        return (expected_load - totalCapacity)
    }

    for i := int32(0); i < (1 << n); i++ {
        var count int32 = 0
        for j := int32(0); j < n; j++ {
            if i&(1<<j) > 0 {
                count++
            }
        }

        var selectedCapacity int32 = 0
        for j := int32(0); j < n; j++ {
            if i&(1<<j) > 0 {
                selectedCapacity += servers[j]
            }
        }

        if selectedCapacity == expected_load && count < minServers {
            minServers = count
        }
    }

    if minServers == n+1 {
        return -1
    }
    
    if minServers == 0 {
        return -1
    }
    
    return (totalCapacity - expected_load)
}

func main() {

    servers := []int32{1,1, 2, 4,4}
    expected_load := int32(10)

    minServers := getMinServers(expected_load, servers)
    if minServers == -1 {
        fmt.Println("Cannot find a combination of servers to handle expected load")
    } else if minServers == 0 {
        fmt.Println("The total capacity of servers is already sufficient to handle expected load")
    } else {
        fmt.Println("Need", minServers, "more servers to handle expected load")
    }
}
