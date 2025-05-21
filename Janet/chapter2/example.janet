(def skadi @{:name "Skadi" :type "German Shepherd"})
(def odin @{:name "Odin" :type "German Shepherd"})

(def people
    [{:name "ian" :dogs [skadi odin]}
        {:name "kelsey" :dogs [skadi odin]}
        {:name "jeffrey" :dogs []}])

(pp people)

(defn main[&]
    (set (odin :type)
        "Well mostly German Shepherd but he's mixed with some")
    (pp people))
