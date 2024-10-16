tilde
=====

Expand tilde `~` to home dir.


The expansion follows this specific flow:

On windows at first HOME is used, and if it's not present or is empty,
there goes %USERPROFILE% (via `os.UserHomeDir`).
It allows for having HOME pointed to some extra layer like Cygwin environment.

On other systems expansion uses `os.UserHomeDir` directly, so it is
$HOME on Unix or $home on Plan9.
