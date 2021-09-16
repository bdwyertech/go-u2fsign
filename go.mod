module github.com/bdwyertech/go-u2fsign

go 1.17

// Windows Fix
replace github.com/bearsh/hid => github.com/bdwyertech/hid v1.3.1-0.20210915214714-a7960f4cb384

require (
	github.com/hashicorp/go-hclog v0.14.1
	github.com/hashicorp/go-plugin v1.4.3
	github.com/marshallbrekka/go-u2fhost v0.0.0-20210111072507-3ccdec8c8105
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/bearsh/hid v1.3.0 // indirect
	github.com/fatih/color v1.7.0 // indirect
	github.com/golang/protobuf v1.3.4 // indirect
	github.com/hashicorp/yamux v0.0.0-20180604194846-3520598351bb // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/mitchellh/go-testing-interface v0.0.0-20171004221916-a61a99592b77 // indirect
	github.com/oklog/run v1.0.0 // indirect
	golang.org/x/net v0.0.0-20210916014120-12bc252f5db8 // indirect
	golang.org/x/sys v0.0.0-20210915083310-ed5796bab164 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55 // indirect
	google.golang.org/grpc v1.27.1 // indirect
)

