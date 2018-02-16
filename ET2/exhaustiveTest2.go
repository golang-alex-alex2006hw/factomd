package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/FactomProject/electiontesting/controller"
	"github.com/FactomProject/electiontesting/election"
	"github.com/FactomProject/electiontesting/imessage"

	"reflect"
)

var _ = reflect.DeepEqual

//================ main =================
func main() {
	recurse(3, 3, 2000)
}

// newElections will return an array of elections (1 per leader) and an array
// of volunteers messages to kick things off.
//		Params:
//			feds int   Number of Federated Nodes
//			auds int   Number of Volunteers
//			noDisplay  Passing a true here will reduce memory consumption, as it is a debugging tool
//
//		Returns:
//			controller *Controller  This can used for debugging (Printing votes)
//			elections []*election   Nodes you can execute on (returns msg, statchange)
//			volmsgs   []*VoluntMsg	Volunteer msgs you can start things with
func newElections(feds, auds int, noDisplay bool) (*controller.Controller, []*election.Election, []*mymsg) {
	con := controller.NewController(feds, auds)

	if noDisplay {
		for _, e := range con.Elections {
			e.Display = nil
		}
		con.GlobalDisplay = nil
	}
	var msgs []*mymsg
	fmt.Println("Starting")
	for _, v := range con.Volunteers {
		for i, _ := range con.Elections {
			my := new(mymsg)
			my.leaderIdx = i
			my.msg = v
			msgs = append(msgs, my)
			fmt.Println(my.msg.String(), my.leaderIdx)
		}
	}

	global := con.Elections[0].Display.Global
	for i, ldr := range con.Elections {
		con.Elections[i] = CloneElection(ldr)
		con.Elections[i].Display.Global = global
	}

	return con, con.Elections, msgs
}

type mymsg struct {
	leaderIdx int
	msg       imessage.IMessage
}

var solutions = 0
var breadth = 0
var loops = 0
var depths []int
var cuts []int
var hitlimit int

var globalRunNumber = 0

var extraPrints = true
var insanePrints = false

func dive(msgs []*mymsg, leaders []*election.Election, depth int, limit int) {
	incDepths(depth)
	depth++
	if depth > limit {
		if extraPrints {
			fmt.Print("Loop/solution/limit at ", depth)
			for _, v := range cuts {
				fmt.Print(v, " ")
			}
			fmt.Println()
		}
		breadth++
		hitlimit++
		return
	}

	if LoopingDetected(leaders[0].Display.Global) == len(leaders) {
		// TODO: Paul you can move this check wherever you need
		incCuts(depth)
		loops++
		return
	}

	fmt.Println("=============== ",
		" Depth=", depth,
		", Hit the Limits=", hitlimit,
		" Breadth=", breadth,
		", solutions so far =", solutions,
		", global count= ", globalRunNumber,
		", loops detected=", loops)
	fmt.Print("Loop/solution/limit at ", depth)
	for _, v := range cuts {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Print("Working at depth: ", depth)
	for _, v := range depths {
		fmt.Print(v, " ")
	}
	fmt.Println()

	if extraPrints {
		// Lots of printing... Not necessary....
		fmt.Println(leaders[0].Display.Global.String())

		for _, ldr := range leaders {
			fmt.Println(ldr.Display.String())
		}

		if insanePrints {
			// Example of a run that has a werid msg state
			if globalRunNumber > -1 {
				fmt.Println("Leader 0")
				fmt.Println(leaders[0].PrintMessages())
				fmt.Println("Leader 1")
				fmt.Println(leaders[1].PrintMessages())
				fmt.Println("Leader 2")
				fmt.Println(leaders[2].PrintMessages())
			}
		}
	}

	done := 0
	for _, ldr := range leaders {
		if ldr.Committed {
			done++
		}
	}
	if done == len(leaders)/2+1 {
		incCuts(depth)
		if extraPrints {
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>> Solution Found @ ", depth)
		}
		breadth++
		solutions++
		return
	}

	for d, v := range msgs {
		var msgs2 []*mymsg
		msgs2 = append(msgs2, msgs[0:d]...)
		msgs2 = append(msgs2, msgs[d+1:]...)
		ml2 := len(msgs2)
		globalRunNumber++

		cl := CloneElection(leaders[v.leaderIdx])

		//if !spewSame(cl, leaders[v.leaderIdx]) {
		//	fmt.Println("Clone Failed")
		//	debugClone(cl, leaders[v.leaderIdx])
		//	os.Exit(0)
		//}

		msg, changed := leaders[v.leaderIdx].Execute(v.msg)

		if changed {
			if msg != nil {
				for i, _ := range leaders {
					if i != v.leaderIdx {
						my := new(mymsg)
						my.leaderIdx = i
						my.msg = msg
						msgs2 = append(msgs2, my)
					}
				}
			}
			gl := leaders[v.leaderIdx].Display.Global
			for _, ldr := range leaders {
				ldr.Display.Global = gl
			}
			// Recursive Dive
			dive(msgs2, leaders, depth, limit)
			for _, ldr := range leaders {
				ldr.Display.Global = cl.Display.Global
			}
			msgs2 = msgs2[:ml2]
		} else {
			incCuts(depth)
		}
		leaders[v.leaderIdx] = cl
	}
}

func incCuts(depth int) {
	for len(cuts) <= depth {
		cuts = append(cuts, 0)
	}
	cuts[depth]++
}

func incDepths(depth int) {
	for len(depths) <= depth {
		depths = append(depths, 0)
	}
	depths[depth]++
}

func recurse(auds int, feds int, limit int) {

	_, leaders, msgs := newElections(feds, auds, false)

	dive(msgs, leaders, 0, limit)
}

// reuse encoder/decoder so we don't recompile the struct definition
var enc *gob.Encoder
var dec *gob.Decoder

// LoopingDetected will the number of looping leaders
func LoopingDetected(global *election.Display) int {
	return global.DetectLoops()
}

// UniqueLeaders returns true if all leaders are unique and copied correctly
func UniqueLeaders(leaders []*election.Election) bool {
	for i := 0; i < len(leaders); i++ {
		for j := i + 1; j < len(leaders); j++ {
			if !leaders[i].IsDifferent(leaders[j]) {
				return false
			}
		}
	}
	return true
}

func init() {
	buff := new(bytes.Buffer)
	enc = gob.NewEncoder(buff)
	dec = gob.NewDecoder(buff)
}

func CloneElection(src *election.Election) *election.Election {
	return src.Copy()
	dst := new(election.Election)
	err := enc.Encode(src)
	if err != nil {
		panic(err)
	}
	err = dec.Decode(dst)
	if err != nil {
		panic(err)
	}
	return dst
}
