#!/bin/sh
set -e

SQLITE_FILEPATH='./sqlite3.db'
SQL_FILEPATH='./scripts/seeds.sql'

echo ">>> Checking if sqlite3 commands installed ..."
if ! [ $(which sqlite3) ] ; then
    echo "sqlite3 commands not found, please install it first."
    exit 1
else
    echo "OK"
fi

echo ">>> Checking SQLite3 database exists ..."

if [ -f $SQLITE_FILEPATH ] ; then
    echo "Found, will recreate with initial data."
    rm -rf $SQLITE_FILEPATH
else
    echo "No database file exists, will create new."
    touch $SQLITE_FILEPATH
fi

echo ">>> Loading initital data ..."
sqlite3 $SQLITE_FILEPATH < $SQL_FILEPATH
if [ $? -ne 0 ] ; then
    echo "Failed to load seeds data"
else
    echo "Done"
fi
