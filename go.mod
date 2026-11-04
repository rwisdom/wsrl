module github.com/rwisdom/wsrl

go 1.25.0

// forever is a pinned submodule at third_party/forever (376a6aa288), vendored into vendor/.
// After `make proto` inside the submodule: `go mod tidy` then `go mod vendor`; go.sum is generated.
require github.com/wowsims/forever v0.0.0

require (
	golang.org/x/exp v0.0.0-20250408133849-7e4ce0ab07d0 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)

replace github.com/wowsims/forever => ./third_party/forever
