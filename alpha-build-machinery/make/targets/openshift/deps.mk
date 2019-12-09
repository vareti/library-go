self_dir :=$(dir $(lastword $(MAKEFILE_LIST)))
scripts_dir :=$(self_dir)/../../../scripts

# We need to force localle so different envs sort files the same way for recursive traversals
deps_diff :=LC_COLLATE=C diff --no-dereference -N

update-deps:
	$(scripts_dir)/$@.sh
.PHONY: update-deps

# $1 - temporary directory to restore vendor dependencies from go.sum
define restore-deps
	ln -s $(abspath ./) "$(1)"/current
	cp -R -H ./ "$(1)"/updated
	$(RM) -r "$(1)"/updated/vendor
	cd "$(1)"/updated && go mod vendor
	cd "$(1)" && $(deps_diff) -r {current,updated}/vendor/ > updated/gomodules.diff || true
endef

verify-deps: tmp_dir:=$(shell mktemp -d)
verify-deps:
	$(call restore-deps,$(tmp_dir))
	@echo $(deps_diff) '$(tmp_dir)'/{current,updated}/gomodules.diff
	@     $(deps_diff) '$(tmp_dir)'/{current,updated}/gomodules.diff || ( \
		echo "ERROR: Content of 'vendor/' directory doesn't match 'go.sum' and the overrides in 'gomodules.diff'!" && \
		echo "If this is an intentional change (a carry patch) please update the 'gomodules.diff' using 'make update-deps-overrides'." && \
		exit 1 \
	)
.PHONY: verify-deps

update-deps-overrides: tmp_dir:=$(shell mktemp -d)
update-deps-overrides:
	$(call restore-deps,$(tmp_dir))
	cp "$(tmp_dir)"/{updated,current}/gomodules.diff
.PHONY: update-deps-overrides
