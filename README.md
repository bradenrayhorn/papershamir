# archive notice

This project is superseded by https://github.com/bradenrayhorn/paper-backup

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

## Security

papershamir has not been audited.

Anyone using papershamir should consider their individual thread model and use at their own risk.

## Usage

*Single downloadable HTML file is currently in progress*

papershamir has a few features to make it easier to store the secret on paper.

First, a special character set has been chosen to prevent confusion.
For example, "B" has been removed to prevent confusion with "8".

Second, there is a checksum on the end of each line. When combining the parts, papershamir
will provide warnings if the checksum does not match. This helps to narrow down where a typo
might be.

