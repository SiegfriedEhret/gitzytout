# gitzytout

> git z'y tout !<br>
(git them all!)

A very small thing that reads a yaml config file and pushes to the given git repositories.

Made to solve the problem of having your code centralized only on Github or Gitlab...

It is made in [Go](https://golang.org) because [Vincent](https://github.com/vbehar) thinks this language is awesome (and yes, it is basically 2 commands in your shell).

## how to

Install using `go get -u gitlab.com/SiegfriedEhret/gitzytout`.

First, create a `gitzytout.yaml` file, with some content:

```yaml
main: git@gitlab.com:SiegfriedEhret/gitzytout.git
mirrors:
- git@github.com:SiegfriedEhret/gitzytout.git
```

Then, run `gitzytout` it will automagically configure your `.git/config` to allow you to push to multiple repositories.

## todo

- [ ] Check with https/ssl urls
- [x] Check with passphrase keys
- [x] [Add a main item](https://gitlab.com/SiegfriedEhret/gitzytout/issues/1) ?

## licence

```
Copyright (c) 2016 Siegfried Ehret <siegfried@ehret.me>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```
