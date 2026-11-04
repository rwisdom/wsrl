// Package engine owns one simulation behind one DLL handle (plan 3.1, section 5).
//
// Nothing here knows about C. lib/library.go calls these methods; cmd/rlsim does not use this package.
//
// TODO (sim.go):
//   - type Status int32: Running = 0, FightOver = 1, TankDead = 2.
//   - type Sim struct (the plan's rlSim):
//     sim    *core.Simulation
//     healer *core.Character   // Raid.Parties[0].Players[0]
//     tank   *core.Unit        // the Heal Target
//     spells []*core.Spell     // one per id in Setup.SpellIDs; nil when the talents did not teach it
//     attr   *attribution.Attribution
//     damage *damageWindow     // damage.go
//     gained float64           // effective healing since the last Reward()
//     setup  encounter.Setup
//     seed   int64
//   - func New(setupJSON []byte) (*Sim, error): encounter.ParseSetup, register.Once, encounter.Build,
//     core.NewSim, resolve spells by id from the healer's Spellbook, attr.Attach(healer, tank),
//     attachDamageAura(tank).
//   - func (s *Sim) Reset(seed int64): Cleanup if the last episode did not finish, Reseed(seed), Reset(),
//     PrePull(), re-arm intake, attr.Reset(), clear damage window and gained. Same path run() takes
//     between iterations (sim/core/sim.go:359-398).
//   - func (s *Sim) Close(): Cleanup.
//
// Test (advance_test.go): two Sims reset with the same seed and advanced with no casts produce
// identical tank health traces.
package engine
