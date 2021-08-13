# general
go test -run . -bench=. -benchtime=5s -count 5 -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out | tee bench.txt
go tool pprof -http :8080 cpu.out
go tool pprof -http :8080 mem.out
go tool trace trace.out
go tool pprof <yourbin> cpu.out
go get -u golang.org/x/perf/cmd/benchstat
benchstat bench.txt

# http pprof
import _ "net/http/pprof"
go tool pprof <yourbin> http://127.0.0.1:8080/debug/pprof/profile
go tool pprof -alloc_objects <yourbin> http://127.0.0.1:8080/debug/pprof/heap
go tool pprof <yourbin> http://127.0.0.1:8080/debug/pprof/heap
go tool pprof -alloc_objects <yourbin> http://127.0.0.1:8080/debug/pprof/heap
go tool pprof <yourbin> http://127.0.0.1:8080/debug/pprof/goroutine
go tool pprof <yourbin> http://127.0.0.1:8080/debug/pprof/block




