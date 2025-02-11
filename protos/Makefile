PROTO_DIRS = .
OUT_DIRS = ./golang

proto:
	@for dir in $(PROTO_DIRS); do \
	     protoc --go_out=$(OUT_DIRS) --go_opt=paths=source_relative \
	     	    --go-grpc_out=$(OUT_DIRS) --go-grpc_opt=paths=source_relative \
		    $$dir/*.proto; \
	done
