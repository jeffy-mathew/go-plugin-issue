package main

import "fmt"

func init() {
	fmt.Println("initialized plugin")
}
func PluginFunc(pluginInputNumber int) { fmt.Printf("My custom code prints PluginInputNumber: %d\n", pluginInputNumber) }

func Greet(pluginInputNumber int) {fmt.Printf("hello!! greetings!! x %d", pluginInputNumber)}

func main() {}