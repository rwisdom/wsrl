# Glossary

The capitalized forms are the ones docs/RL_Plan.md uses. Project terms first, then the sim-side terms
that come from WoWSims Forever.

## Project terms

- **Policy**: the trained neural network that picks one action at every Decision Point. Its action
  space is the Setup's fixed spell-id list plus Wait; its target is always the Heal Target. A Policy
  is stored as a checkpoint: the MaskablePPO model plus the VecNormalize statistics.
- **Setup**: everything that defines one training problem except the Policy: item ids per slot, the
  talents string, consumables, the encounter name, fight duration, the tank preset, the spell-id list
  and the aura labels. Python passes a Setup JSON to the DLL; `lib/encounter` expands it into the
  sim's `RaidSimRequest`. One Setup is one draw or one pick from the pool plus the duration.
- **Training Run**: one PPO run on one Setup under a wall-clock budget (30 s to 5 min), warm-started
  from the previous checkpoint unless started from scratch. The outer loop does one per cycle.
- **Stat Weight**: how much effective healing one point of a stat is worth, relative to +Healing,
  which is fixed at 1.00. Six stats: +Healing, Crit, Haste, Intellect, Spirit, MP5.
- **Attribution**: the method that produces Stat Weights from how the Policy actually plays: at every
  heal, the derivative of that heal with respect to each stat is accumulated. It is the HSW addon's
  method, implemented in Go in `lib/attribution`; ADR 0001 says why it is not perturbation.
- **Filler**: a heal that costs mana. Every Restoration Shaman heal is a Filler (none is a raid
  cooldown), so filler healing per mana is the exchange rate that prices the mana stats.
- **Agreement Test**: the pytest in `tests/agreement/` that replays one episode's event log through
  the addon's own Lua and asserts that its six weights match the DLL's. It is what makes "same
  method as the addon" a tested claim rather than an intention.

## Sim-side terms

- **Sim Agent**: forever's term for the Go object that plays one character: its spells, talents,
  auras and rotation. wsrl registers its own Restoration Shaman agent (`lib/restosham`) because
  forever's has no compiled heals.
- **Decision Point**: a moment at which the Policy is asked for an action: after a sim event, the
  healer's GCD is ready, no hardcast is in progress and nothing is channeling. The event that
  completes a hardcast is one. The sim's own `NeedsInput` flag is not used.
- **Interactive Mode**: `SimOptions.Interactive = true`. Here it does one job: the healer's rotation
  returns without running an APL, so the DLL can drive the healer from outside. Pets and the boss
  keep running their own rotations.
- **Heal Target**: the unit every Policy action is cast on: the Protection Warrior tank. Chain Heal
  bounces beyond it by the spell's own rules.
- **Damage Intake**: periodic `RemoveHealth` on the friendlies in the second encounter, re-armed
  after every reset, so Chain Heal has bounce targets that are actually hurt.
- **Effective Healing**: the health a heal actually added after clamping at max health. The reward
  is Effective Healing per step; the heal hooks fire after the clamp, so it is measured as the
  tank's health difference across the hook.
- **Overheal**: the part of a heal that was lost to the max-health clamp. Raw healing minus
  Effective Healing; logged as a fraction, never rewarded.
