#!/bin/bash 

set -e

setup() {
  if [ "$(vagrant status | grep running)" == "" ]; then 
    vagrant up 
  fi

  vagrant ssh-config frontend_ci >> /tmp/lambdacd-dev-env-ssh-config
  vagrant ssh-config backend_ci >> /tmp/lambdacd-dev-env-ssh-config

  mkdir -p /tmp/mockrepo
  echo "[SUCCESS] You are good to go"
}

testallClojure() {
  lein test :all
}

testallClojureScript() {
  lein cljsbuild test
}

autotestClojureScript() {
  lein cljsbuild auto test
}


testall() {
    testallClojure && testallClojureScript
}

testunit() {
  lein test
}

release() {
  testall && lein clean && lein with-profile +release release $1 && scripts/github-release.sh
}

releaseLocal() {
  lein with-profile +release install
}

push() {
  testall && git push
}

serve() {
  lein run
}

serveClojureScript() {
  lein figwheel app
}

repl-server() {
  lein repl :headless :port 58488
}

if [ "$1" == "setup" ]; then
    setup
elif [ "$1" == "testall" ]; then
    testall
elif [ "$1" == "test" ]; then
    testunit
elif [ "$1" == "test-cljs" ]; then
    testallClojureScript
elif [ "$1" == "autotest-cljs" ]; then
    autotestClojureScript
elif [ "$1" == "release" ]; then
    release $2
elif [ "$1" == "release-local" ]; then
    releaseLocal
elif [ "$1" == "push" ]; then
    push
elif [ "$1" == "serve" ]; then
    serve
elif [ "$1" == "serve-cljs" ]; then
    serveClojureScript
elif [ "$1" == "repl-server" ]; then
    repl-server
else
    echo "usage: $0 <goal>

goal:
    setup         -- to set up your environment
    test          -- run unit tests
    testall       -- run all tests
    testcljs      -- run all ClojureScript tests (i.e. unit tests for frontend)
    autotest-cljs -- starts autotest-session for frontend
    serve         -- start a server with a demo-pipeline
    serve-cljs    -- compile clojurescript and watch for changes
    push          -- run all tests and push current state
    release       -- release current version
    release-local -- install current version in local repository
    repl-server   -- start a repl cursive can use to run tests in"
    exit 1
fi
