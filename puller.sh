#!/bin/sh

full_filename='/tmp/rpt-full.txt'
ext_filename='/tmp/rpt-ext.txt'
simple_filename='/tmp/rpt-simple.txt'
whereami=`pwd`

curl 'http://arma2.unitedoperations.net/dump/SRV1/SRV1_RPT.txt' | \
grep -P 'has been.*?by|has died at|####|has bled out' | \
tee $full_filename | \
$whereami/rpt-trimmer -e | \
tee $ext_filename | \
$whereami/rpt-trimmer -s \
> $simple_filename
