load("//tools:build_rules/shims.bzl", "go_test")
load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//kythe:default_visibility"])

go_library(
    name = "packdb",
    srcs = ["packdb.go"],
    deps = [
        "//kythe/go/platform/indexpack",
        "//kythe/go/platform/kcd",
        "@go_protobuf//:proto",
    ],
)

go_test(
    name = "packdb_test",
    size = "small",
    srcs = ["packdb_test.go"],
    library = "packdb",
    visibility = ["//visibility:private"],
    deps = [
        "//kythe/go/platform/kcd/kythe",
        "//kythe/go/platform/kcd/testutil",
    ],
)
