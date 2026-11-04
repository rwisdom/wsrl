package engine

// Observation vector and reward (plan section 7, order step 3).
//
// float32, fixed length, every duration in seconds, every fraction in 0..1. S = len(Setup.SpellIDs),
// A = len(Setup.AuraLabels). Layout, in order:
//
//	0        fight time remaining fraction      sim.GetRemainingDurationPercent()
//	1        mana fraction                       CurrentMana() / MaxMana()
//	2        mana absolute / 10000               CurrentMana()
//	3        five-second-rule time left          PseudoStats.FiveSecondRuleRefreshTime - CurrentTime
//	4        GCD remaining                       GCD.TimeToReady(sim)
//	5        hardcast remaining                  Hardcast.Expires - CurrentTime
//	6..6+S   per-spell cooldown remaining        spell.TimeToReady(sim); 0 for nil spells
//	next     tank health fraction                CurrentHealth() / MaxHealth()
//	next 2   tank damage last 1 s and 3 s / max health   damage.go window
//	next     boss time to next melee swing       the boss unit's auto-attack timer
//	next A   remaining duration of each aura label (stacks where it stacks)   GetAura(label).RemainingDuration, GetStacks()
//	last 3   +Healing/1000, spell crit %/100, haste/100   GetStat(stats.HealingPower), GetStat(stats.SpellCritPercent), haste
//
// TODO:
//   - func ObsLen(S, A int) int.
//   - func (s *Sim) Observe(out []float32).
//   - func (s *Sim) Reward() float64: return gained, set gained = 0. gained is fed by the attribution
//     aura's effective-heal tracking (every heal on the tank adds its effective amount).
//
// Test (observe_test.go): scripted Lesser Healing Wave on a damaged tank gives a reward equal to the
// health gained; the vector length matches ObsLen.
