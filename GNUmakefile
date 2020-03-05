PHONY: deps clean build

dirs = $(shell find * -type d -maxdepth 0)
baseDir = $(shell pwd)


test-all:
	@$(foreach dir,$(dirs), \
		echo $(dir); \
		cd $(baseDir); \
		cd $(dir); \
		go test -v ; \
	)
	cd $(baseDir)