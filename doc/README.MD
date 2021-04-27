https://www.cnblogs.com/gutousu/p/9951956.html

protoc --java_out=. message.proto

生成model
protoc --proto_path=E:/Java/protobuf --java_out=E:/Java/protobuf message.proto

生成service
protoc --plugin=protoc-gen-grpc-java=E:/RPC/protoc/bin/protoc-gen-grpc-java-1.10.0-windows-x86_64.exe
--proto_path=E:/Java/protobuf --grpc-java_out=E:/Java/protobuf message.proto


protoc --grpc-java_out=. message.proto

protoc --plugin=protoc-gen-grpc-java-1.37.0-linux-x86_32.exe --grpc-java_out=./ message.proto

