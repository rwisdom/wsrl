"""Shared fixtures (plan, File map: tests).

TODO:
- example_setup: setups/example.json as a wsrl.setups.Setup.
- handle: one fresh Handle on example_setup, closed after the test.
- rlsim_bin: path to the built cmd/rlsim binary (go build into build/), skip when missing.
"""

import pytest

from wsrl.dll import REPO_ROOT, Dll, default_path


@pytest.fixture(scope="session")
def repo_root():
    return REPO_ROOT


@pytest.fixture(scope="session")
def dll_path():
    path = default_path()
    if not path.exists():
        pytest.skip(f"DLL not built: {path} (run make dll)")
    return path


@pytest.fixture(scope="session")
def dll(dll_path):
    return Dll(dll_path)
