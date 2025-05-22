(defn print-dots []
  (while true
    (prin ".")
    (flush)
    (ev/sleep 0)))

(ev/call print-dots)

(def f (file/open "lorem-ipsum.txt" :r))
(print "About to read")
(def bytes (file/read f 10))
(print "Done reading")
(file/close f)
(print "Done closing the file")
(printf "read %q" bytes)
(os/exit 0)
