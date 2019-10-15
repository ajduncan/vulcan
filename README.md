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
  - Ambassador can be used as a front between APIs (e.g. https://vulcantracker.biz/
    and https://report.vulcantracker.biz/);
  - Separation between services which isolate:
      a. HTTP GET/web beacon API requests and log data for queing analytic data,
      b. workers which transform, validate and store data,
      c. retrieval and presentation of composit analytics

Further [rationale](docs/rationale.md) provided.

To see how Vulcan works, see the [architecture](docs/architecture.md).
To see what services Vulcan provides, see the [services](docs/services.md).

## license ##

MIT - See [LICENSE.md](license.md)

## development ##

To run locally and develop, see [development.md](docs/development.md)

## contributing ##

Please review [standards](docs/standards.md) before submitting issues and pull
requests.  Thank you in advance for feedback, criticism, and feature requests.
