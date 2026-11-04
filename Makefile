# wsrl build targets. See docs/RL_Plan.md, "wsrl scaffolding".
# Windows: run with `mingw32-make` (MSYS2 ucrt64, same dir as gcc). The msys `make` strips USERPROFILE and Go
# then cannot find its module cache. cgo needs gcc on PATH. protoc and protoc-gen-go are needed by `proto`.

FOREVER := third_party/forever
DLL     := build/wsrl.dll

.PHONY: proto vendor dll test agreement bump clean

# Generate forever's sim/core/proto/*.pb.go inside the submodule (gitignored upstream). Once per pin.
proto:
	$(MAKE) -C $(FOREVER) proto

# Refresh vendor/ after a pin move. Check that assets/database/db.bin came across (with_db embeds it).
vendor:
	go mod tidy
	go mod vendor
	@test -f vendor/github.com/wowsims/forever/assets/database/db.bin || (echo "db.bin missing from vendor/" && exit 1)

# The c-shared DLL. with_db embeds the item and spell database; without it rlNew fails on every item id.
dll:
	mkdir -p build
	CGO_ENABLED=1 go build -mod=vendor -buildmode=c-shared -o $(DLL) --tags=with_db ./lib

# Go tests (agent, encounter, engine, attribution) then the Python tests.
test: dll
	go test -mod=vendor --tags=with_db ./...
	python -m pytest tests -q

# Agreement Test against the addon's Lua (section 10).
agreement: dll
	python -m pytest tests/agreement -q

# Move the forever pin: make bump COMMIT=<sha>. Then follow "Upgrading forever" in the plan.
bump:
	cd $(FOREVER) && git fetch && git checkout $(COMMIT)
	$(MAKE) proto vendor
	go build -mod=vendor --tags=with_db ./...

clean:
	rm -rf build
