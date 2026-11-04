"""Baseline effective HPS through cmd/rlsim (plan 3.4 step 3, section 8).

TODO:
- run(setup, siml_path, seeds, duration, rlsim_bin=None) -> per-iteration effective healing:
  subprocess cmd/rlsim with --setup --siml --seed --iterations, parse the RaidSimResult JSON.
- effective healing = sum over the healer's spells of the health resource metrics' actual_gain
  (proto/api.proto:342-346), never the raw hps field; HPS = that / duration.
- mean and 90% confidence interval; death_penalty(setup) = baseline full-fight effective healing in
  health-bar units (used by env.py).
- main(): CLI `wsrl-baseline --setup --siml --seeds --duration`, prints mean, CI, per-seed values.

Tests: tests/test_g1_scripted.py (G1), tests/test_evaluate.py (G2 comparison).
"""


def main() -> None:
    raise NotImplementedError("see module docstring")
