# bazel

Simple example of bazel + gazelle for go projects


`bazel run //:gazelle` to automatically generate BUILD files in each repo
`bazel build //cmd/...` to build all targets in cmd

Might need to to `bazel run //:gazelle -- update-repos github.com/klahssen/bazel`
in case of an error about inexistant package