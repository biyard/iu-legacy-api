SERVICE ?= $(shell basename `git rev-parse --show-toplevel`)
ENV ?= local
ARTIFACTS_DIR ?= .aws-sam/build/Function
BUILD_DIR = $(ARTIFACTS_DIR)
BUILD_CONFIG_FILE = $(BUILD_DIR)/config.yaml

ifeq ("$(ENV)", "prod")
	DomainName ?= api.incheon.world
endif

ifeq ("$(ENV)", "stg")
	DomainName ?= stg-api.incheon.world
endif

ifeq ("$(ENV)", "dev")
	DomainName ?= dev-api.incheon.world
	ADMIN_CODE ?= a2FbaIy9et65GwI
endif

DomainName ?= dev-api.incheon.world
AWS_KEY_ID ?= $(shell aws configure get aws_access_key_id)
AWS_SECRET_KEY ?= $(shell aws configure get aws_secret_access_key)
ADMIN_CODE ?= "test"
FEEPAYER_KEY ?=
OWNER_KEY ?=

DOCKER_FLAGS ?=

RESTORE_KEY ?= 
BOT_TOKEN ?=
HINT_SECRET ?= 

run: $(BUILD_CONFIG_FILE)
	@go run . -config $(BUILD_CONFIG_FILE)

.PHONY: $(BUILD_CONFIG_FILE)
$(BUILD_CONFIG_FILE):
	@echo "Building config file for $(ENV)"
	@mkdir -p $(BUILD_DIR)
	@cp ./fixtures/config-$(ENV).yaml  $@
	@cp ./fixtures/firebase-credential.json  $(BUILD_DIR)/firebase-credential.json
	@sed -i 's|{AWS_KEY_ID}|$(AWS_KEY_ID)|g' $@
	@sed -i 's|{AWS_SECRET_KEY}|$(AWS_SECRET_KEY)|g' $@
	@sed -i 's|{ADMIN_CODE}|$(ADMIN_CODE)|g' $@
	@sed -i 's|{FEEPAYER_KEY}|$(FEEPAYER_KEY)|g' $@
	@sed -i 's|{OWNER_KEY}|$(OWNER_KEY)|g' $@
	@sed -i 's|{RESTORE_KEY}|$(RESTORE_KEY)|g' $@
	@sed -i 's|{BOT_TOKEN}|$(BOT_TOKEN)|g' $@
	@sed -i 's|{HINT_SECRET}|$(HINT_SECRET)|g' $@

.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $(ARTIFACTS_DIR)/bootstrap -buildvcs=false .
	@cp ./fixtures/firebase-credential.json  $(ARTIFACTS_DIR)/firebase-credential.json

clean:
	@rm -rf $(BUILD_DIR) .aws-sam

build-Function:

build-sam: clean $(BUILD_CONFIG_FILE)
	@sam build
	DOCKER_FLAGS="$(DOCKER_FLAGS)" $(MAKE) .aws-sam/build/Function/bootstrap

.aws-sam/build/Function/bootstrap: $(BUILD_CONFIG_FILE)
	docker run --rm -v $(PWD):/workdir $(DOCKER_FLAGS) --workdir /workdir public.ecr.aws/sam/build-provided.al2 sh -c "make build"

test-sam: build-sam
	@sam local invoke Function

run-sam: build-sam
	@sam local start-api -p 3000

deploy-sam: build-sam deploy-only-sam

deploy-only-sam:
	@sam deploy --config-file samconfig.toml --no-confirm-changeset --stack-name $(ENV)-$(SERVICE) --parameter-overrides Stage=$(ENV) --parameter-overrides DomainName=$(DomainName) --tags Stage=$(ENV) --no-fail-on-empty-changeset
