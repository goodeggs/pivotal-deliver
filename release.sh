#!/usr/bin/env bash
set -e
set -o pipefail

indent() {
  sed -u 's/^/       /'
}

version=$(cat ../../VERSION)

make test

go get github.com/mitchellh/gox
go get github.com/tcnksm/ghr

gox -osarch "darwin/amd64 linux/amd64" -ldflags "-X main.VERSION=$version" -output "releases/$version/{{.OS}}_{{.Arch}}/pivotal-deliver"

rm -rf "releases/$version/dist" && mkdir -p "releases/$version/dist"
cp "releases/$version/darwin_amd64/pivotal-deliver" "releases/$version/dist/pivotal-deliver-Darwin-x86_64"
cp "releases/$version/linux_amd64/pivotal-deliver" "releases/$version/dist/pivotal-deliver-Linux-x86_64"

echo "releasing v${version}..."

ghr -t "$GITHUB_TOKEN" -u goodeggs -r pivotal-deliver --replace "v$version" "releases/$version/dist/"
