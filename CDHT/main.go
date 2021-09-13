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

    if core.Config.RUN_TCP_TEST_APPLICATION || core.Config.RUN_UDP_TEST_APPLICATION { 
        time.Sleep(time.Minute * core.Config.TEST_APP_RUNNING_TIME)
    }
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