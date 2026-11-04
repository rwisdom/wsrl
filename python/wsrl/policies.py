"""Policy callables sharing one signature: (obs: np.ndarray, mask: np.ndarray) -> int (plan section 7 step 4).

Anything that takes a policy (evaluate.py, sequence.py handover, the Agreement Test) accepts one of these.

TODO:
- random_masked(rng): uniform over actions with mask == 1.
- lhw_spam(spell_ids): the Lesser Healing Wave index when castable, else Wait. The DLL half of gate G1.
- from_checkpoint(path, vecnormalize_stats): wraps a MaskablePPO .predict(deterministic=True),
  applying the saved observation normalization.

Test: tests/test_g1_scripted.py.
"""
