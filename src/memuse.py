#!/usr/bin/python3

import subprocess

size = 500000

print("メモリ獲得前のシステム全体のメモリ量を表示")
subprocess.run("free")

array = [0]*size

print("メモリ獲得後のシステム全体のメモリの空き容量を表示")
subprocess.run("free")