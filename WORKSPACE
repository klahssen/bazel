workspace(name = "bazel")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
#load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

#load go rules
http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.5/rules_go-0.18.5.tar.gz"],
    sha256 = "a82a352bffae6bee4e95f68a8d80a70e87f42c4741e6a448bec11998fcc82329",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_klahssen_bazel",
    commit = "55dba024d1ea9f752e184987e52019bc1d613e71",
    importpath = "github.com/klahssen/bazel",
)
