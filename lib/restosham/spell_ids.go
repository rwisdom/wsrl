package restosham

// The canonical action list (plan 3.2, section 7).
//
// Every heal spell id the spec can learn, talented or not, in a fixed order. A Setup copies this list;
// the DLL's action index i casts SpellIDs[i]; an id the current talents did not teach is masked.
// APPEND ONLY: reordering invalidates every checkpoint (plan, "Upgrading forever" step 4).
//
// TODO:
//   - var SpellIDs = []int32{ ... }  highest rank per ladder: Healing Wave, Lesser Healing Wave, Chain Heal,
//     Riptide, Healing Stream Totem, Nature's Swiftness, Mana Tide Totem, Water Shield (refresh).
//     Decide whether lower ranks get their own indices (cheap downranked heals) before training starts.
//   - var SpellNames = map[int32]string for labels.
//   - the ladder id slices used by the registerX functions.
//
// Test: tests/test_dll.py compares rlSpellIds with this list.
