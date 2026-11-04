"""Stat weights from per-episode attribution (plan section 9, section 11 step 3).

TODO:
- normalize(rows): divide every slot by the +Healing slot per episode so +Healing reads 1.00.
- summarize(rows) -> mean and stdev per stat (+Healing, Crit, Haste, Intellect, Spirit, MP5),
  mean effective HPS, death rate.
- report(...) text/markdown for the loop log and the console.

Test: tests/test_statweights.py on synthetic rows (no DLL needed).
"""
