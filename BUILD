#gazelle is a too that prepares bazel for go specifics (generate BUILD.bazel files for all subpackages etc)
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/klahssen/bazel
gazelle(
    name = "gazelle",
    prefix = "github.com/klahssen/bazel",
)

# bazel run //:gazelle
