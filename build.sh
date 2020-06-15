export PKG_CONFIG=/Users/sahaschitlange/go/bin/pkg-config

go generate ./stdlib ./libflux/go/libflux
go generate ./libflux/go/libflux
go build ./cmd/flux