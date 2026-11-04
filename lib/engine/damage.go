package engine

// Tank damage window for the observation (plan section 5).
//
// TODO:
//   - type damageWindow: ring buffer of (time.Duration, float64) with Push and Sum(since).
//   - attachDamageAura(tank): an aura registered on the tank (unit.RegisterAura, sim/core/aura.go:542)
//     with OnSpellHitTaken and OnPeriodicDamageTaken pushing (sim.CurrentTime, result.Damage).
//     Never keep the SpellResult: it is pooled and disposed after the hooks (spell_result.go:827).
//   - observe.go reads Sum(1 s) and Sum(3 s) as fractions of the tank's max health.
//   - cleared on Reset.
