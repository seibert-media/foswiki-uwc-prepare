# foswiki-uwc-prepare

foswiki-uwc-prepare is a command line tool to prepare a [Foswiki](http://foswiki.org/) installation for a migration to
[Confluence](https://www.atlassian.com/software/confluence) with the
[Universal Wiki Converter](https://migrations.atlassian.net/wiki) (UWC).

Unfortunately the development of UWC was discontinued. Therefore the following preparation steps are required to migrate
a current version of Foswiki to Confluence with UWC:

* Replace ``<verbatim>...</verbatim>`` to ``%CODE%...%ENDCODE%`` tags
* Convert base64 inline images to image files in Foswiki's /pub folder

foswiki-uwc-prepare takes over these steps on all Foswiki pages in the /data folder.

## Installation

1. Install Go: https://golang.org/doc/install
1. Get foswiki-uwc-prepare:

    ``go get github.com/seibert-media/foswiki-uwc-prepare``

## Usage

    foswiki-uwc-prepare -dir /path/to/foswiki