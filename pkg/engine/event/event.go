// package event provides an event system that all other functionalities can hook on to
// and either subscribe or emit events
package event

import "github.com/simimpact/srsim/pkg/engine/event/handler"

type System struct {
	AttackEvent AttackEventHandler
	DamageEvent DamageEventHandler

	// test placeholder until we get actual events defined
	Ping handler.EventHandler[int]
}