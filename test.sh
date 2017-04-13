#!/bin/sh

die() {
  echo "$*"
  exit 1
}

out=$(pivotal-deliver </dev/null)
[ $? -eq 1 ] || die "test 1"
echo "$out" | grep -q PIVOTAL_TOKEN || die "test 2"

out=$(PIVOTAL_TOKEN=abcdefg pivotal-deliver </dev/null)
[ $? -eq 1 ] || die "test 3"
echo "$out" | grep -q PIVOTAL_PROJECT || die "test 4"

out=$(echo "foo #4444 bar #143629003 baz" | PIVOTAL_TOKEN=abcdefg PIVOTAL_PROJECT=12345 pivotal-deliver)
[ $? -eq 1 ] || die "test 5"
echo "$out" | grep -q 4444 && die "test 6"
echo "$out" | grep -q 143629003 || die "test 7"
echo "$out" | grep -q "Error fetching stories from Pivotal" || die "test 8"

echo "PASS"
