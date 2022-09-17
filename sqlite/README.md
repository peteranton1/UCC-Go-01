SQLite Exercises
-------------------

# How to setup

Data from 
[kaggle Shakespeare](https://www.kaggle.com/datasets/kingburrito666/shakespeare-plays?resource=download)

Then install sqlite

    sudo apt install sqlite3

Unzip data and load into a local db

    cd ../data
    unzip Shakespeare_data.csv.zip
    sqlite3 shakespeare.db
    sqlite> .mode csv
    sqlite> .import Shakespeare_data.csv import
    sqlite> .schema
        ...
    sqlite> select * from import limit 10;
        ...
    sqlite> select count(*) from import;
        111396
    sqlite> create table plays (\
        rowid integer primary key, \
        play, linenumber, act, scene, \
        line, player, text);
    sqlite> .schema
        CREATE TABLE import(
        "Dataline" TEXT,
        "Play" TEXT,
        "PlayerLinenumber" TEXT,
        "ActSceneLine" TEXT,
        "Player" TEXT,
        "PlayerLine" TEXT
        );
        CREATE TABLE plays (rowid integer primary key, play, linenumber, act, scene, line, player, text);

    sqlite> insert into plays \
        select dataline as rowid, play, \
        playerlinenumber as linenumber, \
        substr(actsceneline, 1, 1) as act, \
        substr(actsceneline, 3, 1) as scene, \
        substr(actsceneline, 5, 5) as line, \
        player, playerline as text \
        from import;
    sqlite> select * from plays limit 10;
    ...

    sqlite> select * from plays \
        where text like "whether tis nobler";
    <not found>
    
    sqlite> create virtual table playsearch \
        using fts5(playsrowid, text);
    sqlite> insert into playsearch select \
        rowid, text from plays;

    sqlite> select * from playsearch \
        where text match "whether tis nobler";
        34231,"Whether 'tis nobler in the mind to suffer"

    sqlite> select play, act, scene, line, \
        player, plays.text \
        from playsearch inner join plays \
        on playsearch.playsrowid = plays.rowid \
        where playsearch.text match "whether tis nobler";

        Hamlet,3,1,65,HAMLET,"Whether 'tis nobler in the mind to suffer"

    sqlite> drop table import;
    sqlite> vacuum;
    sqlite> <CTRL-D>


