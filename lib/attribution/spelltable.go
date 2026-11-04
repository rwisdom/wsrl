package attribution

// Spell table for the Agreement Test (plan section 9 order step 4; section 10 step 1).
//
// TODO:
//   - type SpellRow struct { ID int32; Name string; Coeff, Base, CastTime, Mana float64 }.
//   - func SpellTable(unit *core.Unit, ids []int32) []SpellRow, resolved from sim/core/spelldata by
//     spell id (spelldata.SpellConfig, HealEffect, CastTime), i.e. the numbers the sim actually used.
//   - JSON encoding for rlSpellTable.
//
// Test: tests/agreement/test_coefficients.py prints these beside Parsers/Spells_Generated.lua.
