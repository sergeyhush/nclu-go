package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

const NCLU_ADDR = "/var/run/nclu/uds"

func main() {
	conn, err := net.Dial("unix", NCLU_ADDR)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	cmdJson, _ := json.Marshal(os.Args[1:])
	fmt.Fprintf(conn, string(cmdJson))
	b, err := ioutil.ReadAll(bufio.NewReader(conn))
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(b))
}
