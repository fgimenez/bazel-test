load(
    "@bazel_tools//tools/build_defs/repo:http.bzl",
    "http_archive",
    "http_file",
)
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

# Additional bazel rules

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "08369b54a7cbe9348eea474e36c9bbb19d47101e8860cec75cbf1ccd4f749281",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.24.0/rules_go-v0.24.0.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.24.0/rules_go-v0.24.0.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "d4113967ab451dd4d2d767c3ca5f927fec4b30f3b2c6f8135a2033b9c05a5687",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.0/bazel-gazelle-v0.22.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.0/bazel-gazelle-v0.22.0.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "4521794f0fba2e20f3bf15846ab5e01d5332e587e9ce81629c7f96c793bb7036",
    strip_prefix = "rules_docker-0.14.4",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.14.4/rules_docker-v0.14.4.tar.gz"],
)

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_register_toolchains",
    "go_rules_dependencies",
)

go_rules_dependencies()

go_register_toolchains(
    go_version = "1.14"
)

load(
    "@bazel_gazelle//:deps.bzl",
    "gazelle_dependencies"
)

gazelle_dependencies()


load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()
