package main

import(
  "fmt"
  "log"
  "os/exec"
)
  
func GetCmd() []string {
    return []string{"ifconfig", "-a"}
}

func main() {

     cmd := GetCmd()
     out, err := exec.Command(cmd[0], cmd[1:]...).Output()
    
    // out, err := exec.Command("date").Output()
	 if err != nil {
	      log.Fatal(err)
     }
	 
	 fmt.Printf("The date is %s\n", out)
}

