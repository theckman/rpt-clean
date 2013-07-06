RPT-Cleaner
===========

This is used for trimming the `ArmA 2 <http://www.arma2.com/>`_ RPT log
provided by the `United Operations <http://forums.unitedoperations.net/index.php/page/index.html`_
community. Their RPT log contains additional information, and this project
strips out the noise and trims the data down.

Not much to say here, just compile the go program and then run using the
rpt-trimmer shell script::

    cd $RPT_TRIMMER_DIR
    go build rpt-trimmer.go
    ./puller.sh
