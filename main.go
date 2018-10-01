package main

import(
	"net/http"
	"log"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const VERSION = 0.1

func main(){
	greetings()
	devices := newDevices("192.168.1.1/24")
	http.HandleFunc("/api/devices", devices.devicesHandler)
	http.HandleFunc("/api/verification", verificationHandler)
	fmt.Println("Starting server on http://127.0.0.1")
	fmt.Println("Press Ctrl + C to quit...")
	log.Fatal(http.ListenAndServe(":80", nil))
}


func greetings(){
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
logo :=
"               ,^,\n" +
"             ,'   `,\n" +
"           .'       `,\n" +
"         .`:  1   0  :`.\n" +
"           :         :\n" +
"           :  1   1  :\n" + 
"           ...........\n" +
"       <Moonshard Project>\n" +
"  Host backend version: "
	fmt.Print(logo)
	fmt.Println(VERSION)
}






