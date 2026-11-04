// Package main is the cgo surface of the wsrl DLL (plan section 5).
//
// Only this file and handles.go import "C". No sim logic lives here: each export converts its C
// arguments, looks up the handle, and calls one method on engine.Sim. Build with
// `make dll` (go build -buildmode=c-shared --tags=with_db ./lib).
//
// TODO: one //export per row of the plan's export table:
//
//	rlNew(json *C.char) int32            engine.New(setupJSON); register handle; -1 on error
//	rlReset(h int32, seed int64)         sim.Reset(seed)
//	rlAdvance(h int32) int32             sim.Advance(): 0 running, 1 fight over, 2 tank dead
//	rlAct(h int32, i int32) bool         sim.Act(i); i == len(spells) is Wait
//	rlStepAll(hs *int32, n int32, out *int32)   handles.go fan-out
//	rlObserve(h int32, out *float32, n int32)   sim.Observe
//	rlMask(h int32, out *uint8, n int32)        sim.Mask; Wait always 1
//	rlReward(h int32) float64            sim.Reward(): returns and zeroes gained
//	rlAttribution(h int32, out *float64) six slots, +Healing first (section 9)
//	rlEvents(h int32) *C.char            CSV; caller frees with rlFree
//	rlSpellIds(h int32, out *int32, n int32)
//	rlSpellTable(h int32) *C.char        JSON per spell id: coeff, base, cast time, mana (section 10)
//	rlItemStats(ids *int32, n int32, out *float64)   six stats per item from core.GetItemByID (section 11)
//	rlClose(h int32)                     sim.Close(); delete handle
//
// Rules: never run two exports on one handle concurrently; never cache a SpellResult.
package main

/*
#include <stdlib.h>
*/
import "C"

import "unsafe"

// rlFree releases a string returned by rlEvents or rlSpellTable. Safe on NULL.
//
//export rlFree
func rlFree(p *C.char) {
	C.free(unsafe.Pointer(p))
}

func main() {}
