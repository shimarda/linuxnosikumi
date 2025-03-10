#!/usr/bin/python3

import mmap
import time
import datetime

ALLOC_SIZE = 100 * 1024 * 1024
ACCESS_UNIT = 10 * 1024 * 1024
PAGE_SIZE = 4096

def show_message(msg):
    print("{} : {}".format(datetime.datetime.now().strftime("%H:%M:%S"), msg))

show_message("新規メモリ領域獲得前。Enterで100MiBのメモリ領域を獲得")
input()

memregion = mmap.mmap(-1, ALLOC_SIZE, flags=mmap.MAP_PRIVATE)
show_message("新規メモリ領域を獲得。Enterで1秒に10MiBずつ、合計100MiBの新規メモリ領域にアクセス")
input()

for i in range(0, ALLOC_SIZE, PAGE_SIZE):
    memregion[i] = 0
    if i%ACCESS_UNIT == 0 and i != 0:
        show_message("{} MiBアクセス".format(i//(1024*1024)))
        time.sleep(1)

show_message("新規獲得したメモリ領域の全てへアクセス。Enterで終了")
input()