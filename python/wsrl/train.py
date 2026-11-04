"""One Training Run: MaskablePPO on ForeverVecEnv with a wall-clock budget (plan section 8).

TODO:
- args: --setup, --budget (seconds, 30 to 300), --envs (8), --checkpoint (warm start path),
  --from-scratch, --out (runs/<setup hash>/<timestamp>/).
- model: sb3_contrib.MaskablePPO("MlpPolicy", net_arch=[128, 128]); 256 steps per env per rollout,
  batch 512, lr 3e-4, gamma 0.99, gae_lambda 0.95, ent_coef 0.01. VecNormalize with norm_reward=False.
- warm start (default): MaskablePPO.load(path, env=env) and load the VecNormalize stats; reset the
  learning-rate schedule so a short fine-tune does not inherit a decayed rate. --from-scratch builds new.
- callbacks.BudgetCallback and callbacks.EvalCallback; TensorBoard under the out dir.
- save model.zip AND vecnormalize.pkl every time; print the out dir.
- main(): CLI `wsrl-train`.

Test: tests/test_train.py (60 s run writes a checkpoint that plays 10 episodes; warm start does not
start below the previous run's HPS).
"""


def main() -> None:
    raise NotImplementedError("see module docstring")
