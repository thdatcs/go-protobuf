IMAGE=thdatcs/go-protobuf

ci-base:
	docker build --rm -f ci.base.Dockerfile --build-arg IMAGE=$(IMAGE) -t base-$(IMAGE) .

ci-lint:
	docker run --rm -i base-$(IMAGE) bash -c "make lint"

ci-unit-test:
	docker run --rm -i base-$(IMAGE) bash -c "make unit-test"

ci-integration-test:
	docker-compose -f ci.test.docker-compose.yml up integration-test

ci-build:
	docker build --rm -f ci.build.Dockerfile --build-arg IMAGE=$(IMAGE) -t $(IMAGE) .