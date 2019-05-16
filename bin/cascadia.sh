#!/usr/bin/env bash

/Users/timfulmer/go/bin/cascadia -i ${1} -c 'text' -o | sed -e 's/<[^>]*>//g' >> /Users/timfulmer/go/src/legislation-analysis/txt/legislation.txt