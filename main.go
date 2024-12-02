package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "time"
)

func cpuIntensiveTask() {
    // Lower iteration count to reduce load
    for i := 0; i < 1_000_000; i++ {
        _ = rand.Intn(1000) * rand.Intn(1000)
    }
}

func memoryIntensiveTask() {
    // Reduce the size to avoid excessive memory consumption
    data := make([][]int, 100)
    for i := range data {
        data[i] = make([]int, 100_000) // Reduced size
        for j := range data[i] {
            data[i][j] = rand.Intn(1000)
        }
    }
    fmt.Println("Memory task completed.")
}

func main() {
    fmt.Println("Starting lighter resource-consuming tasks...")
    
    for i := 0; i < runtime.NumCPU()/2; i++ {
        go cpuIntensiveTask()
    }
    go memoryIntensiveTask()

    time.Sleep(30 * time.Second)
    fmt.Println("Resource consumption test completed.")
}
