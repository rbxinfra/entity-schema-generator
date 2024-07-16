# Entity Schema Generator

Command line tool for parsing Roblox Entity files and converting them to .NET entity files and MSSQL database migrations.

## Compatibility

Right now this supports the following:

- [ ] Business Logic Layer
	- [x] Business Logic Layer v1
    - [ ] Entities as a Service (EaaS)
    - [ ] Roblox Entity Framework Core (REFC) -- Currently not possible due to no BLL existing
- [ ] Data Logic Layer
    - [x] Data Logic Layer v1 (Database Access)
    - [x] Data Logic Layer v2 (Mssql + Data Access Patterns)
    - [ ] Data Logic Patterns (EaaS)
    - [ ] Roblox Entity Framework Core (REFC)
- [ ] Platform Entity Patterns
    - [x] Platform Entity interface and implemention (Roblox.Entities)
    - [ ] Platform Entity factory interfaces and implementations (BLL Proxy)
- [x] MSSQL
    - [x] Database Generation with Sharding
    - [x] Support for all methods across DALv1 and DALv2
- [ ] CockroachDB
    - [ ] CDB SQL Generation

## Building

Ensure you have [Go 1.20.3+](https://go.dev/dl/)

1. Clone the repository via `git`:

    ```txt
    git clone git@github.com:rbxinfra/entity-schema-generator.git
    cd entity-schema-generator
    ```

2. Build via [make](https://www.gnu.org/software/make/)

    ```txt
    make build-debug
    ```

## Usage

`cd src && go run main.go --help` (use the build binary found in the bin directory if you downloaded a prebuilt or built it yourself)

```txt
Usage: entity-schema-generator
Build Mode: debug
        [-h|--help]
        [--configuration-directory <directory>] [--output-directory <directory>]
        [--recurse]

  -configuration-directory string
        The directory where the configuration files are located.
  -help
        Print usage.
  -output-directory string
        The directory where the output files will be located. (default "./out")
  -recurse
        Recurse into subdirectories. (default true)
```
Example: 
entity-schema-generator --configuration-directory ./configurations --recurse
