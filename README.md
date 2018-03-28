# Page

Pagination for stdin.

## Install

```sh
$ go get github.com/andystanton/page
```

## Usage

* Page size and page number can be passed in.
* Default page size is 8.
* Page number starts from 1.

```sh
$ ls | page -size 5 -num 3
```
