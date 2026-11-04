package restosham

// Restoration talents (plan 3.2 step 4). Ranks come from a.Talents (public proto on shaman.Shaman).
//
// TODO: func (a *Agent) applyTalents():
//   - Tidal Mastery: spell crit for healing spells.
//   - Purification: healing multiplier.
//   - Tidal Focus: mana cost multiplier on the heals.
//   - Healing Way: stacking aura on the target, multiplies Healing Wave; registered per target.
//   - Improved Healing Wave: shorter Healing Wave cast.
//   - Nature's Swiftness, Mana Tide Totem: spells (also in spell_ids.go so the Policy can cast them).
//   - Water Shield: already in sim/shaman/shields.go, registered by the embedded shaman; nothing here.
//   - Mental Quickness: only after confirming the row exists in forever's spelldata.
//
// Test (talents_test.go): each talent changes exactly the number it should.
