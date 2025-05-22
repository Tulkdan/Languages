(def hosts [
	    {:name "claudius"
	     :ip "45.63.9.183"
	     :online true
	     :services
	       [{:name "janet.guide"}
		{:name "bauble.studio"}
		{:name "ianthehenry.com"}]}
	    {:name "caligula"
	     :ip "45.63.9.184"
	     :online false
	     :services [{:name "basilica.horse"}]}])

(each host hosts
  (if (host :online)
    (each service (host :services)
      (print (service :name)))))

(loop [host :in hosts
       :when (host :online)
       service :in (host :services)]
  (print (service :name)))

(print)

(each host hosts
  (if (host :online)
    (let [ip (host :ip)]
      (eachp [service-name available] (host :services)
	     (if available
	       (for instance 0 3
		 (pp [ip service-name instance])))))))

(loop [host :in hosts
       :when (host :online)
       :let [ip (host :ip)]
       [service-name available] :pairs (host :services)
       :when available
       instance :range [0 3]]
  (pp [ip service-name instance]))
