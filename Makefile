sol:
	solc --optimize --abi ./contracts/contract.sol -o build
	solc --optimize --bin ./contracts/contract.sol -o build
	/home/chicmic/go/bin/abigen --abi=./build/BalanceCheck.abi --bin=./build/BalanceCheck.bin --pkg=api --out=./api/contract.go