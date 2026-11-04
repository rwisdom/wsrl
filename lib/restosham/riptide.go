package restosham

// Riptide (plan 3.2 step 3): a direct heal plus a HoT.
//
// TODO: func (a *Agent) registerRiptide():
//   - direct part as in lesser_healing_wave.go.
//   - HoT through the Dot config on the spell, periodic coefficient from the row's periodic effect
//     (spelldata/resolve_aura.go:93), tick count and period from the row.
//   - the aura label must match the Setup's aura list so observe.go can read RemainingDuration.
//
// Test (heals_test.go): tick count equals the row's duration over period.
