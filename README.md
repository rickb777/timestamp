# Timestamp

Timestamp is a simple utility, similar to the Unix `date` command. Timestamp outputs the current time in a format
suitable for using in URLs as a marker for a version. The purpose is to support far-future expiry of assets
by giving each new version of an asset a distinct URL.

To this end, the default output format is a base-36 decimal number, with base-36 chosen because it is a simple
way to contract the number of characters needed, and therefore keep the URL short.

## Examples

    timestamp -precision sec -base 10

prints the number of seconds since the Unix Epoch, which is the same result as given by

    date '%+s'

The default

    timestamp

will print the number of minutes since the Epoch in base-36.

Instead of the 1st Jan 1970 Epoch, a different inception point can be specified, so

    timestamp -zero 2014-01-01

will print in base-36 the number of minutes since 1st Jan 2014, this being a much smaller number at
the time of writing.
