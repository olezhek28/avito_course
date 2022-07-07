PHONY: bench
bench:
	go test -bench=BenchmarkGetDomainStat -count 5 | tee result.txt
	benchstat result.txt
	rm result.txt
