# subspace
File sharing application.

## Supported Platforms

| OS      | 386 | amd64 | arm6 | arm64 |
| ---     | --- | ----  | ---  | ----  |
| Linux   | ✅   | ✅    | ✅    | ✅     |
| Windows | ✅   | ✅    |      | ✅    |
| MacOS   |     | ✅    |      | ✅     |

## Building
Requirements:
- Golang 1.16+
- [Goreleaser](https://goreleaser.com/) (optional, to build linux packages)

```shell
go test ./...

mkdir -p ./licensing/my-licenses
touch ./licensing/my-licenses/make_embedFS_happy
go-licenses save github.com/jimmale/subspace --force --save_path="licensing/my-licenses"

go build -buildmode=pie -ldflags "-s -w" -o ./subspace ./cmd/main.go
```

### Build packages
`goreleaser release --rm-dist --snapshot`

## Dependencies
### Buildtime 

### Runtime
| Library                                                     | License | Purpose                                      |
| -------                                                     |---------|----------------------------------------------|
| [Sirupsen/Logrus](https://github.com/Sirupsen/logrus)       | MIT     | Pretty Logging                               |
| [urfave/cli/v2](https://github.com/urfave/cli/v2)           | MIT     | CLI args parsing, configuration file reading |
| [lucas-clemente/quic-go](github.com/lucas-clemente/quic-go) | MIT     | Peer Connection                              |

## TODO

- Peer discovery
  - Cache
  - Out of Band
  - Bittorrent Trackers
  - LAN multicast
  - DHT
  - Peer Exchange
- Cert Management
  - Generate new Cert
  - Keep Peer Certs
  - Trust Peer Certs
- Network Management
  - UPnP Port Forwarding
- Transfers
  - Transfer Resume
  - Compression
  - BT-style sharing (eg if Alice and Bob are downloading from Charlie, Alice and Bob will share chunks of the file amongst themselves)
- CI/CD
  - Build in licenses for `--license` invocation
  - Semantic Versioning
  - Include Buildinfo (git revision, release version, GOOS, GOARCH, Build Time)
  - RPM Packages
  - Deb Packages
  - Windows EXEs
  - OSX DMG