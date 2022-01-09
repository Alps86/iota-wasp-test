build-sc:
	make -C sc

build-frontend:
	cd frontend && npm run build

start-frontend: build-frontend
	cd frontend && npm run start

start-node:
	cd docker && docker-compose up

destroy-node:
	cd docker && docker-compose down -v

init-node:
	wasp-cli chain deploy --committee=0 --quorum=1 --chain test --description="only a test"
	wasp-cli chain deposit IOTA:10000
	wasp-cli chain deploy-contract wasmtime test "test" sc/test/sc_bg.wasm