package main

import (
    "cdht/CoreModule"
    "cdht/Config"
    "sync"
    "fmt"
    "time"
)


func main() {
    var wg sync.WaitGroup
    
    fmt.Println(logo)

    wg.Add(1)
    go testNode(&wg);
    wg.Wait()


}


func testNode(wg *sync.WaitGroup) {
    configMngr := Config.Config{}
    config := configMngr.LoadConfig()
    core := CoreModule.Core{ Config: &config }

    core.START()

    time.Sleep(time.Second * 30000)
    wg.Done()
}




var logo  string = `
                                                ▄████▄  ▓█████▄  ██░ ██ ▄▄▄█████▓
                                                ▒██▀ ▀█  ▒██▀ ██▌▓██░ ██▒▓  ██▒ ▓▒
                                                ▒▓█    ▄ ░██   █▌▒██▀▀██░▒ ▓██░ ▒░
                                                ▒▓▓▄ ▄██▒░▓█▄   ▌░▓█ ░██ ░ ▓██▓ ░ 
                                                ▒ ▓███▀ ░░▒████▓ ░▓█▒░██▓  ▒██▒ ░ 
                                                ░ ░▒ ▒  ░ ▒▒▓  ▒  ▒ ░░▒░▒  ▒ ░░   
                                                  ░  ▒    ░ ▒  ▒  ▒ ░▒░ ░    ░    
                                                ░         ░ ░  ░  ░  ░░ ░  ░      
                                                ░ ░         ░     ░  ░  ░         
                                                ░         ░
`