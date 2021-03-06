load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("//tools:build_rules/testing.bzl", "shell_tool_test")

package(default_visibility = ["//kythe:default_visibility"])

go_binary(
    name = "kindex_file_generator",
    srcs = ["kindex_file_generator.go"],
    deps = [
        "//kythe/go/platform/kindex",
        "//kythe/go/util/kytheuri",
        "//kythe/proto:storage_go_proto",
        "@go_protobuf//:proto",
    ],
)

shell_tool_test(
    name = "kindex_file_generator_test",
    script = ["""
    tmp="$$(mktemp -d 2>/dev/null || mktemp -d -t kythetest)"
    trap 'rm -rf "$$tmp"' EXIT ERR INT

    kindex_output="$$tmp/temp.kindex"
    $$generator \\
      --uri "kythe://kythe?lang=bazel?path=kythe/BUILD#//kythe/BUILD" \\
      --output "$$kindex_output" \\
      --source "$$input"

    size="$$(du -k "$$kindex_output" | cut -f1)"
    if [[ "$$size" -eq 0 ]] ; then
       echo "FAILED: output $$kindex_output is empty"
       exit 1
    fi

    $$kindex -explode "$$kindex_output"

    if [[ ! -e "$${kindex_output}_UNIT" ]] ; then
      echo "FAILED: Exploded kindex UNIT file not found"
      exit 1
    fi
    """],
    tools = {
        "generator": ":kindex_file_generator",
        "kindex": "//kythe/cxx/tools:kindex_tool",
        "input": "//kythe/testdata:BUILD",
    },
)
