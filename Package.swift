// swift-tools-version:5.3
import PackageDescription

let package = Package(
	name: "WormholeWilliamSwift",
	products: [
		.library(
			name: "WormholeWilliamSwift",
			targets: ["WormholeWilliamGo"]
		),
	],
	targets: [
		.target(
			name: "WormholeWilliamSwift"
		),
		.binaryTarget(
			name: "WormholeWilliamGo",
			path: "Frameworks/WormholeWilliamGo.xcframework"
		),
		.testTarget(
			name: "WormholeWilliamSwiftTests",
			dependencies: ["WormholeWilliamSwift", "WormholeWilliamGo"]
		),
	]
)
