
# benchmarking flow

## OpCode

- To enable benchmarking

```sh
export ENABLE_BENCHMARKING=true 
```

- To print out opcode details on console
```sh
export OPCODE_DETAILS=true

./gno.land/build/gnoland start | grep "benchmark.Op"
```
The benchmarking contract is located at `examples/gno.land/r/x/benchmark/ops.gno`

- To execute the benchmarking functions, add and edit file located at `gno.land/genesis/genesis_txs.txt`

After we starts gnoland, it automatically generates a file benchmark.log

we need to run following command. It converts the binary format result in benchmarks.log to human readable csv in result.csv

```
./benchmarking/build/benchmark -path benchmarks.log

```

TODO:
- benchmark binary operations on different data type with large numbers.


## Store Access 
