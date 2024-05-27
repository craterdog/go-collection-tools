echo "Compiling the following tools:"
mkdir -p ./bin/
echo "	version"
go build -o ./bin/ ./src/version
echo "	validate"
go build -o ./bin/ ./src/validate
echo "	format"
go build -o ./bin/ ./src/format
echo "Done."
