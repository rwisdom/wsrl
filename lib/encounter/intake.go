package encounter

// Damage Intake on the friendlies (plan 3.3 step 3).
//
// TODO:
//   - type Intake struct { Rate, Variance float64 } (health per second, fraction).
//   - func (in Intake) Arm(sim *core.Simulation, units []*core.Unit): core.StartPeriodicAction
//     (sim/core/periodic_action.go:96) whose action calls unit.RemoveHealth(sim, amount)
//     (sim/core/health.go:73) with the rate and variance.
//   - Reset truncates pending actions, so the engine calls Arm after every reset.
//
// Test (encounter_test.go): a friendly loses health at the configured rate; a heal on it is clamped
// at max health.
