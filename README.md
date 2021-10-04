# goload-manager

## About
A terminal-based download manager written in Go!
This project uses [grab](github.com/cavaliercoder/grab) for downloading and [tcell](github.com/gdamore/tcell) for the terminal interface.

Heavily work-in-progress.

## F.A.Q.

### Why is this written in Go?
I could ramble about how fast Go is and how well it does concurrency but the truth is I just wanted to learn it.

### Why a download manager? Aren't there tons of them already?
I thought it would be an interesting challenge to make in Go in my learning journey. It can be expanded to incorporate various cool features that will help me learn more about uses of the language.

### Why terminal-based? If I want to download something from the terminal I can just use `curl` or `wget`.
I always wanted to make a terminal application with a pretty user interface (as pretty as a terminal can be). These tools (`curl` and `wget`) can't cover the situation where you want to download multiple files and want to monitor the progress, pause and continue them from a terminal in an simple an sightly manner.

### When will the be done? I wanna use it already!
Sorry, this is a learning project that I'm working on in parallel with my university studies and everything else I have going on in my life. So I can't give a time schedule.

### I want to contribute
I want to work on the major features of this myself; however If you have an interesting idea in mind, notice a bug, have a refactoring proposal etc you are welcome to open an issue.

## Future features
 - [ ] Ability to add new links to download
 - [ ] Ability to pause and resume a download
 - [ ] Sort by columns
 - [ ] Remote! manage downloads on a different machine
 - [ ] GUI using Qt

 ##### This is made in-part as a learning project, that's why the above features are so ambitious D:
