"""Predefined cast sequence from CSV, exposed as a policy callable (plan section 7, option).

CSV columns: spell (name or id, or WAIT), optional at_s, optional repeat.

TODO:
- load(path, spell_ids) -> list of rows resolved to action indices.
- ordered mode: consume one row per Decision Point. timed mode (at_s present): cast each row at its
  time, Wait until then. Needs the fight time from the observation (slot 0) and the duration.
- on exhaustion: "wait" | "loop" | handover to a policy callable (policies.from_checkpoint).
- Sequence(...).as_policy() -> (obs, mask) -> int; masked rows are skipped or waited on (decide and
  document).

Used by evaluate.py --sequence, the Agreement Test, and an optional behavior-cloning warm start.
Test: tests/test_sequence.py.
"""
