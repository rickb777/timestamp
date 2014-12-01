# Timestamp

Timestamp is a simple utility, similar to the Unix `date` command. Timestamp outputs the current time in a format
suitable for using in URLs as a marker for a version. The purpose is to support far-future expiry of assets
by giving each new version of an asset a distinct URL.

To this end, the output format can be to any integer base between 2 and 36. A useful choice is a base-36 number,
because it is a simple way to contract the number of characters needed, and therefore keep the URL short.

The precision can be set anywhere between days and nanoseconds, with seconds being the default.

## Examples

    timestamp -precision sec -base 10

prints the number of seconds since the Unix Epoch, which is the same result as given by

    date '%+s'

The default

    timestamp

will print the number of seconds since the Epoch in base-10. A more compact timestamp will be produced by

    timestamp -precision min -base 36

which is ideal for asset tagging if the assets are released no more often than once per minute. You might
even prefer once per hour or per day.

Instead of the 1st Jan 1970 Epoch, a different inception point can be specified, so

    timestamp -zero 2014-01-01

will print in base-36 the number of minutes since 1st Jan 2014, this being a much smaller number at
the time of writing.
