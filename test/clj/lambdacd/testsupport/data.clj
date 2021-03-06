(ns lambdacd.testsupport.data
  (:require [lambdacd.util :as utils]
            [lambdacd.event-bus :as event-bus]
            [clojure.core.async :as async]))


(defn- some-ctx-template []
  (let [config {:home-dir    (utils/create-temp-dir)}
        step-results-channel (async/chan (async/dropping-buffer 100))]
    (-> {:initial-pipeline-state   {} ;; only used to assemble pipeline-state, not in real life
         :step-id                  [42]
         :result-channel           (async/chan (async/dropping-buffer 100))
         :step-results-channel     step-results-channel
         :pipeline-state-component nil ;; set later
         :config                   config
         :is-killed                (atom false)
         :_out-acc                 (atom "")}
        (event-bus/initialize-event-bus))
    ))

(defn- add-pipeline-state-component [template]
  (if (nil? (:pipeline-state-component template))
    (assoc template :pipeline-state-component
                    (lambdacd.internal.default-pipeline-state/new-default-pipeline-state (atom (:initial-pipeline-state template))
                                                                                         (:config template)
                                                                                         (:step-results-channel template)))
    template))

(defn some-ctx []
  (add-pipeline-state-component
    (some-ctx-template)))

(defn some-ctx-with [& args]
  (add-pipeline-state-component
    (apply assoc (some-ctx-template) args)))