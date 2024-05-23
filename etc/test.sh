directory=tst
echo "Validating the test files:"
for file in `ls $directory`; do
	echo "    $file"
	bin/validate $directory/$file
done
echo

echo "Formatting the test files:"
for file in `ls $directory`; do
	# Go maps are not deterministic so skip them.
	if [ "${file:0:3}" != "map" ]; then
		echo "    $file"
		bin/format $directory/$file
	fi
done
echo
