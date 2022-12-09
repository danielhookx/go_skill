go build -o snippet_mem && GODEBUG='gctrace=1' ./snippet_mem >runningtrace.out 2>&1
