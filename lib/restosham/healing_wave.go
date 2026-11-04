package restosham

// Healing Wave ladder (plan 3.2 step 3). Pattern: sim/paladin/holy_light.go.
//
// TODO: func (a *Agent) registerHealingWave():
//   - spelldata.Ranked(ids...) for the ladder (sim/core/spelldata/ladder.go:19); ids in spell_ids.go.
//   - per rank: spelldata.SpellConfig(&a.Unit, rank, spelldata.Magic(core.ProcMaskSpellHealing)) for
//     coefficient and variance (resolve_spell.go:87, 189); Flags: SpellFlagHelpful | SpellFlagAPL;
//     cast time from rank.CastTime() shortened by Improved Healing Wave; mana cost from the row times
//     Tidal Focus; ApplyEffects: spell.CalcAndDealHealing(sim, target, rank.HealEffect().Average(),
//     spell.OutcomeHealing).
//   - consume one Healing Way stack's multiplier on the target (talents.go) and apply the Earthfury
//     8-piece interaction (item_sets.go).
//
// Test (heals_test.go): average equals the row's tooltip midpoint plus coefficient times +Healing on a
// damaged player.
