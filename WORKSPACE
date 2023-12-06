workspace(
    name = "aspect-repro",
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

rules_go_version = "v0.43.0"
http_archive(
    name = "io_bazel_rules_go",
    patch_args = [
        "-E",
        "-p1",
    ],
    sha256 = "d6ab6b57e48c09523e93050f13698f708428cfd5e619252e369d377af6597707",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/{version}/rules_go-{version}.zip".format(version = rules_go_version),
        "https://github.com/bazelbuild/rules_go/releases/download/{version}/rules_go-{version}.zip".format(version = rules_go_version),
    ],
)


gazelle_version = "v0.34.0"
http_archive(
    name = "bazel_gazelle",
    patch_args = [
        "-E",
        "-p1",
    ],
    sha256 = "b7387f72efb59f876e4daae42f1d3912d0d45563eac7cb23d1de0b094ab588cf",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/{version}/bazel-gazelle-{version}.tar.gz".format(version = gazelle_version),
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/{version}/bazel-gazelle-{version}.tar.gz".format(version = gazelle_version),
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    version = "1.20.2",
)

http_archive(
    name = "build_aspect_cli",
    sha256 = "d30da98533e2d6949948f1c6b0b53942203b2e997be363e1fc5f5479e4944dd7",
    strip_prefix = "aspect-cli-5.5.4",
    urls = ["https://github.com/aspect-build/aspect-cli/archive/refs/tags/5.5.4.tar.gz"],
)

load("@build_aspect_cli//:go.bzl", aspect_cli_deps = "deps")

aspect_cli_deps()

http_archive(
    name = "aspect_gcc_toolchain",
    sha256 = "3341394b1376fb96a87ac3ca01c582f7f18e7dc5e16e8cf40880a31dd7ac0e1e",
    strip_prefix = "gcc-toolchain-0.4.2",
    urls = [
        "https://github.com/aspect-build/gcc-toolchain/archive/refs/tags/0.4.2.tar.gz",
    ],
)

load("@aspect_gcc_toolchain//toolchain:repositories.bzl", "gcc_toolchain_dependencies")

gcc_toolchain_dependencies()

load("@aspect_gcc_toolchain//toolchain:defs.bzl", "ARCHS", "gcc_register_toolchain")

gcc_register_toolchain(
    name = "gcc_toolchain_x86_64",
    sysroot_variant = "x86_64-X11",
    target_arch = ARCHS.x86_64,
)
