#!/bin/sh

set -e 
CONTAINER=$1
HOSTPATH=$2
CONTPATH=$3

REALPATH=$(readlink --canonicalize $HOSTPATH)
FILESYS=$(df -P $REALPATH | tail -n 1 | awk '{print $6}')

printf "REALPATH - %s\n" "$REALPATH"
printf "FILESYS - %s\n" "$FILESYS"

df > /tmp/df

while read DEV BLOCK USED AVAIL USE MOUNT 
do [ $MOUNT = $FILESYS ] && break
done < /tmp/df
[ $MOUNT = $FILESYS ]

printf "DEV - %s\n" "$DEV"

SUBPATH=$(echo $REALPATH | sed s,^$FILESYS,,)
DEVDEC=$(printf "%d %d" $(stat --format "0x%t 0x%T" $DEV))

printf "SUBPATH - %s\n" "$SUBPATH"
printf "DEVDEC - %s\n" "$DEVDEC"

docker-enter $CONTAINER sh -c \
	"[ -b $DEV ] || mknod --mode 0600 $DEV b $DEVDEC"

docker-enter $CONTAINER mkdir /tmpmnt
docker-enter $CONTAINER mount $DEV /tmpmnt
docker-enter $CONTAINER mkdir -p $CONTPATH
docker-enter $CONTAINER mount -o bind /tmpmnt/$SUBPATH $CONTPATH
docker-enter $CONTAINER umount /tmpmnt
docker-enter $CONTAINER rmdir /tmpmnt
