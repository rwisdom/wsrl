// Package attribution implements the HSW addon's per-heal stat attribution in Go (plan section 9).
//
// One aura on the healer plus hooks on the tank. Six accumulators in the addon's slot order
// (+Healing, Crit, Haste, Intellect, Spirit, MP5); Python divides every slot by +Healing so it reads 1.00.
// The Agreement Test (section 10) proves the method matches the addon's Lua.
//
// TODO (aura.go):
//   - type Attribution struct: the six accumulators, knownHealth float64, the filler fields (filler.go),
//     FSR time split, a pointer to the engine's gained and event log (or callbacks for both).
//   - func Attach(healer *core.Character, tank *core.Unit) *Attribution: unit.RegisterAura
//     (sim/core/aura.go:542) with OnHealDealt, OnPeriodicHealDealt, OnCastComplete on the healer and
//     OnSpellHitTaken, OnPeriodicDamageTaken on the tank.
//   - effective heal per event: at a damage or heal hook on the tank,
//     effective = tank.CurrentHealth() - knownHealth; knownHealth = CurrentHealth().
//     Every change to the tank's health fires one of those hooks, so the difference is this heal's
//     effective amount (the heal hooks fire after GainHealth, spell_result.go:807-825, so the clamped
//     amount is not on the SpellResult).
//   - un-crit: if result.DidCrit(), heal = effective / the spell's crit multiplier (spell_outcome.go:147-160).
//   - func (a *Attribution) Reset().
//
// Test (attribution_test.go): summed effective healing equals the episode's reward total.
package attribution
