# Backs up an Arch/Manjaro system in all the ways I care about
# NOTE: This assumes certain things are in certain places, so won't work
# for a general system - only mine.
if [ $# -ne 1 ]; then
    echo "USAGE: $0 <path/to/external/harddrive>"
    exit 1
fi

if [ ! -d "$1" ]; then
    echo "$1 is not a directory."
    exit 2
fi

BACKUPDIR="$1"/backup
if [ -f "$BACKUPDIR" ]; then
    echo "$BACKUPDIR already exists."
    exit 3
fi

mkdir "$BACKUPDIR"

# Grab the list of Pacman installed stuff
echo "Grabbing paclist..."
pacman -Q > "$BACKUPDIR/paclist.txt"

# Grab the aur directory
echo "Copying aur directory..."
rsync -ah --progress "/home/$USER/aur" "$BACKUPDIR/"

# Grab the repo directory (use rsync in case of circular crap)
echo "Copying repos..."
rsync -ah --progress "/home/$USER/repos" "$BACKUPDIR/"

# Grab the data directories I care about
cp -r "/home/$USER/.zsh_history" "$BACKUPDIR/.zsh_history"

# Compress the backup
tar czvf "$BACKUPDIR.tar.gz" "$BACKUPDIR"
rm -rf "$BACKUPDIR"
