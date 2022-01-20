PROTOC=protoc 
PBPATH=./protobuf
PBFILE=athena.proto

protobuf: $(PBPATH)/$(PBFILE)
	$(PROTOC) -I=$(PBPATH) --go_out=$(PBPATH) $(PBPATH)/$(PBFILE)

