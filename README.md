# semver-highest

A tool to determine the highest version number that's smaller than a target version number.

## Installation

```bash
go install github.com/marten-seemann/semver-highest@latest
```

## Usage

```
./semver-highest -target v0.2.0 -versions v0.1.0,v0.1.1,v0.3.0 # v0.1.1
```

By default, pre-releases are skipped
```
./semver-highest -target v0.2.0 -versions v0.1.0,v0.1.1,v0.1.2-alpha,v0.3.0 # v0.1.1
./semver-highest -target v0.2.0 -versions v0.1.0,v0.1.1,v0.1.2-alpha,v0.3.0 -prerelease # v0.1.2-alpha
```
