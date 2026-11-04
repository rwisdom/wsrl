"""Evaluation episodes with attribution (plan section 8 order step 3, section 11 step 3).

TODO:
- evaluate(policy, setup, seeds=seeds.EVAL_SEEDS, attribution=True) -> list of per-episode records:
  six raw attribution sums (rlAttribution), effective HPS, death flag, overheal fraction, events CSV
  when asked (for the Agreement Test).
- deterministic Policy from policies.from_checkpoint; --sequence <csv> uses sequence.py instead.
- comparison report against baseline.run: means, 90% CIs, "beats baseline" when the intervals do not
  overlap (gate G2).
- main(): CLI `wsrl-evaluate --setup --checkpoint | --sequence --seeds --report`.

Test: tests/test_evaluate.py.
"""


def main() -> None:
    raise NotImplementedError("see module docstring")
