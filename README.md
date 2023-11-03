该项目为Go Pprof/Go viewcore工具调优实践案例，使用该项目可以快速完成pprof、viewcore的入门
项目已经内置改良版viewcore 在debug/cmd/viewcore下执行go build .即可

运行该项目至少需要2c2g的资源

程序会占用约 2G 内存；
程序占用 CPU 最高约 100%（总量按“核数 * 100%”来算）；
程序不涉及网络或文件读写；
程序除了吃资源之外没有其他危险操作

代码会在cpu、memory、mutex、goroutine、gc五个方面对资源进行压测

pprof默认暴露端口6060