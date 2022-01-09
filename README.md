# iota-wasp-test

wasp-cli request-funds
wasp-cli chain deploy --committee=0 --quorum=1 --chain test --description="only a test"
wasp-cli chain deposit IOTA:10000
wasp-cli chain deploy-contract wasmtime test "test" sc/test/sc_bg.wasm