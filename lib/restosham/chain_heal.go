package restosham

// Chain Heal ladder (plan 3.2 step 3).
//
// TODO: func (a *Agent) registerChainHeal():
//   - bounce count and per-bounce falloff from the spelldata rows, not hard-coded.
//   - next target after the primary: the raid player with the most missing health among the
//     remaining ones (sim.Raid players with a health bar), never the same unit twice.
//   - each bounce is its own CalcAndDealHealing so the attribution hooks see one heal per bounce.
//
// Test (heals_test.go): bounce count and falloff against the rows; a lone tank gets one heal.
