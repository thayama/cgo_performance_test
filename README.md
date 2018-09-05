# Benchmark for Logging Facility

Measure benchmarks for diffrent way to pass C message to logging
facility in Go.

Idea is to create a message in C, and pass it to Go. In Go, the
message is collected with channel.

* BenchmarkCallGoFromC
 * Pass static memory to Go. No copies in Go nor C.
* BenchmarkCallGoFromC2 
 * Pass static memory to Go. Make copy in Go before sending to channel.
* BenchmarkCallGoFromC3
 * Pass dynamically allocated message to Go.
* BenchmarkCallGoFromCNop
 * NOP. Measure base perforamnce of calling C from Go.

This benchmark calls C functions from Go. So the base measurement for NOP
must be subtracted frmo each benchamrk.
