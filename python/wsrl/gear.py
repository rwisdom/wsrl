"""Gear pick from stat weights (plan section 11 step 4).

TODO:
- item_stats(dll, ids) -> six stats per item through rlItemStats (core.GetItemByID on the Go side).
- score(item_stats, weights) = sum of weight * stat over the six stats; items scored independently,
  set bonuses ignored in the pilot.
- pick(pool, weights, dll) -> item id per slot: best candidate per slot from pools/resto.json honoring
  unique and slot rules (two rings, two trinkets, off-hand vs two-hander).
- Later: hand the weights to forever's optimizer instead.

Test: tests/test_gear.py with a tiny synthetic pool.
"""
