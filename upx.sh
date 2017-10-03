#!/bin/sh
UPX(){
	echo -n "Building $1..."
	upx $1 -o ./upxed/"upxed-"$1
}

for arg in "$@"
do
     UPX $arg
done
