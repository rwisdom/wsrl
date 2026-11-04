package engine

// Actions and masks (plan section 5, order step 2).
//
// TODO:
//   - func (s *Sim) Act(i int32) bool:
//       i == len(spells): Wait. One sim.Step(), then the Decision Point test again (Advance handles it).
//       spells[i] == nil (talents did not teach it): return false, change nothing.
//       otherwise spell.CanCast(sim, tank) then spell.Cast(sim, tank); return the cast result.
//       Never call Rotation.DoNextAction on the healer.
//   - func (s *Sim) Mask(out []uint8): CanCast(sim, tank) per spell as 0/1, nil spells 0, Wait always 1.
//   - func (s *Sim) SpellIds(out []int32): the Setup's list, so Python can label actions.
//
// Test (act_test.go): casting an uncastable index returns false and changes nothing (mana, GCD, health).
