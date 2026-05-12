# A9R

A fast terminal UI for managing AWS resources.

## Features

- EC2 instance viewer
- Live search filtering
- Auto refresh
- Vim-style navigation
- AWS profile support
- Multi-region support
- Responsive terminal UI

## Preview

(screenshot later)

## Installation

### Clone

git clone ...

### Run

go run .

## Requirements

- Go 1.24+
- AWS credentials configured

## AWS Config Example

~/.aws/config

## Controls

| Key | Action |
|-----|--------|
| TAB | switch focus |
| j/k | move |
| / | search |
| r | refresh |
| a | auto refresh |
| q | quit |

## Architecture

configs/
models/
services/
ui/

## Future Roadmap

- S3
- IAM
- VPC
- EKS
- ASG
- CloudWatch
- Fuzzy search

## Tech Stack

- Go
- tview
- AWS SDK v2

## License

MIT