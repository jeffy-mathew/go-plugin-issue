package main

import "fmt"


func PluginFunc(pluginInputNumber int) { fmt.Printf("My custom code prints PluginInputNumber: %d\n", pluginInputNumber) }

func Greet(pluginInputNumber int) {fmt.Printf("hello!! greetings!! x %d", pluginInputNumber)}

func main() {}