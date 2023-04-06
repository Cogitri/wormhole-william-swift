package WormholeWilliam

import (
	"github.com/psanford/wormhole-william/wormhole"
)

func GetWormholeCLIAppID() string {
	return wormhole.WormholeCLIAppID
}

func GetDefaultRendezvousURL() string {
	return wormhole.DefaultRendezvousURL
}

func GetDefaultTransitRelayAddress() string {
	return wormhole.DefaultTransitRelayAddress
}
