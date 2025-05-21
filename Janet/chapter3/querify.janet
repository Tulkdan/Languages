(import sqlite3)

(defmacro querify [schema-file]
    ~(defn ,(symbol table-name) [conn]
        (sqlite3/eval conn ,(string/format "select * from %s;" table-name)))

    (def conn (sqlite3/open ":memory:"))
    (sqlite3/eval conn (string (slurp schema-file)))
    (def tables
        (->> (sqlite3/eval conn "select name from sqlite_schema where type = 'table';")
            (map |($ :name))
            (filter |(not (string/has-prefix? "sqlite_" $)))))
    (sqlite3/close conn)

    ~(upscope
        ,;(map table-definition tables)))

(querify "schema.sql")

(defn main [&]
    (def conn (sqlite3/open "db.sqlite"))
    (pp (people conn))
    (pp (grudges conn)))
