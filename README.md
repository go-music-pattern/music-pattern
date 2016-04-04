[![Build Status](https://travis-ci.org/go-pattern/pattern.svg?branch=master)](https://travis-ci.org/go-pattern/pattern) [![GoDoc](https://godoc.org/github.com/go-pattern/pattern?status.svg)](https://godoc.org/github.com/go-pattern/pattern) [![Coverage](https://img.shields.io/badge/coverage-35%-red.svg?style=flat)](https://gocover.io/github.com/go-pattern/pattern) [![Go Report Card](https://goreportcard.com/badge/github.com/go-pattern/pattern)](https://goreportcard.com/report/github.com/go-pattern/pattern) [![codebeat badge](https://codebeat.co/badges/cba9cc87-ae77-4894-9c3a-c8e3796f9b02)](https://codebeat.co/projects/github-com-go-pattern-pattern)

# Pattern

## Musical Pattern-to-Notes Language, parsed in Go

https://github.com/go-pattern/pattern

##### Credit

[Charney Kaye](http://w.charney.io)

[Outright Mental](http://w.outright.io)

### Time

To the X.J.Ontic API, time is specified as a time.Duration-since-epoch, where the epoch is the moment that player.Start() was called.

Internally, time is tracked as samples-since-epoch at the master output playback frequency (e.g. 48000 Hz). This is most efficient because source audio is pre-converted to the master output playback frequency, and all audio maths are performed in terms of samples.

### Patterns

1. Pattern is split into words (split on space " ")
2. Each word is parsed using regex to see if any of the following functions apply.

#### xN 

Repeats a note, e.g. `x2`:

    KCKx2 SNR MRCx2 KCK SNR

is equivalent to:

    KCK KCK SNR MRC MRC KCK SNR

#### Pick

Pick one note randomly from this comma-separated set, e.g. `[MRC,HAT]`:

    KCK [MRC,HAT] SNR [MRC,HAT]x2 KCK SNR [MRC,HAT]

#### Shuffle

Shuffle randomly and play each of the notes once, from this hyphen-separated set e.g. `-SNR-MRC-HAT-`

    KCK HAT -SNR-MRC-HAT-KCK-SNR-MRC-

#### Scratch

Randomly "scratch" beginning, end, how much, backwards/forwards, count

    T.B.D.
