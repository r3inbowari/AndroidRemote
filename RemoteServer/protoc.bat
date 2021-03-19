echo "protoc running"
cd hello
protoc *.proto --go_out=.
protoc.exe *.proto --go-grpc_out=.
cd ..

cd jpeg
protoc *.proto --go_out=.
protoc.exe *.proto --go-grpc_out=.
cd ..

cd touch
protoc *.proto --go_out=.
protoc.exe *.proto --go-grpc_out=.
cd ..

cd control
protoc *.proto --go_out=.
protoc.exe *.proto --go-grpc_out=.
cd ..