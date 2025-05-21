(use ./set-x-verbose)

(defn main [&]
    (set-x
        (var sum 0)
        (for i 0 10
            (+= sum i))
        (print sum)))

(print)
(print "and this is what main looks like:")
(print)
(pp (disasm main))
