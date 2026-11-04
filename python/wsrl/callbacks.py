"""stable-baselines3 callbacks for train.py (plan section 8).

TODO:
- BudgetCallback(BaseCallback): _on_step returns False once the wall clock passes --budget seconds;
  learn() then returns and train.py writes the checkpoint.
- EvalCallback(BaseCallback): every N rollouts, 20 deterministic episodes on seeds.EVAL_SEEDS[:20],
  logs mean effective HPS, overheal fraction, death rate, mana left, action histogram, Wait share,
  entropy to TensorBoard; keeps the last 5 means and logs their slope as the plateau signal.
  Written, never acted on.

Test: tests/test_train.py.
"""
