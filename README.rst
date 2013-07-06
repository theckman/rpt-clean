RPT-Cleaner
===========

This is used for trimming the `ArmA 2`_ RPT log provided by the
`United Operations`_ community. Their RPT log contains additional information,
and this project strips out the noise and trims the data down.

.. _Arma 2: http://www.arma2.com/
.. _United Operations: http://forums.unitedoperations.net/index.php/page/

Not much to say here, just compile the go program and then run using the
rpt-trimmer shell script::

    cd $RPT_TRIMMER_DIR
    go build rpt-trimmer.go
    ./puller.sh
