/EXCEL_LIKE_DB
    /cmd
        /my-app
            main.go
    /pkg
        /my-package
    /internal
        /my-internal-code
    /api
    /web
    /configs
    /scripts
    /test
    /docs
    README.md
    LICENSE
    go.mod
    go.sum

Contact: https://github.com/Shintaro-run

This project includes cute icons throughout its documentation and structure.

Basic directory structure:
/cmd: Contains the entry points of the programs. The main application file (main.go) is located here.
/pkg: Contains reusable package code. This includes library code that can be imported by other projects.
/internal: Contains private application and library code. The code in this directory is not intended to be imported by external projects.
/api: Contains external API definitions (e.g., Swagger specifications, protocol buffer files).
/web: Contains web application-specific components (e.g., frontend files).
/configs: Contains configuration files.
/scripts: Contains scripts for building, installing, analyzing, etc.
/test: Contains additional external test applications and test data.
/docs: Contains project documentation (e.g., design documents, user guides).

Common files:
README.md: Describes the project overview, build instructions, and usage.
LICENSE: Contains the project's license information.
go.mod: Manages module dependencies (used with Go Modules).
go.sum: Contains checksums for ensuring specific versions of dependencies listed in go.mod.
# EXCEL_LIKE_DB
