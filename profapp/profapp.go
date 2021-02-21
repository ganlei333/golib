package profapp

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"time"

	"go.uber.org/zap"
)

func StartCpuProf() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("create cpu profile file error: ", zap.Error(err))
		return
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("can not start cpu profile,  error: ", zap.Error(err))
		f.Close()
	}
}

func StopCpuProf() {
	pprof.StopCPUProfile()
}

//--------Mem
func ProfGc() {
	runtime.GC() // get up-to-date statistics
}

func SaveMemProf() {
	f, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("create mem profile file error: ", zap.Error(err))
		return
	}

	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Println("could not write memory profile: ", zap.Error(err))
	}
	f.Close()
}

// goroutine block
func SaveBlockProfile() {
	f, err := os.Create("block.prof")
	if err != nil {
		fmt.Println("create mem profile file error: ", zap.Error(err))
		return
	}

	if err := pprof.Lookup("block").WriteTo(f, 0); err != nil {
		fmt.Println("could not write block profile: ", zap.Error(err))
	}
	f.Close()
}

// 生成 CPU 报告
func CpuProfile() {
	f, err := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Println("CPU Profile started")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	time.Sleep(60 * time.Second)
	fmt.Println("CPU Profile stopped")
}

// 生成堆内存报告
func HeapProfile() {
	f, err := os.OpenFile("heap.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	time.Sleep(30 * time.Second)

	pprof.WriteHeapProfile(f)
	fmt.Println("Heap Profile generated")
}

// 生成追踪报告
func TraceProfile() {
	f, err := os.OpenFile("trace.out", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Println("Trace started")
	trace.Start(f)
	defer trace.Stop()

	time.Sleep(60 * time.Second)
	fmt.Println("Trace stopped")
}
