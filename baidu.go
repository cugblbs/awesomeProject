package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
)

func main() {
	response,err:=http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Printf("%s", err)
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		fmt.Println(response.StatusCode)
	}

	buf := make([]byte, 1024)
	f, err1 := os.OpenFile("baidu.html", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err1 != nil {
		panic(err1)
		return
	}
	defer f.Close()
	for {
		n, _:=response.Body.Read(buf)
		if 0 == n {
			break
		}
		f.WriteString(string(buf[:n]))
	}
}
