# TensorFlow external dependencies that can be loaded in WORKSPACE files.
load("//vendor:repo.bzl", "flywk_http_archive")


# Sanitize a dependency so that it works correctly from code that includes
# TensorFlow as a submodule.
def clean_dep(dep):
  return str(Label(dep))




def flywk_workspace(path_prefix="", tf_repo_name=""):
  # Note that we check the minimum bazel version in WORKSPACE.
  flywk_http_archive(
          name = "pcre",
          build_file_content = all_content,
          strip_prefix = "pcre-8.43",
          build_file = clean_dep("//vendor:pcre.BUILD"),
  )
