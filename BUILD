load ("@io_bazel_rules_go//go:def.bzl", "go_library", "go_binary")

go_library(
	name = "demo",
	srcs = [
		"demo.go",
		"unsafe.c",
		"unsafe.h",
	],
	cdeps = [
		## Something
	],
	cgo = True,
	clinkopts = ["-ldl"],
	importpath = "aspect/demo",
)

go_binary(
	name = "demobin",
	srcs = ["main.go"],
	deps = [":demo"],
)
