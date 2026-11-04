"""Talent strings: validation, random draw, and the two-kind pick (plan section 11 steps 1 and 5).

TODO:
- parse(s) / format(tree ranks) for the forever talent string; validate(s): tier and prerequisite
  rules, 51 points max; random_valid(rng).
- stat talents priced by the weights: Tidal Mastery as crit, Purification as a share of effective
  healing, Tidal Focus through filler healing-per-mana; a point goes where value per point is highest.
- structural talents (Riptide, Water Shield, Healing Way, Nature's Swiftness, Mana Tide Totem,
  Improved Healing Wave): candidates = current string with one toggled; each fine-tuned from the
  current checkpoint for --talent-tune seconds (default budget / 10) then judged on 20 episodes; the
  best goes forward with its fine-tuned checkpoint. Works only because the action list is fixed.
- every produced string is validated by building a sim (setups.validate).

Test: tests/test_talents.py.
"""
