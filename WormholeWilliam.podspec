Pod::Spec.new do |s|
    s.name = "WormholeWilliam"
    s.version = "0.0.4"
    s.osx.deployment_target = '11'

    s.prepare_command = "make"
    s.vendored_frameworks = "Frameworks/WormholeWilliam.xcframework"

    s.authors = { "Rasmus Thomsen" => "oss@cogitri.dev" }
    s.license = { :type => 'MIT', :file => 'LICENSE' }
    s.summary = "Swift bindings for wormhole-william (go)"
    s.homepage = "https://github.com/Cogitri/wormhole-william-swift.git"
    s.source  = { :git => "https://github.com/Cogitri/wormhole-william-swift.git", :tag => "v#{s.version}" }

end
