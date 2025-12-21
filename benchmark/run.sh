echo "Building dirgrep from source..."
cd ..
go build .
cd benchmark/
echo "Generating files..."
python3 generate_files.py
cd ..
time ./dirgrep --directory benchmark --recursive --pattern "A password forgot itself at dawn" > benchmark.txt