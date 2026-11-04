package restosham

// Heal tests (plan 3.2 step 6).
//
// TODO: build a sim with one Restoration Shaman and one damaged raid player (core.NewSim on a request
// from lib/encounter, or core.RunTestSuite with a .results pin in testdata/):
//   - each heal's average equals the row's tooltip midpoint plus coefficient times +Healing;
//   - Riptide tick count; Chain Heal bounce count and falloff; Healing Stream tick count;
//   - Earthfury 5-piece mana return and 8-piece Healing Wave change.
//
// Needs the with_db tag: go test --tags=with_db ./lib/...
