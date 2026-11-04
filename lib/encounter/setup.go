// Package encounter turns a Setup into the proto.RaidSimRequest the sim runs (plan 3.3, section 7).
//
// Shared by the DLL (lib/engine) and cmd/rlsim: this is the one place a request is built, so the
// Policy and the SIML baseline always fight the same encounter.
//
// TODO (setup.go):
//   - type Setup struct matching setups/example.json: Items map[slot]int32 (or ordered slice),
//     Talents string, Consumables, Encounter string ("tank" | "tank_friendlies"), DurationSeconds,
//     TankPreset string, SpellIDs []int32 (the action list, copied from restosham.SpellIDs),
//     AuraLabels []string (the A observation slots), Intake {Rate, Variance} for the friendlies.
//   - func ParseSetup(b []byte) (Setup, error) with validation: every slot id resolves in the database,
//     the talents string has the right length, SpellIDs non-empty, duration > 0.
//
// Test (encounter_test.go): setups/example.json parses and validates.
package encounter
