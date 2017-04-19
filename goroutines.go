package main

import "time"

func say(s string) {
	for i:= 0; i < 5; i++ {
		time.Sleep(100*time.Millisecond)
		println(s)
	}
}

func say2(s string) {
	println("test")
}

func main() {
	go say("World")
	say2("Hello")

}
