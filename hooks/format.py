#!/usr/bin/env python2.7

import re
import subprocess
import os
from common import get_root_path

tracked_keywords = ["modified", "new file"]

def format_opened_go_files():
    for f in opened_go_files():
        print ("Formatting {}".format(f))
        subprocess.check_call(["gofmt", "-w", f])

def check_code():
        path = get_root_path()
        shell_path = path + '/hooks/checkcode.sh'
        subprocess.check_call([shell_path])

def opened_go_files():
    out = subprocess.check_output(["git", "status"])
    pattern = r"({}):\s*(\S+\.go)".format("|".join(tracked_keywords))
    files = []
    for line in out.splitlines():
        line = line.decode("utf-8")
        match_obj = re.search(pattern, line)
        if match_obj:
            files.append(match_obj.group(2))
    return files

if __name__ == "__main__":
    format_opened_go_files()
    check_code()

