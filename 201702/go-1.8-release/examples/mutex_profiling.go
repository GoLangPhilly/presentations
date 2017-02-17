go test bench=. -mutexprofile=mutex.out

go tool pprof runtime.test mutex.out

