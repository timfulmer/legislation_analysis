# @LegiBot

@LegiBot is a Twitter chatbot with an understanding of legislation introduced during the 116th US congress.  Chat with me on Twitter!

## Methodology

@LegiBot performs NLP and statistical analysis on the full archive of congressional legislation introduced by the 116th US congress available at https://govinfo.gov:

1. Plain text is extracted from the XML files made available online.
1. Plain text is tokenized and sanitized.
1. Statistical occurrence rates are calculated for interesting tokens.

Full implementation is found in `analyze.go`.

## Architecture

@LegiBot is built using GCP + go, and is designed to scale to any amount of traffic without requiring work on the back end.  One trade off was made for cost, by using Google Cloud SQL instead of Cloud Bigtable.  Known issues using a relational database (tech debt) are mitigated with the following:

- Runtime monitoring and alerting is in place around the Google Cloud SQL instance.
- Can re-generate the database given about 10min of downtime, no backups are required.

## Roadmap

1. Improve analysis to remove more noise from the data set.
1. Include Bill name and sponsor information in the data set.
1. Include Congressional map data, to relate user location to bill sponsors.
1. Show data related user's elected congresspeople in chat conversation.
1. Dynamically pull congressional legislation database from https://govinfo.gov.

## Setup

1. Install go:

    ```
    brew install go
    ```

1. Install Google Cloud SDK:

    ```
    brew tap caskroom/cask
    brew cask install google-cloud-sdk
    ```

1. Add variables to environment:

    ```
    export GOPATH="${HOME}/.go"
    export GOROOT="$(brew --prefix golang)/libexec"
    export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin
    source '/usr/local/Caskroom/google-cloud-sdk/latest/google-cloud-sdk/path.bash.inc'
    source '/usr/local/Caskroom/google-cloud-sdk/latest/google-cloud-sdk/completion.bash.inc'
    ```

1. Setup local App Engine emulation, enter `Y` to install: 

    ```
    dev_appserver.py
    ```

1. Install this repository to your local Go project:

    ```
    go get github.com/timfulmer/legislation_analysis
    ```

1. Install App Engine dependencies:

    ```
    go get -u google.golang.org/appengine/...
    ```

1. Run the local App Engine emulator to install the supporting packages, enter `Y` to install:

    ```
    dev_appserver.py ./app.yaml
    ```

1. Install GoLand for IDE
1. Open the project in GoLand and happy coding!
