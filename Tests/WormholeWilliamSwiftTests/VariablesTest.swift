import XCTest

@testable import WormholeWilliam

class VariablesTests: XCTestCase {
	func testVariables() {
		XCTAssert(WormholeWilliamGetWormholeCLIAppID().contains("lothar.com/wormhole/text-or-file-xfer"))
		XCTAssert(WormholeWilliamGetDefaultRendezvousURL().contains("ws://relay.magic-wormhole.io:4000/v1"))
		XCTAssert(WormholeWilliamGetDefaultTransitRelayAddress().contains("transit.magic-wormhole.io:4001"))
	}
}
