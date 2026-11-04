package engine

// TODO (plan 3.1 step 3; section 5 order step 1):
//   - TestSameSeedSameTrace: two Sims from setups/example.json, Reset(42) each, Advance with no casts
//     until FightOver, tank health sampled at every Decision Point; traces identical.
//   - TestHardcastDecisionPoints: cast Healing Wave at the first Decision Point; no Decision Point
//     until the event that completes the cast; then one.
//   - TestIdleHealerOnePerEvent: with only Wait, every Advance returns after exactly one event.
//   - TestTankDeath: an Intake rate high enough to kill the tank returns TankDead.
