module u2fsign

go 1.17

// Windows Fix
replace github.com/bearsh/hid => github.com/bdwyertech/hid v1.3.1-0.20210915214714-a7960f4cb384

require (
	github.com/marshallbrekka/go-u2fhost v0.0.0-20210111072507-3ccdec8c8105
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/bearsh/hid v1.3.0 // indirect
	golang.org/x/sys v0.0.0-20210915083310-ed5796bab164 // indirect
)
