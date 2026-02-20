.PHONY: all docs
all: speakeasy docs

original.yaml:
	curl https://docs.api.epilot.io/webhooks.yaml > original.yaml

original_modified.yaml: original.yaml overlay.yaml
	speakeasy overlay apply -s original.yaml -o overlay.yaml > original_modified.yaml

overlay.yaml:
	speakeasy overlay compare -a original.yaml -b original_modified.yaml > overlay.yaml

speakeasy:
	$(eval TMP := $(shell mktemp -d))
	curl https://docs.api.epilot.io/webhooks.yaml > $(TMP)/openapi.yaml
	speakeasy overlay apply -s $(TMP)/openapi.yaml -o overlay.yaml > $(TMP)/final.yaml
	speakeasy generate sdk --lang terraform -o . -s $(TMP)/final.yaml

docs:
	go mod tidy
	go generate ./...

