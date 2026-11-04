package engine

// TODO (plan section 5 order step 3; section 7):
//   - TestObsLen: Observe fills exactly ObsLen(S, A) values and nothing is NaN.
//   - TestRewardEqualsHealthGained: damage the tank, cast Lesser Healing Wave, Advance; Reward equals
//     the tank's health difference; a second Reward returns 0.
//   - TestDamageWindow: boss swings land in the 1 s and 3 s sums as fractions of max health.
