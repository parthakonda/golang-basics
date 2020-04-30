#!/bin/bash
echo "Running load test on $1"
locust -f test.py --no-web -c 1000 -r 100 --run-time 10m --step-load --step-clients 1000 --step-time 10s --host $1