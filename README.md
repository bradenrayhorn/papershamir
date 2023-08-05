# papershamir

A CLI for securely splitting a secret into parts in a format easily writable on paper using Shamir's Secret Sharing.
Inspired by David Shaw's [Paperkey](https://www.jabberwocky.com/software/paperkey/) project.

This project uses the [hashicorp/vault](https://github.com/hashicorp/vault) Shamir's Secret Sharing algorithm.

Shamir's Secret Sharing allows a "secret" to be split up into multiple parts. The "secret" can only be revealed
when a number of parts greater than or equal to the defined threshold are combined.

Read more on [Wikipedia](https://wikipedia.org/wiki/Shamir's_secret_sharing).

A use case for this application is storing a physical copy of a "master password" (such as for a password manager).
The parts can be distributed in different physical locations. This method of storing a secret
can reduce the risk of losing the information due to natural disasters, misplaced paper, etc. while also
reducing the risk of theft due to the fact multiple parts are required to reassemble the secret.

## Installation

```bash
go install github.com/bradenrayhorn/papershamir/cmd/papershamir@latest
```

## Usage

papershamir has a few features to make it easier to store the secret on paper.

First, a special character set has been chosen to prevent confusion.
For example, "B" has been removed to prevent confusion with "8".

Second, there is a checksum on the end of each line. When combining the parts, papershamir
will provide warnings if the checksum does not match. This helps to narrow down where a typo
might be.

### Splitting a secret

The split command takes input from stdin.

There are options available to configure the number of parts generated and threshold required to rebuild
the secret:

```
go run papershamir split --help
Usage: papershamir split

Split secret into parts. Pass secret as stdin.

Flags:
  -h, --help           Show context-sensitive help.

      --parts=5        Number of parts to split secret into.
      --threshold=3    Number of parts required to reassemble the secret.
```

Example of using the split command:

```
papershamir split <<< "This is my password."
```

### Rebuilding a secret

It is recommended to enter the secret parts into a file before combining. This
makes editing easier.

The secret parts must have a newline between each other.

Example `secret_file`:

```
HX 16 54 8W 54 31 98 3W X9 1E 3X 4K 84 3E 75 WE1468N3
24 8X 81 8A A6 11                            2HX491E8
                                             HK81HX9N

97 HX 7N 8X XA H7 74 59 5N 4E 22 56 HW XX AW 6N8K681W
72 NW 7H 67 WX 72                            6EA849A8
                                             179HA788

A5 X4 45 78 W9 H8 39 WE 3H N7 K3 68 A7 A8 E1 XK52NXN8
5A A6 K3 WX N9 62                            4778E429
                                             5519E28A
```

The combine command takes input from stdin.

Example of using the combine command:

```
cat secret_file | papershamir combine
```
