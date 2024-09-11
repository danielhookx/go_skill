#!/bin/bash

go test -benchmem -run=^$ -bench ^BenchmarkClient$ -cpuprofile=client.out http_test.go
go test -benchmem -run=^$ -bench ^BenchmarkClientPerHostPool$ -cpuprofile=client_pool.out http_test.go
