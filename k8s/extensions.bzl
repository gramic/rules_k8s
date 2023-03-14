"""Top level extensions to download repos.
"""

load("//toolchains/kubectl:extensions.bzl", _kubectl_toolchain = "kubectl_toolchain")
load("//k8s:repositories.bzl", _python_repositories = "python_repositories")
load("//k8s:with-defaults.bzl", _k8s_defaults = "k8s_defaults")

kubectl_toolchain = _kubectl_toolchain

def _python_repositories_impl(_ctx):
    _python_repositories()

python_repositories = module_extension(
    implementation = _python_repositories_impl,
)

# k8s.defaults

defaults_attrs = {
    "cluster": attr.string(mandatory = False),
    "context": attr.string(mandatory = False),
    "image_chroot": attr.string(mandatory = False),
    "kind": attr.string(mandatory = False),
    "kubeconfig": attr.string(mandatory = False),
    "namespace": attr.string(mandatory = False),
    "resolver": attr.string(mandatory = False),
    "user": attr.string(mandatory = False),
}

_tag_attrs = {
    "name": attr.string(
        mandatory = True,
        doc = "The name of the repo.",
    ),
}
_tag_attrs.update(**defaults_attrs)
_defaults_tag = tag_class(attrs = _tag_attrs)

def _impl(ctx):
    for mod in ctx.modules:
        for tag in mod.tags.defaults:
            _k8s_defaults(
                name = tag.name,
                cluster = tag.cluster,
                context = tag.context,
                image_chroot = tag.image_chroot,
                kind = tag.kind,
                kubeconfig = tag.kubeconfig,
                namespace = tag.namespace,
                resolver = tag.resolver,
                user = tag.user,
            )

k8s = module_extension(
    implementation = _impl,
    tag_classes = {
        "defaults": _defaults_tag,
    },
)
