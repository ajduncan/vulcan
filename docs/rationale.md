# rationale #

So, why create an analytics platform?  

Solving the problems associated with this endeavor provide an excellent learning
opportunity, to look under the hood of how others are gathering and using data
for everything you do online.  There are open source analytics platforms and
standards for sending and collecting data. The Vulcan platform's design
objective is to be effective for both individuals and small companies, to large
scale businesses with millions of tracked subjects.  

Design choices and decisions are based on the problems to be solved.  The Vulcan
Analytics platform will provide an API and mechanism for web sites and app
services to post data for analytics.  It aims to do this with privacy and the
GDPR in mind, in an accessible way for platforms, clients, and apps with a clear,
flexible, and extensible API.

Some discussion points and questions include:

1. Why Go?
2. Why Scylla?
3. Privacy and the GDPR?

## Why Go? ##

Go has several attractive features and design choices for aspects of this
project.  Go is a good choice for programmer productivity in the context of
cloud technologies and large, maintainable projects.

Some of the selling points include:

  * Default statically linked binary with no external dependencies.
  * Remote package management and documentation (go get).
  * Excellent toolchain.
  * Garbage collection.
  * [Well documented](https://golang.org/doc/effective_go.html).
  * Statically typed.
  * Friendly syntax.
  * Concurrency.
  * Scalability.
  * Compiled.
  * Simple!
  * Fast.

Go is a statically typed language developed at Google by Rob Pike, Robert
Griesemer, and Ken Thompson.  The language is simple, clear and unambiguous,
thus a good choice for a maintainable project.  Large successful companies like
Google and Netflix use Go to serve millions of customers concurrently.  Rob Pike
expressed the desire to keep the language specification simple enough to hold in
a programmer's head by omitting certain features (like inheritance).

## Why Scylla? ##

A performant NoSQL database is required to handle the unstructured information
collected for analytics.  If we imagine the potential scale of a successful
analytics platform, we must consider the lowest cost overhead NoSQL database
that has a high throughput and low latency to support storage and retrieval of
millions of requests.  

Apache Cassandra is a proven NoSQL database management system designed to handle
Big Data across commodity servers with high availability and no single point of
failure.  It is a particular type of NoSQL database which contains rows in the
same table that have unmatched or variable columns.  This type of NoSQL database
is called a [wide column store](https://en.wikipedia.org/wiki/Wide_column_store).

Popular wide column stores include:

  * Google's [Bigtable](https://en.wikipedia.org/wiki/Bigtable)
  * Facebook's [Apache Cassandra](https://en.wikipedia.org/wiki/Apache_Cassandra)
  * The NSA's [Apache Accumulo](https://en.wikipedia.org/wiki/Apache_Accumulo)
  * Powerset's [Apache HBase](https://en.wikipedia.org/wiki/Apache_HBase)
  * Zvent's [Hypertable](https://en.wikipedia.org/wiki/Hypertable)
  * ScyllaDB's [Scylla](https://en.wikipedia.org/wiki/Scylla_(database))

Bigtable is proprietary and not appropriate for this project.  Hypertable is no
longer being developed.  Cassandra is 'enterprise,' written in Java, open source,
well established, used by many big companies, and is the most popular wide
column store.  Accumulo and HBase are similarly appropriate for an enterprise,
open source, well established and written in Java.  ScyllaDB is 'lean', new,
written in C++17, and designed to be compatible with Cassandra.

So of Cassandra, Accumulo, HBase, and ScyllaDB, which is appropriate?  

Jeanine Stark writes (see summary of Cap Theorem below) that;

``However, Cassandra is the fastest database in relation to writes to the
database because of the high level of attention that is spent with respect to
how the data is stored on disk when the database has been properly designed.
Cassandra is therefore the correct choice for a database where a high volume of
writes will take place. One common example is to use Cassandra for logs. Logs
have a high volume of writes so having better performance for writes is ideal.``

Cassandra and ScyllaDB are appropriate for storing, essentially log data.

### Cap Theorem ###

The [cap theorem](https://en.wikipedia.org/wiki/CAP_theorem) states that it is
impossible for a distributed data store to simultaneously provide more than two
out of the following three guarantees:

  1. Consistency: Every read receives the most recent write or an error.
  2. Availability: Every request receives a (non-error) response, without the
     guarantee that it contains the most recent write.
  3. Partition tolerance: The system continues to operate despite an arbitrary
     number of messages being dropped (or delayed) by the network between nodes.

An excellent summary of Cap Theorem and design choices for Big Data NoSQL
database management systems is provided (here)[https://blog.ippon.tech/use-cassandra-mongodb-hbase-accumulo-mysql/]
by Jeannine Stark, technology consulting firm Ippon Accelerator.

### ScyllaDB's Claims ###

From [ScyllaDB's website](https://www.scylladb.com/), ScyllaDB is claimed to be
the fastest nosql database.  Further, they claim:

  1. Scylla is a drop-in Apache Cassandra alternative big data database that
  powers your applications with ultra-low latency and extremely high throughput,
  while reducing TCO to a fraction of most NoSQL databases.

  2. We reimplemented Apache Cassandra from scratch using C++ instead of Java to
  increase the raw performance and utilization of modern multi-core servers and,
  through self-tuning and improved uptime, minimize the overhead to DevOps.
  Scylla provides the NoSQL database platform your applications require to scale
  out and up linearly.

These are extremely attractive and compelling features for a project that wants
to scale from individuals or small business up to a huge enterprise, with
millions of requests a second.

It's also claimed that:

  1. Scylla Cloud uses 1/10 the number of servers at BigTable to achieve the
     same performance.
  2. Scylla has handled *2 million* requests *per second* on a single node:

``With extra performance to work with, NoSQL projects can have more flexibility
to focus on other concerns, such as functionality and time to market. Scylla
enables faster cluster scaling, more overhead to handle complex queries, and the
power to do complex analytics tasks at the same time as routine administration
operations.``

[source](https://www.socallinuxexpo.org/scale/14x/presentations/scylladb-cassandra-compatibility-18-million-requests-node)

  3. Scylla handles a data set ten-times bigger than cassandra with comperable
     or better performance.  See [this report](https://www.scylladb.com/2017/02/15/scylladb-vs-cassandra-performance-benchmark-samsung/)

See all ScyllaDB's benchmarks [here](https://www.scylladb.com/product/benchmarks/)

## Privacy and the GDPR? ##

Based on the GDPR, collected data should not allow one to single-out, link
records through correlation, or infer identity based on a set of values.  In
other words, the data should be sufficiently anonymous (perhaps including random
noise, especially for low-record result sets) to safe guard a user's data
privacy.

Some existing open source analytic projects claim to provide anonymous tracking:

https://news.ycombinator.com/item?id=20498259

  - using hashed data, but doesn't provide fully unique visit (users with the same IP)
  - if everyone used IPv6, this wouldn't be as much of an issue?, but we do

See also:
https://www.eff.org/deeplinks/2018/06/gdpr-and-browser-fingerprinting-how-it-changes-game-sneakiest-web-trackers


## General Notes ##

Decisions should be made regarding anonymization techniques and to work under
the GDPR and to, in general, "never be evil."

### Google ###

It's impossible to discuss analytics without bringing up Google, and how they
purport to protect individuals using [anonymization](https://policies.google.com/technologies/anonymization?hl=en).
The general idea is to ensure that:

  1. Data can't be uniquely identified to an individual using various techniques.
  2. The anonymized data is still valuable for analytics.

These two points are at ends with each other.  For example, an IP address under
GDPR is considered [personal data](https://www.jeffalytics.com/gdpr-ip-addresses-google-analytics/)
and consent is required to store this data.  Generalizing an IP address as Google
and others do removes the accuracy of geo-locating a site visitor.  According to
[this](https://www.conversionworks.co.uk/blog/2018/04/16/ip-anonymization-ga-impact-assessment/)
[study](http://rpubs.com/CW_Huiyan/anonymize_ip_ga_impact) by Conversion Works Ltd,
80% of the visists have less than 50km discrepency.  Continent association with
generalized IP address is as you'd expect, with less than a 1% discrepency.

Data generalization techniques that Google claims to use include:

  1. [k-anonymity](https://en.wikipedia.org/wiki/K-anonymity)
  2. [l-diversity](https://en.wikipedia.org/wiki/L-diversity)

These techniques reduce the de-anonymization attacks:

  1. homogeneity
  2. background knowledge

According to [this](https://www.cs.cmu.edu/~jblocki/Slides/K-Anonymity.pdf),
Any actual privacy guarantee must be proved and established mathematically.  K-anonymity
can be achieved by using suppression and generalization, but is NP-Hard.

### GDPR ###

The General Data Protection Regulation [GDPR](https://en.wikipedia.org/wiki/General_Data_Protection_Regulation)
is a regulation in EU law on data protection and privacy.

See [this article](https://www.eff.org/deeplinks/2018/06/gdpr-and-browser-fingerprinting-how-it-changes-game-sneakiest-web-trackers),
under the section Fingerprinting after the GDPR.

## requests ##

The parameters used by, for example, [Google Analytics](https://developers.google.com/analytics/resources/concepts/gaConceptsTrackingOverview)
are prefixed with utm; which stands for Urchin tracking module.  UTM was
specified by the Urchin Software Corporation and later acquired by Google in
2005.  

### utm url codes ###

UTM codes explain how traffic is coming to your site.  A UTM-tagged url includes

  * utm_source (required): the source of traffic, utm_source=twitter
  * utm_medium: the medium of traffic, utm_medium=email, social, cpc
  * utm_content: the content clicked, utm_content=call_to_action, smalltext, etc.
  * utm_term: paid terms(ga), utm_term=coats
  * utm_campaign: the marketing campaign, utm_campaign=winter_special

and is referenced in URLs with:

?utm_source=source&utm_medium=email&utm_content=small_text&utm_campaign=email-campaign

This allows you to create links that track:

  1. Where is the traffic coming from e.g. utm_source=facebook.
  2. Through what medium is the traffic getting the link e.g. utm_medium=social
  3. What campaign is this associated with, e.g. utm_campaign=facebook-marketing

## disabled javascript? ##

What segment of the market disables javascript?  Does this matter?  Some people
are going to market to the security conscious.  What tools can we reference for,
example, log analysis (and should scuttlebutt allow this?)

- https://goaccess.io/

## beacon ##

The beacon service can handle traditional requests which respond with HTTP 204,
No Content.  The beacon service also handles [web beacon](https://en.wikipedia.org/wiki/Web_beacon)
requests.  Beacon uses the [Beacon API](https://www.sitepoint.com/introduction-beacon-api/),
or [Beacon](https://www.w3.org/TR/beacon/)
