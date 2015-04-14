TimeSync CLI
============

A CLI for the OSL's TimeSync Application.


When you

Configuration:
--------------
The program can take configuration from environment variables or from a
configuration file in `~/.timesync.json` or `./.timesync.json`.
Flags can also be passed. The config file specifies things which will almost
always be the same, like the domain where the app is running (i.e. `localhost`
or `timesync.osuosl.org`). Flags specify variables which will change
frequently, like the activity of the duration. If no flags are passed, then it
will pop open an editor, `git commit` style, with info filled in.

Usage:
------

* `timesync checkin` 
	- `-d`: duration, an integer, required.
	- `-u`: user, optional.
	- `-p`: project slug, required.
	- `-a`: activity, required.
	- `-n`: notes, optional.
	- `-i`: issue URI, optional.
	- `-t`: date worked, optional, defaults to current date.

If no flags are passed your favorite editor will be opened with the previous
check in's information already filled in.

