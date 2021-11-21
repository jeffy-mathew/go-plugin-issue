This repository tries to reproduce `plugin already loaded` issue.
When the same plugin is build and loaded once again from the driver program,
it results in a `plugin already loaded error`. 

Step to follow,

    ./setup.sh
    cd driver
    ./main

Current output 

```
2021/11/22 01:02:32 Application started
2021/11/22 01:02:32 loading plugin  plugin1.so
My custom code prints InputNumber: 10
2021/11/22 01:02:32 loaded plugin  plugin1.so
2021/11/22 01:02:32 loading plugin  plugin2.so
2021/11/22 01:02:32 plugin loading failed plugin.Open("plugin2"): plugin already loaded
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x28 pc=0x52f0b4]

goroutine 1 [running]:
plugin.lookup(...)
	/usr/local/go/src/plugin/plugin_dlopen.go:138
plugin.(*Plugin).Lookup(...)
	/usr/local/go/src/plugin/plugin.go:40
main.loadAndCallPlugin(0x55c2ec, 0xa)
	/home/jeff/works/go-plugin-issue/driver/main.go:22 +0x154
main.main()
	/home/jeff/works/go-plugin-issue/driver/main.go:13 +0xb
```

Expected output
```
2021/11/22 01:02:32 Application started
2021/11/22 01:02:32 loading plugin  plugin1.so
My custom code prints InputNumber: 10
2021/11/22 01:02:32 loaded plugin  plugin1.so
2021/11/22 01:02:32 loading plugin  plugin2.so
My custom code prints InputNumber: 10
2021/11/22 01:02:32 loaded plugin  plugin2.so
```