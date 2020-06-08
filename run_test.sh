#!/bin/env sh

g++ main.c -o a.out && \
        echo "g++ build success :)\nrunning test program...\n" && \
        ./shared-lib/output/libjudger.so \
        --exe_path="a.out" \
        --max_cpu_time=10000 \
        --max_real_time=10000 \
        --log_path="judger.log" \
        --input_path="1.in" \
        --output_path="runtime.out" \
        --max_output_size=10000 \
        --max_process_number=1024 \
        --uid=0 \
        --gid=0