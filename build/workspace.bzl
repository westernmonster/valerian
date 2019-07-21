# TensorFlow external dependencies that can be loaded in WORKSPACE files.


# Sanitize a dependency so that it works correctly from code that includes
# TensorFlow as a submodule.
def clean_dep(dep):
  return str(Label(dep))

