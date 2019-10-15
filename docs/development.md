# development #

Where we try to reach dev-prod parity.

## Sandboxing ##

Soon(tm) the vagrant provisioner and Dockerfile(s) will be available, so you
may for example:

  $ vagrant up

And then have an entire ScyllaDB backed system available with:

  http://localhost:8003/ - for tracking,
  http://localhost:8002/ - for reporting.

## Building ##

  $ make

## Running ##

  $ make run

  or specifically;

  $ cd bin
  $ ./vulcan -h

You may then locally run beacon, ellipsis, katric and scuttlebutt.  To generate
some analytics and test the system, first go to:

  http://localhost:8003/


## testing ##

  $ make test
