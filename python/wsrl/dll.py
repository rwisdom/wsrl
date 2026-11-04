"""ctypes bindings for build/wsrl.dll (plan section 7, order step 1).

TODO:
- class Dll: set argtypes and restype on EVERY export listed in lib/library.go
  (rlNew, rlReset, rlAdvance, rlAct, rlStepAll, rlObserve, rlMask, rlReward, rlAttribution,
  rlEvents, rlSpellIds, rlSpellTable, rlItemStats, rlClose). rlItemStats and rlSpellTable are
  module-level (no handle). Only rlFree is bound so far.
- class Handle: one per sim. new(setup_json) -> handle; reset(seed); advance() -> int; act(i) -> bool;
  observe(out: np.ndarray float32); mask(out: np.ndarray uint8); reward() -> float;
  attribution() -> np.ndarray (6); events() -> str (rlFree after copying); spell_ids() -> list[int];
  close(). Pass numpy buffers with arr.ctypes.data_as(POINTER(c_float)) etc.
- step_all(handles) -> np.ndarray int32 statuses, through rlStepAll.
- Never call two exports on one handle from two threads.

Test: tests/test_dll.py.
"""

from __future__ import annotations

import ctypes
import os
from pathlib import Path

REPO_ROOT = Path(__file__).resolve().parents[2]


def default_path() -> Path:
    """The WSRL_DLL env var, else <repo root>/build/wsrl.dll."""
    env = os.environ.get("WSRL_DLL")
    return Path(env) if env else REPO_ROOT / "build" / "wsrl.dll"


class Dll:
    def __init__(self, path: Path | str | None = None):
        self.path = Path(path) if path else default_path()
        self._lib = ctypes.CDLL(str(self.path))
        self._bind()

    def _bind(self) -> None:
        lib = self._lib
        lib.rlFree.argtypes = [ctypes.c_void_p]
        lib.rlFree.restype = None

    def free(self, p) -> None:
        self._lib.rlFree(p)
