// rlsim runs a SIML rotation against the same encounter the DLL builds (plan 3.4 step 2).
//
// Forever's wowsimcli cannot do this: it lacks the Restoration Shaman heals.
//
// TODO:
//   - flags: --setup <json>, --siml <file>, --seed <n>, --iterations <n> (default 100),
//     --duration <s> (overrides the Setup), --attribution (attach lib/attribution and print weights).
//   - register.Once(); encounter.ParseSetup; encounter.Build; siml.Parse (sim/siml/file.go:132) and
//     attach the rotation to the healer player; set RandomSeed and Iterations; core.RunSim
//     (sim/core/sim.go:118); print the RaidSimResult as JSON on stdout.
//   - effective HPS comes from the per-spell health resource metrics (actual_gain, proto/api.proto:342-346),
//     not the raw hps field; python/wsrl/baseline.py does that division.
//
// Used by tests/test_g1_scripted.py and python/wsrl/baseline.py.
package main

func main() {}
