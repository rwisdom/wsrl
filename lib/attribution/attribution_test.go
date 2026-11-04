package attribution

// TODO (plan section 9 order steps 1-4):
//   - TestEffectiveHealingEqualsReward: over an episode of scripted casts, the sum of effective heals
//     equals the engine's reward total.
//   - TestClosedForms: a single non-crit Healing Wave, no procs: dSP/heal == coeff*M/heal; dCrit equals
//     the closed form with the sim's C and CB.
//   - TestMP5: one filler cast with known HPM gives duration/5*HPM.
//   - TestEventLogRows: row count equals casts + heals + ticks.
