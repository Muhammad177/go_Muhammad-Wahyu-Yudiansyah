package main

import (
    "fmt"
    "sync"
)

func main() {
    text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
    freq := make(map[rune]int)
    var wg sync.WaitGroup
    var mu sync.Mutex

    for _, ch := range text {
        wg.Add(1)
        go func(ch rune) {
            defer wg.Done()
            mu.Lock()
            freq[ch]++
            mu.Unlock()
        }(ch)
    }

    wg.Wait()

    for ch, count := range freq {
        fmt.Printf("%c: %d\n", ch, count)
    }
}
