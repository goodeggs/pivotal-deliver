pivotal-deliver
===============
A CLI to deliver Pivotal Tracker stories.

Usage
-----

We originally built this to scan git logs for Pivotal story ids and mark those stories delivered:

```
$ git log --format=full HEAD^..HEAD | pivotal-deliver
```

Stories must be in the `finished` state.

Development
-----------

```
$ brew install golang direnv
$ mkdir -p pivotal-deliver/src/github.com/goodeggs
$ cd pivotal-deliver
$ echo 'layout "go"' > .envrc
$ cd src/github.com/goodeggs
$ git clone https://github.com/goodeggs/pivotal-deliver.git
$ cd pivotal-deliver
$ make
```

We use [dep](https://github.com/golang/dep) for dependency management.

Releasing
---------

To create a release:

```
$ go get github.com/Clever/gitsem
$ gitsem {major,minor,patch}
$ git push
$ GITHUB_TOKEN=xxx ./release.sh
```

