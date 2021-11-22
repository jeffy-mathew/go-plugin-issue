package main

import (
	"log"
	"plugin"
)


func main()  {
	log.Println("Application started")
	loadAndCallPlugin("plugin.so", "PluginFunc")
	loadAndCallPlugin("plugin.so", "Greet")
}

func loadAndCallPlugin(pluginName, funcName string) {
	log.Printf("loading plugin : %s \n symbol: %s", pluginName, funcName)
	loadedPlugin, err := plugin.Open(pluginName)
	if err != nil {
		log.Printf("plugin loading failed %v", err)
	}
	fnSymbol, err := loadedPlugin.Lookup(funcName)
	if err != nil {
		log.Printf("symbol lookup failed %v", err)
	}
	pluginFunc := fnSymbol.(func(pluginInputNumber int))
	pluginFunc(10)
	log.Printf("loaded plugin : %s \n symbol: %s", pluginName, funcName)
}