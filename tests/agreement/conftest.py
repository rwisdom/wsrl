"""Agreement Test fixtures (plan section 10).

TODO:
- sys.path.insert(0, repo_root / "third_party" / "hsw-forever" / "tests"); import harness, parsers.
  Do not copy them.
- addon(talent_tabs): parsers.start("SHAMAN", talent_tabs, ["Parsers/RestoShaman.lua"]) -> (lua, addon, seg).
  talent_tabs come from the Setup's talents string (Tidal Mastery, Mental Quickness, Improved Water Shield).
- load_spell_table(addon, rows): for every row from rlSpellTable,
  addon.Spells:Define(id, addon.SpellType.SHAMAN, name, {coeff, base, cast, mana}).
- set_constants(lua, addon): addon.CritRatingPerPct, addon.ManaPerInt, the db table SetupConversionFactors
  reads (hasteRatingPerPct, intPerCritOverride.SHAMAN) to the sim's values; replace
  addon.GetRegenPerSpirit with a function returning 0.125 (no knee at 50 Spirit in the sim).
- episodes: the three episode kinds (lhw_spam, baseline APL replayed through the DLL, a trained Policy),
  each as (events CSV, spell table, six DLL weights) produced by evaluate.py with attribution on.
"""
