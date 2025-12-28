echo "Building dirgrep from source..."
cd ..
go build .
cd benchmark/
echo "Generating files..."
python3 generate_files.py
cd ..
time ./dirgrep --directory benchmark --recursive --pattern "A password forgot itself at dawn" > benchmark_dirgrep.txt
cd benchmark/
time grep -r "A password forgot itself at dawn" > ../benchmark_grep.txt
time rg "A password forgot itself at dawn" > ../benchmark_ripgrep.txt