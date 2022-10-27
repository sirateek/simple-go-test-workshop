run:
	go run main.go

start-dev-redis:
	docker start test-redis || docker run --name test-redis -p 6379:6379 -d redis:7.0.5

unit-test:
	ginkgo -r --label-filter="unit"

integration-test:
	ginkgo -r --label-filter="integration"