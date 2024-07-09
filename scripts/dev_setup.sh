#!/bin/bash

# This script is intended to be run on the root directory

tmux new-session -d -s dev 

tmux split-window -h
tmux split-window -t dev:0.0 -v
tmux split-window -t dev:0.2 -v

tmux send-keys -t dev:0.0 './scripts/pg_script.sh' C-m # Postgres DB
tmux send-keys -t dev:0.1 'air --build.exclude_dir "volumes"' C-m
tmux send-keys -t dev:0.2 './scripts/tailwind.sh' C-m

tmux attach-session -t dev
