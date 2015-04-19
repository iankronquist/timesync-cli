TimeSync CLI
============

A CLI for the OSL's TimeSync Application.
Feel free to scribble all over this preliminary design document.

SYNOPSIS
`timesync command`

COMMANDS
--------

* `checkin`: Sends a message to the server to check the current user in.
	- `--duration`: duration, an integer, required.
	- `--project`: project slug, required.
	- `--activity`: activity, required.
	- `--user`: user, optional, defaults to current user's name.
	- `--date`: date worked, optional, defaults to current date on server side.
	- `--issue`: issue URI, optional.
	- `--notes`: notes, optional.

If no flags are passed your favorite editor will be opened with the previous
check in's information already filled in (not yet implemented).

EXAMPLES:
---------

This is probably the most basic and common use case:
```sh
$ timesync checkin -duration=12 -project=ww -activity=docs
```
Assuming my username is `iankronquist`, It will send the following data to the
server:
```json
{
  "duration": 12,
  "user": "iankronquist",
  "project": "ww",
  "activity": "docs",
}
```

CONFIGURATION:
--------------

The program can take configuration from environment variables or from a
configuration file in `~/.timesync.json` or `./.timesync.json`.
Command line flags can also be passed. The config file specifies things which
will almost always be the same, like the domain where the app is running (i.e.
`localhost` or `timesync.osuosl.org`). Flags specify variables which will
change frequently, like the activity of the duration.
