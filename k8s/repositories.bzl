"""Register external repositories for python.
"""

load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def python_repositories():
    maybe(
        http_archive,
        name = "subpar",
        sha256 = "c5fa062eb1aff08fda7def89eb6cf914aca0efa696534aa9ea9e78621709394a",
        strip_prefix = "subpar-master",
        urls = ["https://github.com/google/subpar/archive/refs/heads/master.tar.gz"],
    )

    com_github_yaml_pyyaml_build_file = """\
    load("@rules_python//python:defs.bzl", "py_binary", "py_library")
    py_library(
        name = "yaml",
        srcs = glob(["lib/yaml/*.py"]),
        imports = [
            "lib",
        ],
        visibility = ["//visibility:public"],
    )
    py_library(
        name = "yaml3",
        srcs = glob(["lib3/yaml/*.py"]),
        imports = [
            "lib3",
        ],
        visibility = ["//visibility:public"],
    )
    """

    # YAML is used to parse files in the common/ directory that contain
    # information about opcodes, flags, and builtin classes and functions.
    http_archive(
        name = "com_github_yaml_pyyaml",
        build_file_content = com_github_yaml_pyyaml_build_file,
        sha256 = "f33eaba25d8e0c1a959bbf00655198c287dfc5868f5b7b01e401eaa1796cc778",
        strip_prefix = "pyyaml-6.0",
        urls = ["https://github.com/yaml/pyyaml/archive/refs/tags/6.0.tar.gz"],
    )
