
coverage:
	go test -coverprofile=cover.out ./...

	go tool cover -html=cover.out -o cover-inbound-admin.html

	go tool cover -func cover.out

mock: 
	# for generate mockgen service
	# how to use

# make mock INTERFACE_NAME=icheckout_repo
	
	GOPATH=$(shell make get-gopath)
	@echo "set gopath $(GOPATH)"

ifdef MOCKGEN_VERSION
	@echo "Mockgen version $(MOCKGEN_VERSION)"
else
	@echo "Mockgen Not found"
endif
	echo "running mockgen interface service: interface: $(INTERFACE_NAME)"

	# SERVICE_NAME=folder-service INTERFACE_NAME=IFolderService
	mockgen -source=interfaces/${INTERFACE_NAME}.go -destination=mocks/mock_${INTERFACE_NAME}.go 
	@echo "success!"
