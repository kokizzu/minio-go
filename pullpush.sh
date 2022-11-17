#!/usr/bin/env bash

set -x # print all command
set -e # exit on error

if [ $# -eq 0 ] ; then
  echo "Usage: 
  ./pullpush.sh 'the commit message'"
  exit
fi

git add .

git commit -m "$*" && git pull && git push origin master

TAG=$(ruby -e 't = Time.now; print "v1.#{t.month+(t.year-2021)*12}%02d.#{t.hour}%02d" % [t.day, t.min]')
git tag -a "$TAG" -m "$*"
git push --tags 

echo "# to undo this release: 
git tag -d $TAG
git push -d origin $TAG"
