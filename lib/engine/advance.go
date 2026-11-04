package engine

// The Decision Point loop (plan 3.1, section 6).
//
// The sim's own NeedsInput flag is never read: it fires mid-hardcast once the GCD ends, which is noise.
//
// TODO: func (s *Sim) Advance() Status:
//   - loop sim.Step() until one of:
//       Decision Point: healer.GCD.IsReady(sim) && healer.Hardcast.Expires <= sim.CurrentTime &&
//                       !healer.IsChanneling()   (sim/core/unit.go:177, 185; gcd.go:173)  -> Running
//       fight end                                                                    -> FightOver
//       tank.CurrentHealth() <= 0                                                    -> TankDead
//   - the test above is evaluated after every Step, so the event that completes a hardcast is a
//     Decision Point and the Policy sees the cast's outcome in the observation.
//
// Test (advance_test.go): a scripted Healing Wave yields no Decision Point during its cast time and
// one on the event that completes it; an idle healer yields one per event.
