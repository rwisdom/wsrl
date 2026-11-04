"""Gymnasium env and the vector env over DLL handles (plan section 7, order steps 2 and 3).

TODO:
- ForeverEnv(gymnasium.Env): one Handle. observation_space Box(float32, obs_len), action_space
  Discrete(S + 1). reset(seed) -> rlReset + first advance + observe. step(a): act, advance, reward,
  terminated on status 1 or 2. action_masks() for sb3_contrib. Debugging only.
- ForeverVecEnv(stable_baselines3.common.vec_env.VecEnv): N handles, step_wait through rlStepAll,
  action_masks() returning the stacked masks, per-handle seeds from a numpy Generator (evaluation
  passes seeds.EVAL_SEEDS), auto-reset on episode end with the usual VecEnv info keys.
- Reward: rlReward / tank max health (one full bar = 1). Death: terminated and reward minus the
  baseline's full-fight effective healing in the same units (baseline.py provides the number).
  Fight end: terminated, no bonus.
- Training only ever uses ForeverVecEnv.

Tests: tests/test_env.py, tests/test_vec_env.py.
"""
