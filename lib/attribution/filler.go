package attribution

// Filler accounting (plan section 9). Every heal that costs mana is a Filler.
//
// TODO:
//   - at OnCastComplete with spell.CurCast.Cost > 0: fillerCasts++, fillerMana += cost.
//   - at each heal: fillerHealing += effective.
//   - returned mana (Water Shield crit energize, Earthfury 5-piece) from the actual energize the sim
//     applies (resource metrics), not the addon's analytic 2% estimate; comes off fillerMana.
//   - func (a *Attribution) HPM() float64 = fillerHealing / (fillerMana - returned); 0 when no filler mana.
