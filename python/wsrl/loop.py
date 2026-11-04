"""The outer loop (plan section 11).

wsrl-loop --pool pools/resto.json --duration 180s --budget 120s --runs 10 [--gear ...] [--talents ...]
          [--checkpoint ...] [--talent-tune ...]

TODO, per cycle:
1. Start: first Setup from --gear / --talents where given, random from the pool for the rest; validate
   by building a sim. Never stubbed.
2. Train for the budget (train.py); warm start from the last checkpoint after the first cycle.
3. Evaluate 100 episodes on seeds.EVAL_SEEDS with attribution (evaluate.py); statweights.summarize.
4. Pick gear (gear.py).            May be stubbed with a hand-picked Setup in the pilot.
5. Pick talents (talents.py).      May be stubbed in the pilot.
6. Write the next Setup, close handles, append {setup hash, weights, HPS, death rate, checkpoint path}
   to runs/log.jsonl.
7. Stop after --runs cycles or when neither gear nor talents change.

Test: tests/test_loop.py (two stubbed cycles complete and log two lines).
"""


def main() -> None:
    raise NotImplementedError("see module docstring")
