#!/bin/bash

set -x 

echo "generate by antrl4 tool "

java -jar /Users/xingyue/antlr/antlr-4.7.2-complete.jar -Dlanguage=Go -visitor -no-listener JsonQuery.g4 -o ./parser/

echo "download from github.com "

git clone https://github.com/nikunjy/rules tmp

cd tmp/parser/

cp bench_test.go bool_operation.go evaluate.go float_operation.go nester_error.go panic_test.go \
string_operation.go bool_operation.go int_operation.go null_operation.go parse_logical_test.go \
version_operation.go evaluate.go jsonquery_lexer.go jsonquery_visitor_impl.go operation.go parse_simple_test.go \
../../parser

echo "now you can extend this rule engine as you like ... "

cd ../../
rm -rf tmp