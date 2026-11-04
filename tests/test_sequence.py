"""plan section 7, cast sequence option.

TODO (no DLL needed; feed synthetic obs/mask):
- ordered mode consumes one row per call; timed mode waits until at_s; repeat expands rows;
  WAIT maps to the Wait index; exhaustion behaviors: wait, loop, handover.
"""
