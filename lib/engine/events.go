package engine

// Episode event log for the Agreement Test (plan section 9, order step 4; section 10).
//
// One row per cast and per heal, in the addon's column order (HSW-Forever Classes/StatParser.lua:202,
// addon.EventsHeader) plus the six stat columns the DLL appends:
//
//	time,event,spell_id,spell,target,amount,crit,mana,sp,crit_pct,haste,int,spirit,maxmana
//
// event is one of cast, heal, tick; amount is the effective heal before un-critting.
//
// TODO:
//   - type eventLog with Append(row) and String(); the attribution aura appends from its hooks.
//   - func (s *Sim) Events() string; cleared on Reset.
//
// Test (attribution_test.go): row count equals casts plus heals plus ticks.
