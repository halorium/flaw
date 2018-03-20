## Flaw is a Go(lang) custom error package

### Features
This was created because sometimes you need a little more information than what
the builtin error interface can provide.

Flaw allows you to "Wrap" your errors as they bubble up adding a message trace
as well as a stack trace to your errors.

The Flaw interface implements the following interfaces:
* builtin error interface
* fmt.Stringer interface
* json.Marshaler interface

This allows for detailed error message output in JSON formatted logs like
Loggly and many others.

### Usage


#### Credits
This is a refactor from code designed and written by Tom Mornini and Loren Hale