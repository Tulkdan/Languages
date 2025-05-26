(use sh)
(import cmd)

(defn strikethrough [text]
  (string "\e[9m" text "\e[0m"))

(def todo-file (string/format "%s/todo" (os/getenv "HOME")))

(def char-to-state {" " :todo "x" :done})
(def state-to-char (invert char-to-state))

(def task-peg (peg/compile
	       ~{:main (* (any (* :task (+ "\n" -1))) -1)
		 :state (cmt (* "- [" (<- (to "]")) "]") ,|(char-to-state $))
		 :text (/ (<- (to (+ "\n- [" -1))) ,string/trim)
		 :task (/ (* :state :text) ,|@{:state $0 :text $1})}))

(defn parse-tasks []
  (assert (peg/match task-peg (slurp todo-file))
	  "could not parse todo list"))

(def cols (scan-number ($<_ tput cols)))

(defn print-task [{:state state :text text}]
  (def decorate (case state
		  :done strikethrough
		  identity))
  (def prefix (string/format "- [%s] " (state-to-char state)))
  (def indent (string/repeat " " (length prefix)))
  (def wrap-width (- cols (length prefix)))
  (def wrapped-text ($< fold <,text -s -w ,wrap-width))
  (def lines (string/split "\n" wrapped-text))
  (eachp [i line] lines
	 (print
	  (if (= i 0) prefix indent)
	  (decorate line))))

(defn print-tasks [tasks]
  (each task (sort-by |(in $ :state) tasks)
    (print-task task)))

(defn first-word [str]
  (take-while |(not= $ (chr " ")) str))

(defn save-tasks [tasks]
  (def temp-file (string todo-file ".bup"))
  (with [f (file/open temp-file :a)]
	(each {:state state :text text} tasks
	  ($ printf -- "- [%s] %s\n" (state-to-char state) ,text >>,f)))
  ($ mv ,temp-file ,todo-file))

(cmd/defn to-done "cross something off" []
	  (def tasks (parse-tasks))
	  (def input @"")
	  (loop [[i {:state state :text text}] :pairs tasks
		 :when (= state :todo)]
	    (buffer/push-string input
				(string/format "%d %s" i text))
	    (buffer/push-byte input 0))

	  (when (empty? input)
	    (print "nothing to do!")
	    (os/exit 0))

	  (def output @"")
	  (def [exit-status]
	    (run fzf <,input >,output --height 10 --multi --print0 --with-nth "2.." --read0))
	  (def selections
	    (case exit-status
	      0 (drop -1 (string/split "\0" output))
	      1 []
	      2 (error "fzf error")
	      130 []
	      (error "unknown error")))

	  (each selection selections
	    (def task-index (scan-number (first-word selection)))
	    (def task (in tasks task-index))
	    (set (task :state) :done)
	    (print-task task))

	  (unless (empty? selections)
	    (save-tasks tasks)))

(defn append-task [text]
  (with [f (file/open todo-file :a)]
	(file/write f (string/format "- [ ] %s\n" text)))
  (print-task {:state :todo :text text}))

(cmd/defn to-do "add or list tasks"
	  [task (optional ["<task>" :string])]
	  (if task
	    (append-task task)
	    (print-tasks (parse-tasks))))

(cmd/main (cmd/group "A very simple task manager."
		     do to-do
		     done to-done))
