// Package restosham is wsrl's Restoration Shaman agent, built on forever's public shaman API (plan 3.2).
//
// Forever's own Resto Shaman has no compiled heals (sim/shaman/_heals.go is excluded from the build);
// this package registers its own factory for proto.Spec_SpecRestorationShaman. Private shaman internals
// are not reachable: every number comes from sim/core/spelldata rows.
//
// TODO (agent.go):
//   - type Agent struct{ *shaman.Shaman }
//   - Register(): core.RegisterAgentFactory(proto.Player_RestorationShaman{},
//     proto.Spec_SpecRestorationShaman, factory, setter), as sim/shaman/restoration/restoration.go:9-24 does.
//   - factory: shaman.NewShaman(character, talentsString, selfBuffs) (sim/shaman/shaman.go:20); return &Agent{}.
//   - (a *Agent) Initialize(): call the embedded Initialize, then registerHealingWave, registerLesserHealingWave,
//     registerChainHeal, registerRiptide, registerHealingStreamTotem, applyTalents, registerItemSets.
//   - (a *Agent) GetShaman() *shaman.Shaman.
//   - OnGCDReady: no-op. The Policy acts through the DLL; the agent never runs a rotation.
//
// Tests: heals_test.go, talents_test.go.
package restosham

// Register adds wsrl's Restoration Shaman factory to the sim. Called once from lib/register.
func Register() {
	panic("restosham.Register: not implemented yet (plan 3.2 step 1)")
}
