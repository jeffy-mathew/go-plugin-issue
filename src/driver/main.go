package main

import (
	"log"
	"plugin"
)


func main()  {
	log.Println("Application started")
	loadAndCallPlugin("plugin1.so", "PluginFunc")
	loadAndCallPlugin("plugin2.so", "PluginFunc1")
}

func loadAndCallPlugin(pluginName, symbolName string) {
	loadedPlugin, err := plugin.Open(pluginName)
	if err != nil {
		log.Printf("plugin loading failed %v", err)
	}
	fnSymbol, err := loadedPlugin.Lookup(symbolName)
	if err != nil {
		log.Printf("symbol lookup failed %v", err)
	}
	pluginFunc := fnSymbol.(func(pluginInputNumber int))
	pluginFunc(10)
	log.Println("loaded plugin ", pluginName)
}