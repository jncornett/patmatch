/*
Package patmatch implements a simple library for string matching with template expressions.
This package supports 'mad-lib' style template matching on strings.

The list of format control sequences is as follows:

- %(name)s: a named capture group using the verb 's'
- %()s: an anonymous capture group using the verb 's'
- %s: an non-capturing group using the verb 's'
- %%: the literal character '%'

*/
package patmatch
