package main

// Handle table for the DLL (plan section 5).
//
// TODO:
//   - var handles = map[int32]*engine.Sim{}, guarded by a sync.Mutex; a monotonically increasing next id.
//   - get(h) (*engine.Sim, bool); put(*engine.Sim) int32; del(h).
//   - stepAll(hs []int32, out []int32): one goroutine per handle calling sim.Advance(), a sync.WaitGroup,
//     each status written to out[i]. Each sim is touched by exactly one goroutine; the map is read under
//     the mutex before the goroutines start. forever's own async runner already runs sims concurrently
//     in one process, so sim/core tolerates this.
//
// Test: tests/test_vec_env.py (8 handles in parallel equal 8 sequential runs).
