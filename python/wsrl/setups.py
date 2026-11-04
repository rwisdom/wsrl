"""The Setup: what Python hands the DLL instead of a RaidSimRequest (plan section 7, section 11 step 1).

Mirrors lib/encounter/setup.go and setups/example.json.

TODO:
- @dataclass Setup: items (dict slot -> item id), talents (str), consumables, encounter, duration_seconds,
  tank_preset, intake, spell_ids (list[int]), aura_labels (list[str]).
- from_json / to_json; hash() -> short stable digest used for runs/<setup>/ directories and log.jsonl.
- random_from_pool(pool, rng, duration, gear=None, talents=None): one item per slot from pools/resto.json
  (slot and unique rules), a random valid talent string (talents.py), user-given parts kept as is.
- validate(setup, dll): build one handle with it and close it; raises on failure.

Test: tests/test_pool.py.
"""
