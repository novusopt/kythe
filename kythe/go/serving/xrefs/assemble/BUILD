load("//tools:build_rules/shims.bzl", "go_test")
load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//kythe:default_visibility"])

go_library(
    name = "assemble",
    srcs = ["assemble.go"],
    deps = [
        "//kythe/go/services/graphstore",
        "//kythe/go/services/graphstore/compare",
        "//kythe/go/services/xrefs",
        "//kythe/go/storage/stream",
        "//kythe/go/util/encoding/text",
        "//kythe/go/util/kytheuri",
        "//kythe/go/util/pager",
        "//kythe/go/util/schema/edges",
        "//kythe/go/util/schema/facts",
        "//kythe/go/util/schema/nodes",
        "//kythe/go/util/schema/tickets",
        "//kythe/proto:common_go_proto",
        "//kythe/proto:internal_go_proto",
        "//kythe/proto:serving_go_proto",
        "//kythe/proto:storage_go_proto",
        "@go_stringset//:stringset",
    ],
)

go_test(
    name = "assemble_test",
    size = "small",
    srcs = ["assemble_test.go"],
    library = "assemble",
    visibility = ["//visibility:private"],
    deps = [
        "//kythe/go/test/testutil",
        "@go_protobuf//:proto",
    ],
)
