sol:
	solc  --abi ./contract/Store.sol -o build --overwrite
	solc --bin ./contract/Store.sol -o build --overwrite
	/home/chicmic/go/bin/abigen --abi=./build/Voting.abi --bin=./build/Voting.bin --pkg=voting --out=./voting/voting.go