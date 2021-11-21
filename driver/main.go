package main

import (
	"log"
	"plugin"
	"sync"
	"fmt"
)


func main()  {
	log.Println("Application started")
	wg := sync.WaitGroup{}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go loadAndCallPlugin(fmt.Sprintf("./plugin%s.so", i), &wg)
	}
	wg.Wait()
}

func loadAndCallPlugin(pluginName string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("loading plugin ", pluginName)
	loadedPlugin, err := plugin.Open(plugin)
	if err != nil {
		log.Printf("plugin loading failed %v", err)
	}
	fnSymbol, err := loadedPlugin.Lookup("KnownPluginFunction")
	if err != nil {
		log.Printf("symbol lookup failed %v", err)
	}
	pluginFunc := fnSymbol.(func(pluginInputNumber int))
	pluginFunc(10)
	log.Println("loaded plugin ", pluginName)
}