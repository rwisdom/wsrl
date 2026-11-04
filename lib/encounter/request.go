package encounter

// Setup to RaidSimRequest (plan 3.3 steps 1, 2, 4; section 7).
//
// TODO: func Build(s Setup) (*proto.RaidSimRequest, error):
//   - healer: a proto.Player for the Restoration Shaman from the Setup's items and talents, the way
//     forever's tests build one (core.GetGearSet style, sim/core/test_utils.go:268; raid layout as
//     core.SinglePlayerRaidProto, sim/core/raid.go:446). Party 0, player 0 is the healer; the engine
//     relies on that position.
//   - tank: a Protection Warrior from TankPreset, listed in Raid.tanks so the boss melees it, with
//     HealingModel{Hps: 0} to turn the synthetic self-heal off.
//   - friendlies (Encounter == "tank_friendlies"): two more players with an empty rotation, not listed
//     as tanks, so the boss ignores them. Every raid player has a health bar (raid.go:342); target
//     dummies do not, so friendlies are players.
//   - boss: the default raid boss from sim/encounters with Patchwerk-style melee.
//   - SimOptions: Interactive = true (stops the APL; NeedsInput is never read), Duration from the Setup,
//     RandomSeed set by the caller.
//
// Test (encounter_test.go): core.NewSim on the result succeeds; party layout as expected.
