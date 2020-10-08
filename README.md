# Timestamp

[![Build Status](https://travis-ci.org/rickb777/timestamp.svg?branch=master)](https://travis-ci.org/rickb777/timestamp/builds)
[![Issues](https://img.shields.io/github/issues/rickb777/timestamp.svg)](https://github.com/rickb777/timestamp/issues)

Timestamp is a simple utility, similar to the Unix `date` command. Timestamp outputs the current time in a format
suitable for using in URLs as a marker for a version. The purpose is to support far-future expiry of assets
by giving each new version of an asset a distinct URL.

To this end, the output format can be to any integer base between 2 and 36. A useful choice is a base-36 number,
because it is a simple way to contract the number of characters needed, and therefore keep the URL short.

The precision can be set anywhere between days and nanoseconds, with seconds being the default.

## Examples

    timestamp
    timestamp -precision sec -base 10

both print the number of seconds since the Unix Epoch, which is the same result as given by

    date '%+s'

A more compact timestamp will be produced by

    timestamp -precision min -base 36

which specifies minutes precision and is ideal for asset tagging if the assets are released no more 
often than once per minute. You might even prefer once per hour or per day.

Instead of the 1st Jan 1970 Epoch, a different inception point can be specified, so

    timestamp -zero 2014-01-01

will print in base-36 the number of minutes since 1st Jan 2014, this being a much smaller number at
the time of writing.

A short-as-possible example might be

    timestamp -zero 2014-12-01 -precision day -base 36

Instead of the current time being printed as a timestamp, it is also possible to specify the 
output value explicitly. This would use `timestamp` simply for number base conversion

    timestamp -value 100 -base 36

prints the number 100 in base 36, i.e. "2s". This is appropriate if you generate your own sequence of version numbers instead of using the time and you want to benefit from base-36 URL contraction.
