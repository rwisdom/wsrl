// Package register registers every agent factory the sim needs, exactly once (plan 3.2 step 1).
//
// Called by rlNew (lib/library.go) and by cmd/rlsim. Never call sim.RegisterAll(): it registers
// forever's heal-less Restoration Shaman stub, and core.RegisterAgentFactory panics on a second
// factory for the same spec (sim/core/agent.go:181-191).
package register

import (
	"sync"

	"github.com/wowsims/forever/sim/common"
	_ "github.com/wowsims/forever/sim/encounters" // registers the default raid boss
	"github.com/wowsims/forever/sim/warrior/protection"

	"github.com/rwisdom/wsrl/lib/restosham"
)

var once sync.Once

// Once registers the Protection Warrior, the common item and consumable effects, the encounters
// (by import side effect) and wsrl's own Restoration Shaman. Safe to call from every entry point.
func Once() {
	once.Do(func() {
		protection.RegisterProtectionWarrior()
		common.RegisterAllEffects()
		restosham.Register()
	})
}
