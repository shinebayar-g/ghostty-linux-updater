# ghostty-linux-updater

To keep my local [Ghostty](https://ghostty.org/) installation up to date with the upstream [tip release](https://github.com/ghostty-org/ghostty/releases/tag/tip).

## Prerequisites

- [Go](https://go.dev/doc/install) `1.25` or newer installed on your system.
- [Zig](https://ziglang.org/learn/getting-started/) `0.15.2` installed on your system. (Ghostty tip release [requires](https://ghostty.org/docs/install/build) Zig `0.15.2`, other versions aren't guaranteed to work.)
- [Minisign](https://jedisct1.github.io/minisign/) installed on your system. (Used to verify the authenticity of the downloaded tarball using the provided [signature](https://ghostty.org/docs/install/build#getting-the-source-code).)

## Usage

```bash
# Clone the repository
git clone https://github.com/shinebayar-g/ghostty-linux-updater.git
# Change to the repository directory
cd ghostty-linux-updater
# Run the update script
go run ./cmd/
```

Just run the above command whenever you want to update your local Ghostty installation to the latest tip release. The script will handle downloading, verifying, and installing the latest version for you.
