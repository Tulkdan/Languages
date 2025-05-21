(defmacro set-x [& expressions]
    (print "expanding the set-x macro")
    (printf "  here are my arguments: %q" expressions)
    (def result ~(upscope
        ,;(mapcat (fn [expression]
            [~(print ,(string/format "about to execute %q" expression))
                expression])
        expressions)))
    (printf "   and i'm going to return: %q" result)
    result)
