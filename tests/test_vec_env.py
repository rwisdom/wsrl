"""plan section 7 order step 3; section 5 order step 4.

TODO:
- test_vec_matches_single: ForeverVecEnv over 8 handles with seeds s1..s8 gives the same per-episode
  returns as 8 ForeverEnv runs with the same seeds and the same scripted policy.
- test_step_all_parallel_equals_sequential: rlStepAll statuses and observations equal sequential
  rlAdvance calls on the same handles.
"""
