"""Gate G3: the addon's own Lua over one episode's events equals the DLL's six weights (plan section 10).

TODO, per episode kind:
- replay each CSV row: set STUB.time, STUB.sp, STUB.critPct, STUB.haste, STUB.stats, STUB.maxMana from
  the stat columns and call addon:UpdatePlayerStats(); parsers.set_auras(lua, 408510) when the Water
  Shield cast row appears; cast row -> CastTracker:StartCast / seg:IncFillerCasts(mana) / FinishCast;
  heal or tick row -> parsers.heal(addon, "SPELL_HEAL" or "SPELL_PERIODIC_HEAL", id, amount, crit).
- end the segment; read seg.t and GetMP5, GetIntManaValue, GetSpiritRegenValue, GetHasteHPCT;
  normalize by seg.t.heal.
- assert: rel tol 1e-3 on +Healing, Crit, Haste; 2% on Intellect, Spirit, MP5 (Water Shield mana is
  analytic in the addon, actual energize in the DLL). On failure print the first diverging event.
"""
