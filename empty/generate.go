package empty

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go_opt=Mempty.proto=github.com/mniak/ps121/empty empty.proto
