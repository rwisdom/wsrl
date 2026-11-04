"""Builds pools/resto.json from forever's database (plan section 11).

Never curated by hand; regenerated after every forever bump.

TODO:
- through the DLL (rlItemStats / a database export; add an rlItems export if listing needs it):
  every item a Restoration Shaman can equip, grouped by slot, with its six stats and unique flag.
- the pool file also holds: encounter name, spell_ids (copied from lib/restosham/spell_ids.go via
  rlSpellIds), aura_labels. Fight duration is NOT in the pool; --duration sets it per run.
- main(): CLI `wsrl-pool --out pools/resto.json`.

Test: tests/test_pool.py.
"""


def main() -> None:
    raise NotImplementedError("see module docstring")
