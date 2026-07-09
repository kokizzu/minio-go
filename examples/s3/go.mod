module github.com/kokizzu/minio-go/examples/s3

go 1.25.0

require (
	github.com/cheggaaa/pb v1.0.29
	github.com/kokizzu/minio-go v0.0.0
	github.com/minio/sio v0.3.0
	golang.org/x/crypto v0.53.0
)

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/klauspost/cpuid/v2 v2.1.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/net v0.55.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	gopkg.in/ini.v1 v1.66.6 // indirect
)

replace github.com/kokizzu/minio-go => ../..
