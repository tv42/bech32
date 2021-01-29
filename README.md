# eagain.net/go/bech32 -- Human-friendly encoding for binary data

```
go get eagain.net/go/bech32
```

Package `bech32` implements the "Bech32" encoding as specified in https://github.com/bitcoin/bips/blob/master/bip-0173.mediawiki

Note that Bech32 is **not** RFC 4648/3548, for that see [encoding/base32](http://golang.org/pkg/encoding/base32/).
It is also **not** `z-base-32`, for that see [github.com/tv42/zbase32](https://github.com/tv42/zbase32).
Compared to those, Bech32 adds a tag identifying type of stored data and a checksum.
Bech32 was popularized by Bitcoin.

## Command line utilities

Included are simple command-line utilities for encoding/decoding data.
Example:

```console
$ echo hello, world | bech32-encode greet
greet1dpjkcmr09ss8wmmjd3jq54mcuzl
$ bech32-decode greet1dpjkcmr09ss8wmmjd3jq54mcuzl
hello, world
$ printf '\x01binary!!!1\x00' | bech32-encode hckrspk
hckrspk1q93xjmnpwfujzgfpxyqqs8a2ax
$ bech32-decode hckrspk1q93xjmnpwfujzgfpxyqqs8a2ax | hexdump -C
00000000  01 62 69 6e 61 72 79 21  21 21 31 00              |.binary!!!1.|
0000000c
```

## Origins

At the beginning of this project, the library was copied from [filippo.io/age/internal/bech32](https://pkg.go.dev/filippo.io/age@v1.0.0-beta6/internal/bech32) as of version v1.0.0-beta6, in order to make it available outside of the `age` project.
For more information on `age`, see https://age-encryption.org/

It seems that this code was an improved and much cleaned up version of [github.com/btcsuite/btcutil/bech32](https://pkg.go.dev/github.com/btcsuite/btcutil/bech32), that had internally maintained at Google until it was open sourced as part of `age`.
