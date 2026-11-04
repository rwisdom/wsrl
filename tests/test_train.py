"""plan section 8 order steps 1 and 4. Marked slow.

TODO:
- test_budget_run_writes_checkpoint: train.py with --budget 60 on example_setup writes model.zip and
  vecnormalize.pkl; loading them plays 10 episodes without error; wall clock stayed near the budget.
- test_warm_start_not_below_previous: second run from the first's checkpoint; its first evaluation
  mean HPS is not below the first run's final one (within noise).
"""
