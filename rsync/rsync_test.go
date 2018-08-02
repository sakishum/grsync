package rsync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArguments(t *testing.T) {
	t.Run("--verbose", func(t *testing.T) {
		args := getArguments(Options{
			Verbose: true,
		})
		assert.Contains(t, args, "--verbose")
	})

	t.Run("--checksum", func(t *testing.T) {
		args := getArguments(Options{
			Checksum: true,
		})
		assert.Contains(t, args, "--checksum")
	})

	t.Run("--quiet", func(t *testing.T) {
		args := getArguments(Options{
			Quiet: true,
		})
		assert.Contains(t, args, "--quiet")
	})

	t.Run("--archive", func(t *testing.T) {
		args := getArguments(Options{
			Archive: true,
		})
		assert.Contains(t, args, "--archive")
	})

	t.Run("--recursive", func(t *testing.T) {
		args := getArguments(Options{
			Recursive: true,
		})
		assert.Contains(t, args, "--recursive")
	})

	t.Run("--relative", func(t *testing.T) {
		args := getArguments(Options{
			Relative: true,
		})
		assert.Contains(t, args, "--relative")
	})

	t.Run("--no-implied-dirs", func(t *testing.T) {
		args := getArguments(Options{
			NoImpliedDirs: true,
		})
		assert.Contains(t, args, "--no-implied-dirs")
	})

	t.Run("--update", func(t *testing.T) {
		args := getArguments(Options{
			Update: true,
		})
		assert.Contains(t, args, "--update")
	})

	t.Run("--inplace", func(t *testing.T) {
		args := getArguments(Options{
			Inplace: true,
		})
		assert.Contains(t, args, "--inplace")
	})

	t.Run("--append", func(t *testing.T) {
		args := getArguments(Options{
			Append: true,
		})
		assert.Contains(t, args, "--append")
	})

	t.Run("--append-verify", func(t *testing.T) {
		args := getArguments(Options{
			AppendVerify: true,
		})
		assert.Contains(t, args, "--append-verify")
	})

	t.Run("--dirs", func(t *testing.T) {
		args := getArguments(Options{
			Dirs: true,
		})
		assert.Contains(t, args, "--dirs")
	})

	t.Run("--links", func(t *testing.T) {
		args := getArguments(Options{
			Links: true,
		})
		assert.Contains(t, args, "--links")
	})

	t.Run("--copy-links", func(t *testing.T) {
		args := getArguments(Options{
			CopyLinks: true,
		})
		assert.Contains(t, args, "--copy-links")
	})

	t.Run("--copy-unsafe-links", func(t *testing.T) {
		args := getArguments(Options{
			CopyUnsafeLinks: true,
		})
		assert.Contains(t, args, "--copy-unsafe-links")
	})

	t.Run("--safe-links", func(t *testing.T) {
		args := getArguments(Options{
			SafeLinks: true,
		})
		assert.Contains(t, args, "--safe-links")
	})

	t.Run("--copy-dir-links", func(t *testing.T) {
		args := getArguments(Options{
			CopyDirLinks: true,
		})
		assert.Contains(t, args, "--copy-dir-links")
	})

	t.Run("--keep-dir-links", func(t *testing.T) {
		args := getArguments(Options{
			KeepDirLinks: true,
		})
		assert.Contains(t, args, "--keep-dir-links")
	})

	t.Run("--hard-links", func(t *testing.T) {
		args := getArguments(Options{
			HardLinks: true,
		})
		assert.Contains(t, args, "--hard-links")
	})

	t.Run("--perms", func(t *testing.T) {
		args := getArguments(Options{
			Perms: true,
		})
		assert.Contains(t, args, "--perms")
	})

	t.Run("--executability", func(t *testing.T) {
		args := getArguments(Options{
			Executability: true,
		})
		assert.Contains(t, args, "--executability")
	})

	t.Run("--acls", func(t *testing.T) {
		args := getArguments(Options{
			ACLs: true,
		})
		assert.Contains(t, args, "--acls")
	})

	t.Run("--xattrs", func(t *testing.T) {
		args := getArguments(Options{
			XAttrs: true,
		})
		assert.Contains(t, args, "--xattrs")
	})

	t.Run("--owner", func(t *testing.T) {
		args := getArguments(Options{
			Owner: true,
		})
		assert.Contains(t, args, "--owner")
	})

	t.Run("--group", func(t *testing.T) {
		args := getArguments(Options{
			Group: true,
		})
		assert.Contains(t, args, "--group")
	})

	t.Run("--devices", func(t *testing.T) {
		args := getArguments(Options{
			Devices: true,
		})
		assert.Contains(t, args, "--devices")
	})

	t.Run("--specials", func(t *testing.T) {
		args := getArguments(Options{
			Specials: true,
		})
		assert.Contains(t, args, "--specials")
	})

	t.Run("--times", func(t *testing.T) {
		args := getArguments(Options{
			Times: true,
		})
		assert.Contains(t, args, "--times")
	})

	t.Run("--omit-dir-times", func(t *testing.T) {
		args := getArguments(Options{
			OmitDirTimes: true,
		})
		assert.Contains(t, args, "--omit-dir-times")
	})

	t.Run("--super", func(t *testing.T) {
		args := getArguments(Options{
			Super: true,
		})
		assert.Contains(t, args, "--super")
	})

	t.Run("--fake-super", func(t *testing.T) {
		args := getArguments(Options{
			FakeSuper: true,
		})
		assert.Contains(t, args, "--fake-super")
	})

	t.Run("--sparse", func(t *testing.T) {
		args := getArguments(Options{
			Sparse: true,
		})
		assert.Contains(t, args, "--sparse")
	})

	t.Run("--dry-run", func(t *testing.T) {
		args := getArguments(Options{
			DryRun: true,
		})
		assert.Contains(t, args, "--dry-run")
	})

	t.Run("--while-file", func(t *testing.T) {
		args := getArguments(Options{
			WhileFile: true,
		})
		assert.Contains(t, args, "--while-file")
	})

	t.Run("--one-file-system", func(t *testing.T) {
		args := getArguments(Options{
			OneFileSystem: true,
		})
		assert.Contains(t, args, "--one-file-system")
	})

	t.Run("--block-size", func(t *testing.T) {
		args := getArguments(Options{
			BlockSize: 1,
		})
		assert.ElementsMatch(t, args, []string{"--block-size", "1"})
	})

	t.Run("--rsh", func(t *testing.T) {
		args := getArguments(Options{
			Rsh: "test",
		})
		assert.Contains(t, args, "--rsh", "test")
	})

	t.Run("--rsync-programm", func(t *testing.T) {
		args := getArguments(Options{
			RsyncProgramm: "test",
		})
		assert.Contains(t, args, "--rsync-programm", "test")
	})

	t.Run("--existing", func(t *testing.T) {
		args := getArguments(Options{
			Existing: true,
		})
		assert.Contains(t, args, "--existing")
	})

	t.Run("--ignore-existing", func(t *testing.T) {
		args := getArguments(Options{
			IgnoreExisting: true,
		})
		assert.Contains(t, args, "--ignore-existing")
	})

	t.Run("--remove-source-files", func(t *testing.T) {
		args := getArguments(Options{
			RemoveSourceFiles: true,
		})
		assert.Contains(t, args, "--remove-source-files")
	})

	t.Run("", func(t *testing.T) {
		args := getArguments(Options{
			Delete: true,
		})
		assert.Contains(t, args, "--delete")
	})

	t.Run("--delete", func(t *testing.T) {
		args := getArguments(Options{
			DeleteBefore: true,
		})
		assert.Contains(t, args, "--delete-before")
	})

	t.Run("--delete-during", func(t *testing.T) {
		args := getArguments(Options{
			DeleteDuring: true,
		})
		assert.Contains(t, args, "--delete-during")
	})

	t.Run("--delete-delay", func(t *testing.T) {
		args := getArguments(Options{
			DeleteDelay: true,
		})
		assert.Contains(t, args, "--delete-delay")
	})

	t.Run("--delete-after", func(t *testing.T) {
		args := getArguments(Options{
			DeleteAfter: true,
		})
		assert.Contains(t, args, "--delete-after")
	})

	t.Run("--delete-excluded", func(t *testing.T) {
		args := getArguments(Options{
			DeleteExcluded: true,
		})
		assert.Contains(t, args, "--delete-excluded")
	})

	t.Run("--ignore-errors", func(t *testing.T) {
		args := getArguments(Options{
			IgnoreErrors: true,
		})
		assert.Contains(t, args, "--ignore-errors")
	})

	t.Run("--force", func(t *testing.T) {
		args := getArguments(Options{
			Force: true,
		})
		assert.Contains(t, args, "--force")
	})

	t.Run("--max-delete", func(t *testing.T) {
		args := getArguments(Options{
			MaxDelete: 1,
		})
		assert.ElementsMatch(t, args, []string{"--max-delete", "1"})
	})

	t.Run("--max-size", func(t *testing.T) {
		args := getArguments(Options{
			MaxSize: 1,
		})
		assert.ElementsMatch(t, args, []string{"--max-size", "1"})
	})

	t.Run("--min-size", func(t *testing.T) {
		args := getArguments(Options{
			MinSize: 1,
		})
		assert.ElementsMatch(t, args, []string{"--min-size", "1"})
	})

	t.Run("--partial", func(t *testing.T) {
		args := getArguments(Options{
			Partial: true,
		})
		assert.Contains(t, args, "--partial")
	})

	t.Run("--partial-dir", func(t *testing.T) {
		args := getArguments(Options{
			PartialDir: "test",
		})
		assert.ElementsMatch(t, args, []string{"--partial-dir", "test"})
	})

	t.Run("--delay-updates", func(t *testing.T) {
		args := getArguments(Options{
			DelayUpdates: true,
		})
		assert.Contains(t, args, "--delay-updates")
	})

	t.Run("--prune-empty-dirs", func(t *testing.T) {
		args := getArguments(Options{
			PruneEmptyDirs: true,
		})
		assert.Contains(t, args, "--prune-empty-dirs")
	})

	t.Run("--numeric-ids", func(t *testing.T) {
		args := getArguments(Options{
			NumericIDs: true,
		})
		assert.Contains(t, args, "--numeric-ids")
	})

	t.Run("--timeout", func(t *testing.T) {
		args := getArguments(Options{
			Timeout: 100,
		})
		assert.ElementsMatch(t, args, []string{"--timeout", "100"})
	})

	t.Run("--timeout", func(t *testing.T) {
		args := getArguments(Options{
			Contimeout: 100,
		})
		assert.ElementsMatch(t, args, []string{"--contimeout", "100"})
	})

	t.Run("--ignore-times", func(t *testing.T) {
		args := getArguments(Options{
			IgnoreTimes: true,
		})
		assert.Contains(t, args, "--ignore-times")
	})

	t.Run("--size-only", func(t *testing.T) {
		args := getArguments(Options{
			SizeOnly: true,
		})
		assert.Contains(t, args, "--size-only")
	})

	t.Run("--modify-window", func(t *testing.T) {
		args := getArguments(Options{
			ModifyWindow: true,
		})
		assert.Contains(t, args, "--modify-window")
	})

	t.Run("--temp-dir", func(t *testing.T) {
		args := getArguments(Options{
			TempDir: "test",
		})
		assert.Contains(t, args, "--temp-dir", "test")
	})

	t.Run("--fuzzy", func(t *testing.T) {
		args := getArguments(Options{
			Fuzzy: true,
		})
		assert.Contains(t, args, "--fuzzy")
	})

	t.Run("--compare-dest", func(t *testing.T) {
		args := getArguments(Options{
			CompareDest: "test",
		})
		assert.Contains(t, args, "--compare-dest", "test")
	})

	t.Run("--copy-dest=", func(t *testing.T) {
		args := getArguments(Options{
			CopyDest: "test",
		})
		assert.Contains(t, args, "--copy-dest", "test")
	})

	t.Run("--link-dest", func(t *testing.T) {
		args := getArguments(Options{
			LinkDest: "test",
		})
		assert.Contains(t, args, "--link-dest", "test")
	})

	t.Run("--compress", func(t *testing.T) {
		args := getArguments(Options{
			Compress: true,
		})
		assert.Contains(t, args, "--compress")
	})

	t.Run("--compress-level", func(t *testing.T) {
		args := getArguments(Options{
			CompressLevel: 3,
		})
		assert.ElementsMatch(t, args, []string{"--compress-level", "3"})
	})

	t.Run("--skip-compress=", func(t *testing.T) {
		args := getArguments(Options{
			SkipCompress: []string{"1", "2"},
		})
		assert.ElementsMatch(t, args, []string{"--skip-compress", "1,2"})
	})

	t.Run("--cvs-exclude", func(t *testing.T) {
		args := getArguments(Options{
			CVSExclude: true,
		})
		assert.Contains(t, args, "--cvs-exclude")
	})

	t.Run("--stats", func(t *testing.T) {
		args := getArguments(Options{
			Stats: true,
		})
		assert.Contains(t, args, "--stats")
	})

	t.Run("--human-readable", func(t *testing.T) {
		args := getArguments(Options{
			HumanReadable: true,
		})
		assert.Contains(t, args, "--human-readable")
	})

	t.Run("--progress", func(t *testing.T) {
		args := getArguments(Options{
			Progress: true,
		})
		assert.Contains(t, args, "--progress")
	})

	t.Run("--ipv4", func(t *testing.T) {
		args := getArguments(Options{
			IPv4: true,
		})
		assert.Contains(t, args, "--ipv4")
	})

	t.Run("--ipv6", func(t *testing.T) {
		args := getArguments(Options{
			IPv6: true,
		})
		assert.Contains(t, args, "--ipv6")
	})
}
