package util

import (
	"github.com/FactomProject/factomd/util/atomic"
	"fmt"
	"time"
)

var (
	ThreadMutex       atomic.DebugMutex
	ThreadGoidToNames map[string]string // map Goid() to name
	ThreadNameToGoids map[string]string // map Goid() to name
	ThreadIds         map[string]int    // map name to id
	ThreadNames       []string
	ThreadLoopCount   []int
)

func init() {
	ThreadMutex.Lock()
	ThreadGoidToNames = make(map[string]string)
	ThreadNameToGoids = make(map[string]string)
	ThreadIds = make(map[string]int)
	ThreadNames = make([]string, 0)
	ThreadLoopCount = make([]int, 0)
	ThreadMutex.Unlock()
	go ThreadStallReport()
}

func ThreadStallReport() {
	var threadId = ThreadStart("ThreadStallReport")
	for {
		oldThreadLoopCount := make([]int, len(ThreadLoopCount))
		copy(oldThreadLoopCount, ThreadLoopCount) // save current size
		time.Sleep(10 * time.Second)
		ThreadLoopInc(threadId)
		ThreadMutex.Lock()

		for id, name := range ThreadNames {
			if ( id < len(oldThreadLoopCount) && oldThreadLoopCount[id] == ThreadLoopCount[id]) {
				count := ThreadLoopCount[id]
				goid := ThreadNameToGoids[name]
				ThreadMutex.Unlock()
				fmt.Printf("Stalled Thread %33s:%03d (%14s) no iterations in last 10 seconds, %6v total\n", name, id, goid, count)
				ThreadMutex.Lock()
			}
		}
		for id, name := range ThreadNames {
			count := ThreadLoopCount[id]
			goid := ThreadNameToGoids[name]
			ThreadMutex.Unlock()
			fmt.Printf("Thread %33s:%03d (%14s)  %6v itterations\n", name, id, goid, count)
			ThreadMutex.Lock()
		}


		ThreadMutex.Unlock()
	}
	ThreadStop(threadId)
}

func ThreadName() (name string) {
	ThreadMutex.Lock()
	name = ThreadGoidToNames[atomic.Goid()]
	ThreadMutex.Unlock()
	return name
}

func ThreadStart(name string) (id int) {
//	fmt.Printf("ThreadStart(%s)\n", name)
	goid := atomic.Goid()
	ThreadMutex.Lock()
	id, ok := ThreadIds[name]
	if (ok) {
		ThreadMutex.Unlock()
		panic("Can't start thread twice:" + name)
	}
	ThreadGoidToNames[goid] = name
	ThreadNameToGoids[name] = goid
	ThreadNames = append(ThreadNames, name)
	ThreadLoopCount = append(ThreadLoopCount, 0)
	id = len(ThreadNames)-1 // -1 so the first one gets Id 0
	ThreadIds[name] = id
	ThreadMutex.Unlock()
	fmt.Printf("ThreadStart(%33s:%03d) (%14s)\n", name, id, goid)
	return id
}

func ThreadStop(id int) {
	ThreadMutex.Lock()
	name := ThreadNames[id]
	count := ThreadLoopCount[id]
	goid := ThreadNameToGoids[name]
	ThreadMutex.Unlock()
	fmt.Printf("Thread %33s:%03d (%14s) stopped after %v iterations\n", name, id, goid, count)
}

func ThreadLoopInc(id int) {
	ThreadMutex.Lock()
	ThreadLoopCount[id]++
	ThreadMutex.Unlock()
}
