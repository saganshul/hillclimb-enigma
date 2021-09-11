for f in TESTS/*ct.txt
do
 echo "Running on $f"
 { time ./hillclimb $f &> $f.ans.txt; } &> $f.time.txt &
done

