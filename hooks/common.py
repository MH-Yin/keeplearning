"""
Defines some common utility methods for hooks scripts.
"""

import os

def get_root_path():
    hooks_dir = os.path.abspath(os.path.join(os.path.abspath(__file__), os.pardir))
    return os.path.abspath(os.path.join(hooks_dir, os.pardir))
