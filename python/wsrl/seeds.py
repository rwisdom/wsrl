"""The fixed evaluation seed list (plan sections 7 and 8).

Every Policy and every baseline run plays the same 100 fights, so comparisons are paired.
Never change this list once a checkpoint exists; add a second list instead.
"""

# TODO: 100 integers drawn once, written out literally here (no runtime generation).
EVAL_SEEDS: list[int] = []
