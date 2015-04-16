TimeSync CLI
============

A CLI for the OSL's TimeSync Application.
Feel free to scribble all over this preliminary design document.


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
	- `--duration`: duration, an integer, required.
	- `--user`: user, optional.
	- `--project`: project slug, required.
	- `--activity`: activity, required.
	- `--notes`: notes, optional.
	- `--issue`: issue URI, optional.
	- `--date`: date worked, optional, defaults to current date.

If no flags are passed your favorite editor will be opened with the previous
check in's information already filled in.

