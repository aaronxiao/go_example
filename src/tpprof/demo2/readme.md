Go 大杀器之性能剖析 PProf

https://www.topgoer.cn/docs/jianyugo/jianyugo-1cl3tbo6fu6qv

go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60

go tool pprof http://localhost:6060/debug/pprof/heap

go tool pprof http://localhost:6060/debug/pprof/block

go tool pprof http://localhost:6060/debug/pprof/mutex

go tool pprof -http=:8080 cpu.prof

go tool pprof cpu.prof