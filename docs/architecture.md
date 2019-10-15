# architecture #

a high-level overview of what might be

## level 1 ##

```
+---------------+
| Web site      |
| html document |          
| javascript  --+--->-+
|               |     |
+---------------+     |
                      V
                      |
                      | [1] e.g. https://vulcantracker.biz/api/v1/beacon
                  +---+------------------+
                  | Ambassador/Proxy     |
                  |                      |
                  | /api/v1/beacon       |
                  | /api/v1/beacon/log   |
                  | ...                  |
                  |                      |
                  +---+------------------+
                      |
                +-----+
                |
 Beacon pod     V                          Katric pod
+---------------------+                     +-------------------+
|Collector/Producer   |                     |Processor/Consumer |
| [2]                 |                     | [5]               |
| /api/v1/beacon/     |                     | * Validation      |
| /api/v1/beacon/log  +                     | * Post-processing |
|                     |                     | * Modeling        |
| [3] Handlers        |                     | * Sequencing      |
|                     +---+            +--->| * Unifying        | [6a]
|                     |   |            |    |                   |------>+--------+
+---------------------+   |            |    +------------+------+       | Scylla |
                          V            ^                 |              |        | * this small box
                          | [4]        |                 |              +----^---+   is deceptively
                        +-+------------+-+               | [6b]              |       complex.
                        |                |               |                   |
                        | Redis Queue    +---------------+                   |
                        |                |               |                   |
                        +----------------+               +-------------+     |
                                                                       |     |
  e.g. https://report.vulcantracker.biz/              Scuttlebutt pod  |     |
+-------------------+                                +-----------------+-+   |
| Ambassador        |             [7]                |                   |   | [6b]
| API Gateway       | -----------------------------> | * Real time       +---+
|                   |             [8]                | * Composit model  |
|                   | <----------------------------> |                   |
+---+---------+-----+                                |                   |
    ^         ^                                      |                   |
    |         |                                      |                   |
    | [7]     | [8]                                  |                   |
    |         |                                      +-------------------+
    |         |
    |         |
    |         V
+===+===============+
= Analytics Report  =
=                   =
= ...               =
=                   =
= ...               =
=                   =
= ...               =
=                   =
+===================+

```
[1] - The request is load balanced through Ambassador and sent to the
      Beacon app pool, using GET or Web Beacon API.  The requests include:

      /api/v1/beacon     - Web beacon payloads.
      /api/v1/beacon/log - Server log (nginx, apache) payloads.

[2] - Beacon immediately responds (if GET) with HTTP 204, or nothing if using
      Web Beacon API.

[3] - Beacon generalizes, hashes and pre-anonymizes web beacon and server log
      data.

[4] - Payload data is submitted to Redis as a job for Katric to transform.

[5] - Katric performs further validation and post-processing before emitting to
      subscribers and Scylla the data record.

[6] - A composit model of analytics data is logged to Scylla and Scuttlebutt
      subscribers.

[7] - Requests for analytic reports on sites uses a Scuttlebutt instance to
      subscribe to composit events which update the report model(s), and
      generate an initial model to work with;

[8] - Web socket channel established to handle streaming updates from
      Scuttlebutt.

## level 2 ##

A little more thought:

* Steps 1, 5a can make use of [Ambassador](https://www.getambassador.io/) to
  handle api authorization against configured sites, so we don't have garbage
  requests from sites, with the added benefit of a highly performant proxy.

* Steps 3, 5b can be separated with message queue. (Redis?)

* Steps 6a, 6b, 7, 8, don't need as much care in that presumably they would have
  orders of magnitude less traffic from clients which are reading data for
  analytic reports.  Typical proxy/load balancing and Scuttlebutt docker instances
  should cover requests for reports and scale as needed.

## level 3 ##

App pods are container instances which comprise an area of the architecture
outlined above.  [Kubernetes](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/)
provides an excellent framework to run distributed systems resiliently.  Here,
a [pod](https://kubernetes.io/docs/concepts/workloads/pods/pod/) is a collection
of one or more containers. Daniel Sanche has an excellent introduction to
kubernetes specific nomenclature [here](https://medium.com/google-cloud/kubernetes-101-pods-nodes-containers-and-clusters-c1509e409e16).

I. Beacon pod

Beacon's primary job is to serve as an api endpoint for web beacon posts and
server log files.  The data is only minimally prepared, leaving combining and
transforming or other processes to Katric, so that:

  1. Data is prepared for sharded store by generalizing the subject (IP, time
     frame of visit, fingerprint information, etc), and tracker ID used.
  2. *optional* Server log files are prepared and submitted via Redis or
     directly to Katric.
  2. Data is immediately submitted to a queue via Redis or directly to Katric.

Redis is an excellent choice for delivering messages between processes in a
distributed system.  We can further cache and store messages from large volumes
of traffic, such that our platform's primary API endpoint for beacon data can
easily be distributed with load balancing.

II. Katric pod

Katric's primary job is taking the payload directly or via message queue and
transforming this into information for the WCS/ScyllaDB.  This includes taking
the rather unstructured data and further validating it against structures, then
pushing this into ScyllaDB via gocql.

III. Scuttlebutt pod

Scuttlebutt's primary job is to serve as an api endpoint and websocket
connection for getting analytics reports and updates.
