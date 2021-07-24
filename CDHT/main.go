package main;
import "fmt"

func main() {

    messages := []string{"hellp","0hellp0","world"}
    mapper :=make(map[int]string)

    testChanges(mapper, messages)
    fmt.Println( mapper, messages, myport )
}

func printStrs(messages ...string) (length int) {
    for _, mem := range(messages){
        fmt.Println( mem)
    }
    length = len(messages)
    return
}


func testChanges(mapper map[int]string, messages []string){
    mapper[0] = "hello man"
    messages[0] = "hello man"
}


type PORT int
const (
   myport PORT = 12

)