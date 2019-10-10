# architecture #

a high-level overview of what might be

## level 1 ##

+---------------+
| Web site      |
| html document |                 Beacon                       Katric
| javascript ---+-----+ [1]     +-------------------+        +-------------------+
|               |     |         |Collector/Producer |        |Processor/Consumer |
+---------------+     |         |                   |        | [4]               |
                      +-------->| /api/v1/beacon/   |  [3]   | * Validation      |
                                | [2]               +------->| * Post-processing |
                                |                   |        | * Modeling        |
                                |                   |        | * Sequencing      |
                        [5a]    |                   +------->| * Unifying        | [6a]
[server log data]-------------->| /api/v1/log/      |  [5b]  |                   |------>+--------+
                                +-------------------+        +------------+------+       | Scylla |
                                                                          |              |        | * this small box
                                                                          |              +----^---+   is deceptively
                                                                          | [6b]              |       complex.
                                                              Scuttlebutt |                   |
        +===================+                                +------------V------+            |
        = Analytics Report  =             [7]                |                   |            | [6b]
        =                   = -----------------------------> | * Real time       +------------+
        = ...               =                                | * Composit model  |
        =                   =                                |                   |
        = ...               =                                |                   |
        =                   =             [8]                |                   |
        = ...               = <----------------------------> |                   |
        =                   =                                |                   |
        +===================+                                +-------------------+


[1] - Using GET or Web Beacon API, send UTM payload to Beacon's API.

[2] - Beacon immediately responds (if GET) with HTTP 204, or nothing if using
      Web Beacon API.

[3] - Beacon generalizes, hashes and pre-anonymizes data and posts to Katric.

[4] - Katric performs further validation and post-processing before emitting to
      subscribers and Scylla the data record.

[5] - Server logs (nginx, apache) submitted to beacon are also sent to Katric,
      which performs unification to enrich existing data.

[6] - A composit model of analytics data is logged to Scylla and Scuttlebutt
      subscribers.

[7] - Requests for analytic reports on sites uses a Scuttlebutt instance to
      subscribe to composit events which update the report model(s), and
      generate an initial model to work with;

[8] - Web socket channel established to handle streaming updates from
      Scuttlebutt.

## level 2 ##

A little more thought:

* Steps 1, 5a can make use of Kong to handle api authorization against configured
  sites, so we don't have garbage requests from sites.

* Steps 3, 5b can be separated with message queue. (Redis?)

* Steps 6a, 6b, 7, 8, don't need as much care in that presumably they would have
  orders of magnitude less traffic from clients which are reading data for
  analytic reports.  Typical proxy/load balancing and Scuttlebutt docker instances
  should cover requests for reports and scale as needed.
