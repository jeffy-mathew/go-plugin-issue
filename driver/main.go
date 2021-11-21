package main

import (
	"log"
	"plugin"
)


func main()  {
	log.Println("Application started")
	log.Println("loading plugin1")
	plugin1, err := plugin.Open("./plugin1.so")
	if err != nil {
		log.Printf("plugin loading failed %v", err)
	}
	fnSymbol1, err := plugin1.Lookup("KnownPluginFunction")
	if err != nil {
		log.Printf("symbol lookup failed %v", err)
	}
	plugin1Func := fnSymbol1.(func(pluginInputNumber int))
	plugin1Func(10)
	log.Println("finished loading plugin1")
	log.Println("loading plugin second time")
	plugin2, err := plugin.Open("./plugin1.so")
	if err != nil {
		log.Printf("plugin loading failed %v", err)
	}
	fnSymbol2, err := plugin2.Lookup("KnownPluginFunction")
	if err != nil {
		log.Printf("symbol lookup failed %v", err)
	}
	plugin2Func := fnSymbol2.(func(pluginInputNumber int))
	plugin2Func(10)
	log.Println("finished loading plugin2")
}