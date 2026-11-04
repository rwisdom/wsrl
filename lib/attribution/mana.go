package attribution

// Mana-stat values, computed at read time (plan section 9).
//
// TODO:
//   - MP5: duration / 5 * HPM.
//   - Intellect mana pool: 15 * intMult * HPM (mana.go:63-68; intMult from the Int talents).
//   - Spirit: regenPerSpirit * spiritMult * effectiveSeconds * HPM, with regenPerSpirit = 0.125
//     (the sim's 6.25 + Spirit/8 for every class, mana.go:75-76; no knee at 50 Spirit) and
//     effectiveSeconds = timeOutsideFSR + castPct * timeInFSR. Track the FSR split from
//     PseudoStats.FiveSecondRuleRefreshTime at every event.
//   - Crit extras: Water Shield mana per crit valued at HPM, as Parsers/RestoShaman.lua does.
//   - func (a *Attribution) Weights(out []float64): the six slots, +Healing first.
//
// Test (attribution_test.go): one filler cast with known HPM gives MP5 = duration / 5 * HPM.
