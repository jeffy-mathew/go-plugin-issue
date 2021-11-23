This repository intends to reproduce `plugin already loaded` issue with golang plugin

To reproduce issue run `./run-static.sh`

Output

```
2021/11/23 16:14:03 Application started
2021/11/23 16:14:03 loaded plugin  plugin1.so
My custom code prints InputNumber: 10
2021/11/23 16:14:03 plugin loading failed plugin.Open("plugin2"): plugin already loaded
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x28 pc=0x52e2a5]

goroutine 1 [running]:
plugin.lookup(...)
        /usr/local/go/src/plugin/plugin_dlopen.go:138
plugin.(*Plugin).Lookup(...)
        /usr/local/go/src/plugin/plugin.go:40
main.loadAndCallPlugin(0x55b16b, 0xa, 0x55b32d, 0xb)
        /driver/main.go:20 +0xc5
main.main()
        /driver/main.go:12 +0xc5

```

This issue occurs when the plugins are built from same build directory `/go/src/plugin-build` in the plugin compiler.

This can be fixed if we keep the plugin source code in different directories 
while building the plugin using the compiler.

run `./run-dynamic.sh` where the plugin is built from dynamically generated directories.

Output
```
2021/11/23 16:21:16 Application started
2021/11/23 16:21:16 loaded plugin  plugin1.so
My custom code prints InputNumber: 10
2021/11/23 16:21:16 loaded plugin  plugin2.so
My custom code prints InputNumber: 10
```