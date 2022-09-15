#!/bin/bash
kill -9 $(lsof -i :80 | grep LISTEN | awk '{print $2}')
kill -9 $(lsof -i :443 | grep LISTEN | awk '{print $2}')