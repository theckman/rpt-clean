#!/bin/sh

filename='/tmp/rpt-full.txt'

curl 'http://arma2.unitedoperations.net/dump/SRV1/SRV1_RPT.txt' | \
grep -P 'has been.*?by|has died at|####|has bled out' \
> $filename

