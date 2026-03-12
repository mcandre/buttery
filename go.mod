module github.com/mcandre/buttery

go 1.26.1

require (
	github.com/andybons/gogif v0.0.0-20140526152223-16d573594812
	github.com/anthonynsimon/bild v0.14.0
	github.com/magefile/mage v1.16.1
	github.com/mcandre/mx v0.0.44
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/alexkohler/nakedret/v2 v2.0.6 // indirect
	github.com/kisielk/errcheck v1.9.0 // indirect
	github.com/mcandre/mage-extras v0.0.37 // indirect
	github.com/mcandre/tuco v0.0.21 // indirect
	golang.org/x/exp/typeparams v0.0.0-20250408133849-7e4ce0ab07d0 // indirect
	golang.org/x/mod v0.33.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/tools v0.42.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	honnef.co/go/tools v0.6.1 // indirect
)

tool (
	github.com/alexkohler/nakedret/v2/cmd/nakedret
	github.com/kisielk/errcheck
	github.com/magefile/mage
	github.com/mcandre/tuco/cmd/tuco
	honnef.co/go/tools/cmd/staticcheck
)
