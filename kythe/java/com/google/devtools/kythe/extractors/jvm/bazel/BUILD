package(default_visibility = ["//kythe:default_visibility"])

action_listener(
    name = "extract_kindex",
    extra_actions = [":extra_action"],
    mnemonics = ["JavaIjar"],
    visibility = ["//visibility:public"],
)

extra_action(
    name = "extra_action",
    cmd = "$(location :bazel_jvm_extractor) $(EXTRA_ACTION_FILE) $(output $(ACTION_ID).jvm.kindex)",
    out_templates = ["$(ACTION_ID).jvm.kindex"],
    tools = [":bazel_jvm_extractor"],
)

java_binary(
    name = "bazel_jvm_extractor",
    srcs = ["BazelJvmExtractor.java"],
    data = ["//third_party/javac:javac_jar"],
    main_class = "com.google.devtools.kythe.extractors.jvm.bazel.BazelJvmExtractor",
    deps = [
        "//kythe/java/com/google/devtools/kythe/extractors/jvm:jvm_extractor",
        "//kythe/java/com/google/devtools/kythe/extractors/shared",
        "//kythe/java/com/google/devtools/kythe/platform/indexpack",
        "//third_party/bazel:extra_actions_base_java_proto",
        "@com_google_guava_guava//jar",
        "@com_google_protobuf//:protobuf_java",
    ],
)
