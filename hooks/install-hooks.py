import os
import stat
from common import get_root_path

def make_executable(file_path):
    st = os.stat(file_path)
    os.chmod(file_path, st.st_mode | stat.S_IEXEC)


def install_hook(hook_name, hook_file, root_path):
    hook_src_path = os.path.join(root_path, "hooks", hook_file)
    make_executable(hook_src_path)
    hook_dest_path = os.path.join(root_path, ".git", "hooks", hook_name)
    if os.path.isfile(hook_dest_path):
        os.remove(hook_dest_path)
    os.symlink(hook_src_path, hook_dest_path)
    make_executable(hook_dest_path)
    print("{} hook installed.".format(hook_name))


if __name__ == "__main__":
    root_path = get_root_path()
    os.chdir(root_path)

    install_hook("pre-commit", "pre-commit.py", root_path)
