package memorandum

import (
    "context"
    "os/exec"
    "fmt"
)

func StartVue(){
    fmt.Println("Vue Start, Listen on localhost:8080")
    ctx, cancel := context.WithCancel(context.Background())
    cmd := exec.CommandContext(ctx, "npm", "run", "dev")
    if err := cmd.Start(); err != nil{
        fmt.Println(err)
    }
    time.Sleep(time.Second * 10)
    fmt.Println("Process quit")
    cancel()
    cmd.Wait()
}

func StopVue(){

}
