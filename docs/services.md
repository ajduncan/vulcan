# services #

The Vulcan platform provides the following services, which scale and run in
kubernetes pods - see [architecture](architecture.md) - or the Vulcan can be
run as a standalone app.

These services represent respective areas of concern: gathering analytics,
storing and retrieving records, generating and displaying reports, and providing
documentation with examples.

## beacon ##

  - Tracker endpoint which handles the request payload.
  - Uses HTTP 204 No Content instead of a 1x1 transparent gif (Saving 35? bytes)
  - Uses Web Beacon API with no response expected
  - Runs against 127.0.0.1:8000 as vulcan beacon
  - Minimally processes data and queues for katric to process.

## katric ##

  - Uses redis and taskq to process event jobs submitted by beacon
  - Handles database storage operations between gocql and scylla.
  - Runs against 127.0.0.1:8001 as vulcan katric

## scuttlebutt ##

  - Client service to handle reporting.
  - Handles database retrieval operations between gocql and scylla.
  - a report (often malicious) about the behavior of other people.
  - Runs against 127.0.0.1:8002 as vulcan scuttlebutt

## ellipsis ##

  - Server with example index.html that includes a tracker.
  - (hopefully) includes extensive test client runs in excess of 10k requests,
    with different payloads and tests.
  - Runs against 127.0.0.1:8003 as vulcan ellipsis
