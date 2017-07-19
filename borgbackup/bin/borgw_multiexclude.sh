#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

declare -r srcDir=/

printf -v sedExpr -- 's/^.*$/%s %s%s%s %s/g' '--exclude' "'" '\&' "'" '\\'
#printf 'sed exp: "%s"\n' "$sedExpr"

# Edit ABOVE this line #######################
##############################################
die() (
  local msg="$1"; shift
  printf 'Error: '"$msg" "$@" >&2; exit 1
)
isWriteDir() ( [[ -d "$1" || -w "$1" ]]; )

this="$(basename "$0")"; declare -r this

[[ "$#" -ge 2 ]] || die '
  Usage: BACKUP_DISKDIR HOME_EXCLUDES[...]

  Such that a borg repo exists at BACKUP_DISKDIR

  NOTE: Lines in each HOME_EXCLUDES are taken to be relative to root. So a line
  with ".ssh/" refers - probably erroneously - to /.ssh/.
' "$this"
declare -r repoDir="$1"; shift

[[ "$EUID" -eq 0 ]] || die 'MUST be run as root, but found $EUID=%s\n' "$EUID"

[[ "${BORG_PASSPHRASE:-x}" != x ]] || die '$BORG_PASSPHRASE not set\n'

isWriteDir "$repoDir" ||
  die 'Machine backups container not writeable dir:\n\t"%s"\n' "$repoDir"

[[ -d "$srcDir" && -r "$srcDir" ]] ||
  die 'Source for backup is not a readable directory:\n\t"%s"\n' "$srcDir"

printf '[%s] %s\tSTARTING...\n' "$(date --iso-8601=s)" "$this"

borgExcludeBackup="$(mktemp --tmpdir 'borgexc-backup_XXXX.sh')"; declare -r borgExcludeBackup
[[ -e "$borgExcludeBackup" && -w "$borgExcludeBackup" ]] || {
  printf 'failed to start building borgbackup script: %s\n' "$borgExcludeBackup" >&2
  exit 1
}

cleanup() (
  printf '[%s] %s\tCleaning up...\n' "$(date --iso-8601=s)" "$this"
  rm -v "$borgExcludeBackup"
  printf '[%s] %s\tEXITING\n' "$(date --iso-8601=s)" "$this"
)
trap cleanup EXIT

chmod +x "$borgExcludeBackup"

cat <<EOF |
  borg create \
    --debug --verbose --stats \
    "$repoDir"::'{hostname}-'"$(date --iso-8601=ns)" \
    "$srcDir"
EOF
tr -d '\n' >> "$borgExcludeBackup"

# exclude paths, taken from https://wiki.archlinux.org/index.php/full_system_backup_with_rsync
declare -a systemExcludes=(
  '/dev/*'
  '/proc/*'
  '/sys/*'
  '/tmp/*'
  '/run/*'
  '/mnt/*'
  '/media/*'
  '/lost+found'
  '/keybase'
  '/var/lib/lxcfs'
)

autoExcludeRepo() (
  local parent; parent="$(dirname "$repoDir")"
  [[ "$parent" = / ]] || {
    printf '%s' "$parent"
    return 0
  }
  printf '%s' "$repoDir"
)
autoExcluded="$(autoExcludeRepo)"
printf 'auto-excluding repo via exclude: "%s"\n' "$autoExcluded"
systemExcludes+="$autoExcluded"

for exclude in "${systemExcludes[@]}";do
  printf -- " --exclude '%s' " "$exclude" >> "$borgExcludeBackup"
done

printf 'processing excludes files...\n'
for excFile in "$@";do
  printf '\tprocessing file: %s\n' "$excFile"
  awk \
    '{ printf "  " "--exclude '"'"'" $1 "'"'"'" "  " }' \
    < "$excFile" \
    >> "$borgExcludeBackup"
done
printf 'done processing excludes files\n'

printf 'Dynamic script built; going to run the below:\n'
{ cat -n "$borgExcludeBackup"; echo; }

printf 'Running...\n'
(
  set -x
  time "$borgExcludeBackup"
)
