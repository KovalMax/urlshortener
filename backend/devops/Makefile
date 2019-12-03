.DEFAULT_GOAL := help
.PHONY: help run stop

help:
	@echo ''
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
	@echo ''

KEYS := build exec
APP_NAME ?= "url-shorter"

define LOOPBODY
  ifneq ($$(filter $$(KEYS),$(v)),)
    RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
    $(eval $(RUN_ARGS):;@:)
  endif
endef

$(foreach v,$(firstword $(MAKECMDGOALS)),$(eval $(LOOPBODY)))


##Commands:
start: ## docker start applications
	cp -n .env.dist .env || true
	@bash scripts/app_start.sh $(APP_NAME)

stop: ## docker stop applications
	@bash scripts/app_stop.sh $(APP_NAME)

exec: ## docker exec some application for example make exec redis sh - execute redis service container
	docker-compose -p $(APP_NAME) exec $(RUN_ARGS)
%:
@:
