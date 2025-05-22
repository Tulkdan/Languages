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

(def services
  (seq [host :in hosts
	:when (host :online)
	service :pairs (host :services)]
       service))

(pp services)
