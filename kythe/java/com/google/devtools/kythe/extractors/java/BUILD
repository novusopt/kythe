package(default_visibility = ["//kythe:default_visibility"])

java_library(
    name = "java",
    srcs = [
        "InputUsageRecord.java",
        "JavaCompilationUnitExtractor.java",
        "ProcessAnnotation.java",
        "UsageAsInputReportingFileManager.java",
        "UsageAsInputReportingJavaFileObject.java",
    ],
    deps = [
        "//kythe/java/com/google/devtools/kythe/common",
        "//kythe/java/com/google/devtools/kythe/extractors/shared",
        "//kythe/java/com/google/devtools/kythe/extractors/shared:file_vnames",
        "//kythe/java/com/google/devtools/kythe/platform/java:options",
        "//kythe/java/com/google/devtools/kythe/platform/shared:metadata",
        "//kythe/java/com/google/devtools/kythe/util:delete_recursively",
        "//kythe/proto:analysis_java_proto",
        "//kythe/proto:buildinfo_java_proto",
        "//kythe/proto:java_java_proto",
        "//kythe/proto:storage_java_proto",
        "//third_party/javac",
        "@com_google_code_findbugs_jsr305//jar",
        "@com_google_guava_guava//jar",
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:protobuf_java",
        "@javax_annotation_jsr250_api//jar",
    ],
)
