#!/bin/sh

go run main.go && cat head.html test.html tail.html > out.html
