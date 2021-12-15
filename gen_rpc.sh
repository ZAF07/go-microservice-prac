echo 'export GOPATH=$HOME/Go' >> $HOME/.bashrc source $HOME/.bashrc


protoc --go_out=api/rpc --go_opt=paths=source_relative --go-grpc_out=proto/rpc --go-grpc_opt=paths=source_relative ./proto/greet/greet.proto ./proto/greet/greet_service.proto

echo "DONE"
# protoc -I greet.proto -I=$GOPATH/bin --go-grpc_out=:. greet.proto