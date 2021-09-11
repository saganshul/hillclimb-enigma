upperlim=50

for ((i=1; i<=upperlim; i++)); do
   if cmp -s "TESTS/test.$i.config.txt" "TESTS/test.$i.ct.txt.ans.txt" ; then
      echo ""
   else
      echo "Files are not matching for $i"
   fi
done