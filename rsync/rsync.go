package rsync

import (
	"io"
	"os"
	"os/exec"
)

// Rsync is wrapper under rsync
type Rsync struct {
	Source      string
	Destination string

	cmd *exec.Cmd
}

// Options for rsync
type Options struct {
	// Verbose increase verbosity
	Verbose bool
	// Quet suppress non-error messages
	Quiet bool
	// Checksum skip based on checksum, not mod-time & size
	Checksum bool
	// Archve is archive mode; equals -rlptgoD (no -H,-A,-X)
	Archive bool
	// Recurse into directories
	Recursive bool
	// Relative option to use relative path names
	Relative bool
	// NoImliedDirs don't send implied dirs with --relative
	NoImpliedDirs bool
	// Update skip files that are newer on the receiver
	Update bool
	// Inplace update destination files in-place
	Inplace bool
	// Append data onto shorter files
	Append bool
	// AppendVerify --append w/old data in file checksum
	AppendVerify bool
	// Dirs transfer directories without recursing
	Dirs bool
	// Links copy symlinks as symlinks
	Links bool
	// CopyLinks transform symlink into referent file/dir
	CopyLinks bool
	// CopyUnsafeLinks only "unsafe" symlinks are transformed
	CopyUnsafeLinks bool
	// SafeLinks ignore symlinks that point outside the tree
	SafeLinks bool
	// CopyDirLinks transform symlink to dir into referent dir
	CopyDirLinks bool
	// KeepDirLinks treat symlinked dir on receiver as dir
	KeepDirLinks bool
	// HardLinks preserve hard links
	HardLinks bool
	// Perms preserve permissions
	Perms bool
	// Executability preserve executability
	Executability bool
	// CHMOD affect file and/or directory permissions
	CHMOD os.FileMode
	// Acls preserve ACLs (implies -p)
	ACLs bool
	// XAttrs preserve extended attributes
	XAttrs bool
	// Owner preserve owner (super-user only)
	Owner         bool
	Compress      bool
	CompressLevel int
}

// StdoutPipe returns a pipe that will be connected to the command's
// standard output when the command starts.
func (r Rsync) StdoutPipe() (io.ReadCloser, error) {
	return r.cmd.StdoutPipe()
}

// StderrPipe returns a pipe that will be connected to the command's
// standard error when the command starts.
func (r Rsync) StderrPipe() (io.ReadCloser, error) {
	return r.cmd.StderrPipe()
}

// Run start rsync task
func (r Rsync) Run() error {
	return r.cmd.Run()
}

// New returns task with described options
func New(source, destination string, options Options) *Rsync {
	arguments := append(getArguments(options), source, destination)
	return &Rsync{
		Source:      source,
		Destination: destination,
		cmd:         exec.Command("rsync", arguments...),
	}
}

func getArguments(options Options) []string {
	arguments := []string{}
	if options.Verbose {
		arguments = append(arguments, "--verbose")
	}

	if options.Checksum {
		arguments = append(arguments, "--checksum")
	}

	if options.Quiet {
		arguments = append(arguments, "--quiet")
	}

	if options.Archive {
		arguments = append(arguments, "--archive")
	}

	if options.Recursive {
		arguments = append(arguments, "--recursive")
	}

	if options.Relative {
		arguments = append(arguments, "--relative")
	}

	if options.NoImpliedDirs {
		arguments = append(arguments, "--no-implied-dirs")
	}

	if options.Update {
		arguments = append(arguments, "--update")
	}

	if options.Inplace {
		arguments = append(arguments, "--inplace")
	}

	if options.Append {
		arguments = append(arguments, "--append")
	}

	return arguments
}
