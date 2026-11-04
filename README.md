# wsrl

<!-- TODO: ten lines. Clone with submodules, install protoc + protoc-gen-go, `make proto`, `make vendor`,
     `make dll`, `pip install -e .[test]`, `pytest`. Link to docs/RL_Plan.md (local, gitignored) and GLOSSARY.md. -->

Restoration Shaman healing policy trained by reinforcement learning against the WoWSims Forever simulator.

Prerequisites on Windows: Go 1.25 or newer (an older Go fetches the toolchain itself), MSYS2 ucrt64 gcc on PATH
for cgo, `protoc` and `protoc-gen-go` on PATH (unzip the protoc release into `~/go` so `~/go/bin/protoc`
and `~/go/include` exist; `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`), and GNU
make as `mingw32-make` (MSYS2 ucrt64, already on PATH with gcc). Do not use `C:\msys64\usr\bin\make.exe`:
its shell drops `USERPROFILE`, so Go cannot find its module cache or fetch the toolchain.

```sh
git clone --recurse-submodules <url> wsrl
mingw32-make proto vendor dll
pip install -e .[test]
pytest
```

`vendor/` is committed, so a fresh clone that only needs to build and train can skip `make proto vendor`
until the forever pin moves.
