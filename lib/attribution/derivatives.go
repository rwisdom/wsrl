package attribution

// Per-heal derivatives (plan section 9): the addon's _Healing, critPer1Pct, _Haste, _Intellect.
//
//	dSP    = coeff * M
//	dCrit% = heal * CB / ((1 + C * CB) * 100)
//	dHaste = heal / ((1 + H) * HasteConv)        only when the spell has a cast time
//
// TODO:
//   - coeff: spell.BonusCoefficient for a direct heal, dot.BonusCoefficient for a tick
//     (spell.go:188-197, spelldata/resolve_aura.go:93).
//   - M: spell.CasterHealingMultiplier() (spell_result.go:998) times the target's HealingTakenMultiplier.
//   - C: spell.HealingCritChance(target) (spell_result.go:323-330); CB: crit multiplier minus 1.
//   - crit rating per 1%: divide by core.CritRatingPerCritChance.
//   - Intellect: per 1% crit divided by the Shaman Int-to-crit rate (CritPerIntMaxLevel,
//     base_stats_auto_gen.go:38-49), plus Mental Quickness 0.15 * rank * dSP.
//   - H: the healer's spell haste fraction; HasteConv from the sim's haste rating conversion.
//   - chain-cast counting for the segment-level haste cap: copy the rule from HSW-Forever
//     Classes/CastTracker.lua (StartCast:25, FinishCast:40, IncChainCasts:20).
//
// Test (attribution_test.go): one non-crit Healing Wave with no procs matches the closed forms.
