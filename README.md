# vulcan #

🖖 Web analytics. 🖖

## motivation ##

An explorative way of learning how web analytics works at scale.

* Vulcan aims to:

  - Be fast.
  - Handle many requests.
  - Have strong opinions.
  - Be a compelling choice for individuals, organizations and companies.
  - Be self-hosted, open source, easily extended.
  - Follow best practices.

* Area of focus:

  - Provide a beacon API service for collecting data.
  - Store transformed data as collected information.
  - Simple (not feature rich) analytics reporting out of the box.

* Strategy:

  - To scale, use sharding around collection endpoints where possible, so;
  - Kong can be used as a front between APIs;
  - Separation between services which isolate:
      a. HTTP GET/web beacon API requests for analytics,
      b. storage and retrieval of site specific data
      c. presentation of retrieved analytics

Further [rationale](docs/rationale.md) provided.

## services ##

The project is broken down into four major component services which have their
respective areas of concern: gathering analytics, storing and retrieving
records, generating and displaying reports, and providing documentation with
examples;

1. beacon

  - Tracker endpoint which handles the request payload.
  - Uses HTTP 204 No Content instead of a 1x1 transparent gif (Saving 35? bytes)
  - Uses Web Beacon API with no response expected
  - Runs against 127.0.0.1:8000 as vulcan beacon

2. katric

  - Handles database storage and retrieval operations between gocql and scylla.
  - Runs against 127.0.0.1:8001 as vulcan katric

3. scuttlebutt

  - Client service to handle reporting.
  - Works with a local copy of sharded ark data from katric.
  - a report (often malicious) about the behavior of other people.
  - Runs against 127.0.0.1:8002 as vulcan scuttlebutt

4. ellipsis

  - Server with example index.html that includes a tracker.
  - (hopefully) includes extensive test client runs in excess of 10k requests,
    with different payloads and tests.
  - Runs against 127.0.0.1:8003 as vulcan ellipsis

## License ##

MIT - See [LICENSE.md](license.md)

## Development ##

Soon(tm) the vagrant provisioner and Dockerfile(s) will be available, so you
may for example:

  $ vagrant up

And then have an entire ScyllaDB backed system available with:

  http://localhost:8003/ - for tracking,
  http://localhost:8002/ - for reporting.

for now:

### building ###

  $ make

### running ###

  $ make run

  or specifically;

  $ cd bin
  $ ./vulcan -h

You may then locally run beacon, ellipsis, katric and scuttlebutt.  To generate
some analytics and test the system, first go to:

  http://localhost:8003/


### testing ###

  $ make test

## contributing ##

Please review [standards](docs/standards.md) before submitting issues and pull
requests.  Thank you in advance for feedback, criticism, and feature requests.
