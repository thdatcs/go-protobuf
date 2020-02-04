IMAGE=thdatcs/go-protobuf

ci-base:
	docker build --rm -f ci.base.Dockerfile --build-arg IMAGE=$(IMAGE) -t base-$(IMAGE) .

ci-lint:
	docker build --rm -f ci.lint.Dockerfile --build-arg IMAGE=$(IMAGE) -t lint-$(IMAGE) .
	docker run --rm -i lint-$(IMAGE)

ci-test:
	docker build --rm -f ci.test.Dockerfile --build-arg IMAGE=$(IMAGE) -t test-$(IMAGE) .
	docker run --rm -i test-$(IMAGE)

ci-build:
	docker build --rm -f ci.build.Dockerfile --build-arg IMAGE=$(IMAGE) -t $(IMAGE) .