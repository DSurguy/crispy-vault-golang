# Crispy Vault (Golang version)
This repo is a bad idea clone of https://github.com/DSurguy/crispy-vault. I was stuck on correctly interpreting 
a rust error around a sqlite statement. In a blind rage I rewrote the entire backend from rust to golang in one 
day, swapping Tauri for Wails.

I had the same error. See: https://github.com/DSurguy/crispy-vault/commit/019751a4ac2569bbcddae57a68a6387461fbb7a5

This has 100% feature parity with all previous commits. I abandoned it immediately afterward. Fun times!

# Setup
Follow the instructions at https://wails.io/docs/gettingstarted/installation.

This will require you to correctly install go and add it to your path, and then install wails globally.

Make sure you add both the `go` binary AND the go modules binary folder (the one that contains wails), or any `wails` commands will fail. And you will be sad.

Run `wails doctor` and resolve any issues, with the possible exception of `webkit2gtk`

## webkit2gtk weirdness
I run EndeavorOS, a flavor of arch linux. Currently, it only has version 4.1 of webkit2gtk installed, and trying to get 4.0 installed would probably cause a bunch of headaches. Wails docs have a note about this here: https://wails.io/docs/gettingstarted/building.

Based on which version you have installed, you modify the `dev` or `build` commands, don't break your system for this.

# Running the Project

If you have webkit2gtk-4.1: `wails dev -tags "webkit2_41 fts5"`

If not: `wails dev -tags "fts5"`