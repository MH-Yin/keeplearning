#!/usr/bin/env python2.7

"""
Script for git pre-commit hook. It does the following:
* Run go test in all packages, fail if there is panic and failures
"""

import subprocess
import sys
import time

disabled_packages = [
    "chaincfg/chainhash",
    "base",
    "configs",
    "tools", # used
    "logger",
    "model/merkle",
    "rpc/btcjson",
    "rpc/rpcserver/websocket",
    "mvvm",
]


def install_golangciLint():
    subprocess.check_call(["brew", "install", "golangci/tap/golangci-lint"])

def error_out(s):
    return "\033[1;31;40m" + s + "\033[0m"

def success_out(s):
    return "\033[1;32;40m" + s + "\033[0m"

def run_command(type,command,msg):
    if type == "popen":
        p = subprocess.Popen(command, shell=True,stdout=subprocess.PIPE)
        out = p.stdout.readlines()
        if len(out)!=0:
            for line in out:
                print(line.strip())
            return(msg)
    if type == "call":
        p = subprocess.call(command, shell=True)
        if p != 0:
            return(msg)
    if type == "custom":
        p = subprocess.Popen(command, shell=True,stdout=subprocess.PIPE)
        out = p.stdout.readlines()
        return out

def check_golint():
    msg = run_command("popen",
        """golint `go list ./... | grep -E -v  "(vendor|%s)"`""" % "|".join(disabled_packages),
        "Presubmit FAILED! It is means that the code does not conform to the golint specification"
        )
    if msg != None:
        sys.exit(error_out(msg))

def check_golang_cilint():
    pkg = run_command("custom",
    """echo `go list ./... | grep -E -v  "(vendor|%s)"`""" % "|".join(disabled_packages), "")
    pkg = str(pkg[0].strip())
    print pkg
    pkgs = pkg.split(' ')
    failed = []
    for d in pkgs:
        d = d[31:]
        if d[-1] == "'":
            d = d[:len(d)-1]
        print(d)
        err = run_command("call",
            """golangci-lint run %s""" % "".join(d),d
        )
        if err != None :
            failed.append(err)
    if len(failed)!=0:
        print("Failed packages:")
        for f in failed:
            print(f)
        sys.exit("Presubmit FAILED! golangci-lint checking error")

def check_testcase():
    err = run_command("call",
        """go test `go list ./... | grep -E -v  "(vendor|%s)"`""" % "|".join(disabled_packages),
        "Presubmit FAILED! It is means that some test cases are useless"
    )

    if err != None:
        sys.exit(error_out(err))

if __name__ == "__main__":
    # check test case
    print("Begin to check testcase......")
    time.sleep(2)
    check_testcase()
    # check golint
    print("Begin to check golint......")
    check_golint()
    time.sleep(2)
    # check golang-cilint
    print("Begin to check golangci-lint......")
    check_golang_cilint()

    print(success_out("Presubmit PASSED. Submitting change..."))