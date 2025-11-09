solcjs --bin --abi SimpleCounter.sol -o build/

cd build

abigen --bin=SimpleCounter_sol_SimpleCounter.bin --abi=SimpleCounter_sol_SimpleCounter.abi --pkg=counter --out=simpleCounter.go
