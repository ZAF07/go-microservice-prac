echo 'export GOPATH=$HOME/Go' >> $HOME/.bashrc source $HOME/.bashrc

protoc --proto_path=proto --go_out=proto/auto --go_opt=paths=source_relative greet.proto greet_service.proto