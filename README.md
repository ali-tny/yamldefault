# yamldefault: a tiny CLI to read a yaml and optionally update the default value

A CLI to read a yaml file with a number of file system locations, with a default key specifying
which to return (which can be optionally overwritten).

My specific use case is via binding it to a bash function `todo` which opens up notes files - so I
can context switch for a day of personal work via `todo personal`, for example.

## Installation

Run `make build` and put the binary somewhere in your path (or just `go install` if $GOPATH is
reachable).

## Usage

Example yaml file, in say `~/.notes_locations.yml`:
```
DEFAULT: company
LOCATIONS:
  company: ~/company/some/path/to/notes
  personal: ~/personal/notes
```

then call (noting the default being overwritten):
```
➜  ~ echo $(yamldefault ~/.notes_locations.yml)
/Users/your_user/company/some/path/to/notes
➜  ~ echo $(yamldefault ~/.notes_locations.yml personal)
/Users/your_user/personal/notes
➜  ~ echo $(yamldefault ~/.notes_locations.yml)
/Users/your_user/personal/notes
```

In fact, I use a bash function to directly open a vim window via `todo` or `todo personal`:
```
function todo {
    vim $(yamldefault ~/.notes_locations.yml $1)
}
```
