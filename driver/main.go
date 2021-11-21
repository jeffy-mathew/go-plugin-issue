package main

import (
	"log"
	"plugin"
)


func main()  {
	log.Println("Application started")
	loadAndCallPlugin("plugin1.so")
	loadAndCallPlugin("plugin1.so")
	loadAndCallPlugin("plugin2.so")
}

func loadAndCallPlugin(pluginName string) {
	log.Println("loading plugin ", pluginName)
	loadedPlugin, err := plugin.Open(pluginName)
	if err != nil {
		log.Printf("plugin loading failed %v", err)
	}
	fnSymbol, err := loadedPlugin.Lookup("PluginFunc")
	if err != nil {
		log.Printf("symbol lookup failed %v", err)
	}
	pluginFunc := fnSymbol.(func(pluginInputNumber int))
	pluginFunc(10)
	log.Println("loaded plugin ", pluginName)
}