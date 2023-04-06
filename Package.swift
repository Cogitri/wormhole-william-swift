// swift-tools-version:5.3
import PackageDescription

let package = Package(
	name: "WormholeWilliamSwift",
	products: [
		.library(
			name: "WormholeWilliamSwift",
			targets: ["WormholeWilliam"]
		),
	],
	targets: [
		.target(
			name: "WormholeWilliamSwift"
		),
		.binaryTarget(
			name: "WormholeWilliam",
			path: "Frameworks/WormholeWilliam.xcframework"
		),
		.testTarget(
			name: "WormholeWilliamSwiftTests",
			dependencies: ["WormholeWilliamSwift", "WormholeWilliam"]
		),
	]
)
