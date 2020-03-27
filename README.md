# Hotel-Data-Converter
Hotel-Data-Converter is a tool which help coverts hotel data from one format to another.

# Table of Content
* [Supported Formats](#supported-formats)
* [Validations](#validations)
* [Output](#output)
* [Technologies & Tools](#technologies-&-tools)
* [Command Line Flags](#command-line-flags)
* [Setup and Execution](#setup-and-execution)
* [Testing](#testing)
* [Release](#release)
* [Extending](#Extending)
* [Authors](#authors)
* [Acknowledgments](#acknowledgments)
* [License](#license)
* [Notes](#notes)


# Supported Formats
Currently this tool supports:
1. Converting `CSV` data format to `JSON` format.
2. Converting `CSV` data format to `XML` format.


# Validations
These are the validation rules the tool applies to each data format its about to convert
### CSV Data Formats
  * The columns/fields must be up to six
  * The column positions matters to get appropriate results:
    * `First column`: Name
    * `Third column`: Stars
    * `Sixth column`: Uri
  * `Name`: The name is considered valid if it only contains UTF-8 characters
  * `Stars`: The stars are integers ranging from `0` to `5`.
  * `Uri`: A valid uri must be [a full (or absolute) URL](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/What_is_a_URL). The URL protocols `(http or https)`, paths, and parameters `(?key1=value1&key2=value2)` are optional, domain names `(www.github.com)` are required.
  some examples of valid url formats:
    * http://www.github.com
    * https://www.github.com
    * http://github.com
    * www.github.com
    * github.com


# Output
The converted data formats are outputed in the `share/results` folder.
The output files are named according to the timestamp in seconds e.g `timestamp-name-of-file-converted.extension`.
The timestamp is derived from the current time when the tool is ran.
For example converting a `CSV data` to `JSON` and `XML` at timestamp `1582824983` will output the following in the `share/results` folder:
* 1582824983-`<`name-of-file-converted`>`-invalid.csv: this contains csv records that do not pass the [csv validation](###csv-data-formats).
* 1582824983-`<`name-of-file-converted`>`.json
* 1582824983-`<`name-of-file-converted`>`.xml


# Technologies & Tools
* `Mac OSX` - Operating System
* [Go](https://golang.org/) - Run time environment
* [Docker](https://www.docker.com) - Build tool

# Command Line Flags
This flag is only necessary if you're using the `go` commands to `run` or `build` the program.
* `file` - the name of the file that's intended to be converted. Ensure the file is in the `share/files` folder.
  * `required?`: true
  *  `default`: none.
* `sort` - to sort the hotel data by `stars`. if omitted no sort is applied to the result.
  * `allowed values`: `ascend` or `descend`
  * `required?`: false
  *  `default`: none.


# Setup and Execution
  Operating System: `Mac OSX`

## With Go Binaries file
* Ensure you have `Go` [setup](https://golang.org/doc/install) on your local machine
  *  You can also checkout [this go installation and setup tutorial](https://www.digitalocean.com/community/tutorial_series/how-to-install-and-set-up-a-local-programming-environment-for-go)

* Change directory to the root of the project 
  * On your command-line: run `cd path/to/hotel-data-converter`
* On your command-line: run `go build -o bin/main cmd/app/main.go cmd/app/app.go`
* Create the `share/files` folder at the root of the project if it doesn't exist
* Move the file you want to convert to the `share/files` folder located at the root of the project.
* On your command-line: run `bin/main -file=<`name-of-file-to-convert`>`
  * if you want a sorted result: run `bin/main -file=<`name-of-file-to-convert`> -sort=<`[allowed-values](###command-line-flags)`>`
* Check the `share/results` folder for the converted files.

## With Docker
  If you intend to use docker:
  * Ensure you have [Docker setup](https://docs.docker.com/install/#supported-platforms) on your local machine.
  * Ensure you have [Docker Compose setup](https://docs.docker.com/compose/install/#install-compose) on your local machine.
  * Change directory to the root of the project 
    * On your command-line: run `cd path/to/hotel-data-converter` 
  * Create the `share/files` folder at the root of the project if it doesn't exist
  * On your command-line: run `docker-compose build`.
  * Move the file you want to convert to the `share/files` folder located at the root of the project.
  * Create a `.env` file - reference the `.env.sample` file.
    * add the name of the file to `.env` i.e `FILE` = `<`name-of-file-to-convert`>`
  * On your command-line: run `docker-compose up`.
  * Check the `share/results` folder for the converted files.


# Testing
* Change directory to the root of the project 
  * On your command-line: run `cd path/to/hotel-data-converter`
  *  For unit test
      * On your command-line: run  `go test ./tests`


# Release
* Version 1.0.0


# Extending
To support more formats conversion, this tool can be extend. Simply go through the folder structure and architecture to see where extensions can be added.


# Authors
* [**Ogooluwa Akinola**](https://github.com/rovilay)


# Acknowledgments
* [**Ogooluwa Akinola**](https://github.com/rovilay)


# License
*  MIT


# Notes
* The hotel data structure should be in this format:
  * Name
  * Address
  * Stars
  * Contact
  * Phone
  * URI
* Currently you can only sort the hotel data by `Stars`.
