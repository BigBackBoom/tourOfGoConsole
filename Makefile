GOGET := go get
GETFLAG = -d

setup:
	$(GOGET) $(GETFLAG) github.com/jessevdk/go-assets-builder

depend:
	@dep ensure

cleardep:
	@rm -rf vendor