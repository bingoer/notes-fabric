#!/bin/bash
rm -rf data/*
ps -ef | grep order | awk '{print $2}'| xargs kill -9
