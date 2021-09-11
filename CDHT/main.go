package main

import (
    "cdht/CoreModule"
    "sync"
)


func main() {
    var wg sync.WaitGroup
    
    wg.Add(1)
    go testNode(&wg);
    wg.Wait()


}


func testNode(wg *sync.WaitGroup) {

    core := CoreModule.Core{}
    core.START()

    wg.Done()
}




