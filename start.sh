#!/bin/sh
echo Starting
redis-server --protected-mode no \
             --dir /redis \
             --rdbcompression yes \
             --rdbchecksum yes \
             --maxmemory-policy noeviction \
             --maxmemory 128M \
             --appendonly yes \
             --appendfsync everysec \
             --auto-aof-rewrite-percentage 100 \
             --auto-aof-rewrite-min-size 64mb \
             --save 900 1 \
             --save 300 10 \
             --save 60 100 &
echo Starting main
/main

