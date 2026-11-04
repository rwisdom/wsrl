"""plan section 7 order step 1; scaffolding accept check.

TODO:
- test_trivial_export: rlNew on example_setup returns a handle >= 0; rlClose.
- test_spell_ids_match_kit: handle.spell_ids() equals the Setup's spell_ids and lib/restosham/spell_ids.go.
"""


def test_loads(dll):
    assert dll.path.exists()


def test_trivial_export(dll):
    # rlFree on NULL is a no-op; it proves the DLL loads and an export resolves and is callable.
    dll.free(None)
