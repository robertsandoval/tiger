#!/bin/bash

DATE=`date +"%Y-%m-%d"`
BASEDIR=/usr/local/ocp/clusters
CLUSTERNAME=$1
echo "script: $CLUSTERNAME"
INSTALLPATH="$BASEDIR/$USER/$DATE/$CLUSTERNAME"


# create install dir structure
# /usr/local/ocp/installs/$USER/$DATE/$CLUSTERNAME

echo "Install location $INSTALLPATH"

if [ ! -d "$INSTALLPATH" ] 
then
	mkdir -p $BASEDIR/$USER/$DATE/$CLUSTERNAME
else
    echo "Error: Directory $INSTALLPATH already exists."
    exit 0
fi
