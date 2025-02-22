#!/bin/bash

false &
wait $! 
echo "falseコマンドをしゅうりょうしました: $?"