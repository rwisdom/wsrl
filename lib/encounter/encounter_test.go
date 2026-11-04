package encounter

// TODO (plan 3.3 step 5, section 7):
//   - TestParseExampleSetup: ../../setups/example.json parses and validates.
//   - TestBuildRequest: Build on the example gives a request core.NewSim accepts; party 0 player 0 is
//     the healer; the tank is in Raid.tanks; "tank_friendlies" adds two non-tank players.
//   - TestIntake: after Arm, a friendly's health falls at the configured rate over 10 s; a heal is
//     clamped at max health.
//
// Needs register.Once() and the with_db tag.
