## Installation
Because the latest version of Flyway Community Edition does not support MySQL 5.7, the latest version we can use is `7.14.1`. Please follow the below instructions to install this version of flyway onto your system.

### Linux
Run the following command to download the flyway archive and install the binary:
```
cd ~/ && wget -qO- https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/7.14.1/flyway-commandline-7.14.1-linux-x64.tar.gz | tar xvz && sudo ln -s ~/flyway-7.14.1/flyway /usr/local/bin
```

### Mac
Run the following command to download the flyway archive and install the binary:
```
cd ~/ && curl https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/7.14.1/flyway-commandline-7.14.1-macosx-x64.tar.gz | tar xv && sudo ln -s ~/flyway-7.14.1/flyway /usr/local/bin
```

### Windows
Download the [flyway archive](https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/7.14.1/flyway-commandline-7.14.1-windows-x64.zip) from the maven repository.

Extract the `flyway-7.14.1` folder inside of the archive to `C:\` and simply add the new `C:\flyway-7.14.1` directory to the `PATH` enironment variable to make the `flyway` command available from anywhere on your system.

## Usage
While in the root project directory, you have can use the following Makefile commands:

Command | Description
--- | ---
`make migrate env={ENV}` | Run migrations on specified environment.
`make info env={ENV}` | Get migration information about specified environment.
`make repair env={ENV}` | Repair migrations in a specified environment.

Alternatively, you can use the underlying flyway commands:
```
flyway -configFiles=flyway/{ENV}.conf migrate
flyway -configFiles=flyway/{ENV}.conf info
flyway -configFiles=flyway/{ENV}.conf repair
```

For more options, use `flyway --help` or view the [Flyway documentation](https://flywaydb.org/documentation/command/migrate).